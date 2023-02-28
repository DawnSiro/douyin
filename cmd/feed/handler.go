package main

import (
	"context"
	"douyin/pkg/errno"

	"douyin/cmd/feed/service"
	"douyin/kitex_gen/feed"
	"douyin/pkg/util"
)

// FeedServiceImpl implements the last service interface defined in the IDL.
type FeedServiceImpl struct{}

// GetFeed implements the FeedServiceImpl interface.
func (s *FeedServiceImpl) GetFeed(ctx context.Context, req *feed.DouyinFeedRequest) (resp *feed.DouyinFeedResponse, err error) {
	if err := req.IsValid(); err != nil {
		errNo := errno.UserRequestParameterError
		errNo.ErrMsg = err.Error()
		return nil, errNo
	}

	// 解析 Token
	var userID uint64
	if req.Token != nil {
		claim, _ := util.ParseToken(*req.Token)
		userID = claim.UserID
	}
	return service.NewFeedService(ctx).GetFeed(req.LatestTime, userID)
}
