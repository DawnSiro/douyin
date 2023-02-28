package service

import (
	"context"

	"douyin/cmd/feed/pack"
	"douyin/dal/db"
	"douyin/kitex_gen/feed"
	"douyin/pkg/constant"
	"douyin/pkg/errno"

	"github.com/cloudwego/kitex/pkg/klog"
)

type FeedService struct {
	ctx context.Context
}

func NewFeedService(ctx context.Context) *FeedService {
	return &FeedService{
		ctx: ctx,
	}
}

func (s *FeedService) GetFeed(latestTime *int64, userID uint64) (*feed.DouyinFeedResponse, error) {
	videoList := make([]*feed.Video, 0)

	videos, err := db.MGetVideos(s.ctx, constant.MaxVideoNum, latestTime)
	if err != nil {
		klog.Error("service.feed.GetFeed err:", err.Error())
		return nil, err
	}

	// TODO 使用预处理等进行优化
	for i := 0; i < len(videos); i++ {
		u, err := db.SelectUserByID(s.ctx, videos[i].AuthorID)
		if err != nil {
			klog.Error("service.feed.GetFeed err:", err.Error())
			return nil, err
		}

		var video *feed.Video
		// 未登录默认未关注，未点赞
		if userID == 0 {
			video = pack.Video(videos[i], u,
				false, false)
		} else {
			video = pack.Video(videos[i], u,
				db.IsFollow(s.ctx, userID, u.ID), db.IsFavoriteVideo(s.ctx, userID, videos[i].ID))
		}

		videoList = append(videoList, video)
	}

	var nextTime *int64
	// 没有视频的时候 nextTime 为 nil，会重置时间
	if len(videos) != 0 {
		nextTime = new(int64)
		*nextTime = videos[len(videos)-1].PublishTime.UnixMilli()
	}

	return &feed.DouyinFeedResponse{
		StatusCode: errno.Success.ErrCode,
		VideoList:  videoList,
		NextTime:   nextTime,
	}, nil
}
