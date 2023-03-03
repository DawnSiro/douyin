package pack

import (
	"douyin/dal/db"
	"douyin/kitex_gen/comment"
	"douyin/pkg/errno"

	"github.com/cloudwego/kitex/pkg/klog"
)

func Comment(dbc *db.Comment, dbu *db.User, isFollow bool) *comment.Comment {
	if dbc == nil || dbu == nil {
		klog.Error("pack.comment.Comment err:", errno.ServiceError)
		return nil
	}

	return &comment.Comment{
		Id:         int64(dbc.ID),
		User:       User(dbu, isFollow),
		Content:    dbc.Content,
		CreateDate: dbc.CreatedTime.Format("01-02"), // 评论发布日期，格式 mm-dd
	}
}

func User(u *db.User, isFollow bool) *comment.User {
	if u == nil {
		klog.Error("pack.user.User err:", errno.ServiceError)
		return nil
	}
	followCount := u.FollowingCount
	followerCount := u.FollowerCount
	return &comment.User{
		Id:            int64(u.ID),
		Name:          u.Username,
		FollowCount:   &followCount,
		FollowerCount: &followerCount,
		IsFollow:      isFollow,
		Avatar:        u.Avatar,
	}
}

func CommentData(data *db.CommentData) *comment.Comment {
	if data == nil {
		return nil
	}
	followCount := int64(data.FollowingCount)
	followerCount := int64(data.FollowerCount)
	u := &comment.User{
		Id:            int64(data.UID),
		Name:          data.Username,
		FollowCount:   &followCount,
		FollowerCount: &followerCount,
		IsFollow:      data.IsFollow,
		Avatar:        data.Avatar,
	}
	return &comment.Comment{
		Id:         int64(data.CID),
		User:       u,
		Content:    data.Content,
		CreateDate: data.CreatedTime.Format("01-02"), // 评论发布日期，格式 mm-dd
	}
}

func CommentDataList(cdList []*db.CommentData) []*comment.Comment {
	res := make([]*comment.Comment, 0, len(cdList))
	for i := 0; i < len(cdList); i++ {
		res = append(res, CommentData(cdList[i]))
	}
	return res
}
