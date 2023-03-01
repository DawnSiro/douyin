package service

import (
	"context"
	"strconv"
	"strings"

	"douyin/cmd/favorite/pack"
	"douyin/dal/db"
	"douyin/kitex_gen/favorite"
	"douyin/pkg/errno"
	"douyin/pkg/global"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/go-redis/redis"
)

type FavoriteService struct {
	ctx context.Context
}

// NewFavoriteService new favoriteService
func NewFavoriteService(ctx context.Context) *FavoriteService {
	return &FavoriteService{
		ctx: ctx,
	}
}

// FavoriteVideo this is a func for add Favorite or reduce Favorite
func (s *FavoriteService) FavoriteVideo(userID, videoID uint64) (*favorite.DouyinFavoriteActionResponse, error) {
	var builder strings.Builder
	builder.WriteString(strconv.FormatUint(videoID, 10))
	builder.WriteString("_video_like")
	videoLikeKey := builder.String()

	likeCount, err := global.VideoFRC.Get(videoLikeKey).Result()
	if err == redis.Nil {
		likeInt64, err := db.SelectVideoFavoriteCountByVideoID(s.ctx, videoID)
		if err != nil {
			klog.Error("service.favorite.FavoriteVideo err:", err.Error())
			return nil, err
		}
		global.VideoFRC.Set(videoLikeKey, likeInt64, 0)
	}
	var likeUint64 uint64
	if likeCount != "" {
		likeUint64, err = strconv.ParseUint(likeCount, 10, 64)
		if err != nil {
			klog.Error("service.favorite.FavoriteVideo err:", err.Error())
			return nil, err
		}
	}

	err = db.FavoriteVideo(s.ctx, userID, videoID)
	if err != nil {
		klog.Error("service.favorite.FavoriteVideo err:", err.Error())
		return nil, err
	}
	// 如果 DB 层事务回滚了，err 就不为 nil，Redis 里的数据就不会更新
	global.VideoFRC.Set(videoLikeKey, likeUint64+1, 0)

	// TODO 用缓存记录用户点赞数量，防止刷赞

	return &favorite.DouyinFavoriteActionResponse{
		StatusCode: 0,
	}, nil
}

func (s *FavoriteService) CancelFavoriteVideo(userID, videoID uint64) (*favorite.DouyinFavoriteActionResponse, error) {
	var builder strings.Builder
	builder.WriteString(strconv.FormatUint(videoID, 10))
	builder.WriteString("_video_like")
	videoLikeKey := builder.String()

	likeCount, err := global.VideoFRC.Get(videoLikeKey).Result()
	if err == redis.Nil {
		likeInt64, err := db.SelectVideoFavoriteCountByVideoID(s.ctx, videoID)
		if err != nil {
			klog.Error("service.favorite.CancelFavoriteVideo err:", err.Error())
			return nil, err
		}
		global.VideoFRC.Set(videoLikeKey, likeInt64, 0)
	}

	var likeUint64 uint64
	if likeCount != "" {
		likeUint64, err = strconv.ParseUint(likeCount, 10, 64)
		if err != nil {
			klog.Error("service.favorite.CancelFavoriteVideo err:", err.Error())
			return nil, err
		}
	}

	err = db.CancelFavoriteVideo(s.ctx, userID, videoID)
	if err != nil {
		klog.Error("service.favorite.CancelFavoriteVideo err:", err.Error())
		return nil, err
	}
	// 如果 DB 层事务回滚了，err 就不为 nil，Redis 里的数据就不会更新
	global.VideoFRC.Set(videoLikeKey, likeUint64-1, 0)

	return &favorite.DouyinFavoriteActionResponse{
		StatusCode: errno.Success.ErrCode,
	}, nil
}

func (s *FavoriteService) GetFavoriteList(userID, selectUserID uint64) (*favorite.DouyinFavoriteListResponse, error) {
	videos, err := db.SelectFavoriteVideoListByUserID(s.ctx, selectUserID)
	if err != nil {
		klog.Error("service.favorite.GetFavoriteList err:", err.Error())
		return nil, err
	}

	// TODO 优化循环查询数据库问题
	videoList := make([]*favorite.Video, 0)
	for i := 0; i < len(videos); i++ {
		u, err := db.SelectUserByID(s.ctx, videos[i].AuthorID)
		if err != nil {
			klog.Error("service.favorite.GetFavoriteList err:", err.Error())
			return nil, err
		}
		video := pack.Video(videos[i], u,
			db.IsFollow(s.ctx, userID, selectUserID), db.IsFavoriteVideo(s.ctx, userID, videos[i].ID))
		videoList = append(videoList, video)
	}

	return &favorite.DouyinFavoriteListResponse{
		StatusCode: errno.Success.ErrCode,
		VideoList:  videoList,
	}, nil
}
