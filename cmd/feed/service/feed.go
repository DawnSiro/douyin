package service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"

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
	videoData, err := db.MSelectFeedVideoDataListByUserID(s.ctx, constant.MaxVideoNum, latestTime, userID)
	if err != nil {
		klog.Error("service.feed.GetFeed err:", err.Error())
		return nil, err
	}
	var nextTime *int64
	// 没有视频的时候 nextTime 为 nil，会重置时间
	if len(videoData) != 0 {
		nextTime = new(int64)
		*nextTime, err = db.SelectPublishTimeByVideoID(s.ctx, videoData[len(videoData)-1].VID)
		if err != nil {
			hlog.Error("service.feed.GetFeed err:", err.Error())
			return nil, err
		}
	}

	return &feed.DouyinFeedResponse{
		StatusCode: errno.Success.ErrCode,
		VideoList:  pack.VideoDataList(videoData),
		NextTime:   nextTime,
	}, nil
}
