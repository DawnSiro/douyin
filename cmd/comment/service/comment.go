package service

import (
	"context"
	"reflect"
	"strconv"
	"strings"

	"douyin/cmd/comment/pack"
	"douyin/dal/db"
	"douyin/kitex_gen/comment"
	"douyin/pkg/errno"
	"douyin/pkg/global"
	"douyin/pkg/util/sensitive"

	"github.com/cloudwego/hertz/pkg/common/json"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/go-redis/redis"
)

type CommentService struct {
	ctx context.Context
}

// NewCommentService new CommentService
func NewCommentService(ctx context.Context) *CommentService {
	return &CommentService{
		ctx: ctx,
	}
}

func (s *CommentService) PostComment(userID, videoID uint64, commentText string) (*comment.DouyinCommentActionResponse, error) {
	// 删除redis评论列表缓存
	// 使用 strings.Builder 来优化字符串的拼接
	var builder strings.Builder
	builder.WriteString(strconv.FormatUint(videoID, 10))
	builder.WriteString("_video_comments")
	delCommentListKey := builder.String()

	// TODO 业务优化
	keysMatch, err := global.VideoCRC.Do("keys", "*"+delCommentListKey).Result()
	if err != nil {
		klog.Error("service.comment.PostComment err:", err.Error())
	}
	if reflect.TypeOf(keysMatch).Kind() == reflect.Slice {
		val := reflect.ValueOf(keysMatch)
		// 删除key
		for i := 0; i < val.Len(); i++ {
			global.VideoCRC.Del(val.Index(i).Interface().(string))
			klog.Info("删除了RedisKey:", val.Index(i).Interface().(string))
		}
	}

	//检测是否带有敏感词
	if sensitive.IsWordsFilter(commentText) {
		return nil, errno.ContainsProhibitedSensitiveWordsError
	}

	dbc, err := db.CreateComment(s.ctx, videoID, commentText, userID)
	if err != nil {
		klog.Error("service.comment.PostComment err:", err.Error())
		return nil, err
	}

	dbu, err := db.SelectUserByID(s.ctx, userID)
	if err != nil {
		klog.Error("service.comment.PostComment err:", err.Error())
		return nil, err
	}
	authorID, err := db.SelectAuthorIDByVideoID(s.ctx, videoID)
	if err != nil {
		klog.Error("service.comment.PostComment err:", err.Error())
		return nil, err
	}

	return &comment.DouyinCommentActionResponse{
		StatusCode: 0,
		Comment:    pack.Comment(dbc, dbu, db.IsFollow(s.ctx, userID, authorID)),
	}, nil
}

func (s *CommentService) DeleteComment(userID, videoID, commentID uint64) (*comment.DouyinCommentActionResponse, error) {
	// 查询此评论是否是本人发送的
	isComment := db.IsCommentCreatedByMyself(s.ctx, userID, commentID)
	// 非本人评论
	if !isComment {
		klog.Error("service.comment.DeleteComment err:", errno.DeletePermissionError)
		return nil, errno.DeletePermissionError
	}

	dbc, err := db.DeleteCommentByID(s.ctx, videoID, commentID)
	if err != nil {
		klog.Error("service.comment.DeleteComment err:", err.Error())
		return nil, err
	}
	dbu, err := db.SelectUserByID(s.ctx, userID)
	if err != nil {
		klog.Error("service.comment.DeleteComment err:", err.Error())
		return nil, err
	}
	authorID, err := db.SelectAuthorIDByVideoID(s.ctx, videoID)
	if err != nil {
		klog.Error("service.comment.DeleteComment err:", err.Error())
		return nil, err
	}

	return &comment.DouyinCommentActionResponse{
		StatusCode: 0,
		Comment:    pack.Comment(dbc, dbu, db.IsFollow(s.ctx, userID, authorID)),
	}, nil
}

func (s *CommentService) GetCommentList(userID, videoID uint64) (*comment.DouyinCommentListResponse, error) {
	var builder strings.Builder
	builder.WriteString(strconv.FormatUint(userID, 10))
	builder.WriteString("_userId_")
	builder.WriteString(strconv.FormatUint(videoID, 10))
	builder.WriteString("_video_comments")
	commentListKey := builder.String()

	commentList, err := global.VideoCRC.Get(commentListKey).Result()
	if err == redis.Nil {
		dbcList, err := db.SelectCommentListByVideoID(s.ctx, videoID)
		if err != nil {
			klog.Error("service.comment.GetCommentList err:", err.Error())
			return nil, err
		}

		cList := make([]*comment.Comment, 0, len(dbcList))

		for i := 0; i < len(dbcList); i++ {
			u, _ := db.SelectUserByID(s.ctx, dbcList[i].UserID)
			cList = append(cList, pack.Comment(dbcList[i], u, db.IsFollow(s.ctx, userID, dbcList[i].UserID)))
		}

		//序列化
		marshalList, _ := json.Marshal(cList)
		_, err = global.VideoCRC.Set(commentListKey, marshalList, 0).Result()
		if err != nil {
			klog.Error("service.comment.GetCommentList err:", err.Error())
			return nil, err
		}
		commentList, err = global.VideoCRC.Get(commentListKey).Result()
		if err != nil {
			klog.Error("service.comment.GetCommentList err:", err.Error())
			return nil, err
		}
	}
	//反序列化
	var list []*comment.Comment
	err = json.Unmarshal([]byte(commentList), &list)
	if err != nil {
		klog.Error("service.comment.GetCommentList err:", err.Error())
		return nil, err
	}

	return &comment.DouyinCommentListResponse{
		StatusCode:  0,
		CommentList: list,
	}, nil
}
