package main

import (
	"context"
	"douyin/pkg/errno"

	"douyin/cmd/publish/service"
	"douyin/kitex_gen/publish"
	"douyin/pkg/util"

	"github.com/cloudwego/kitex/pkg/klog"
)

// PublishServiceImpl implements the last service interface defined in the IDL.
type PublishServiceImpl struct{}

// PublishAction implements the PublishServiceImpl interface.
func (s *PublishServiceImpl) PublishAction(ctx context.Context, req *publish.DouyinPublishActionRequest) (resp *publish.DouyinPublishActionResponse, err error) {
	if err := req.IsValid(); err != nil {
		errNo := errno.UserRequestParameterError
		errNo.ErrMsg = err.Error()
		return nil, errNo
	}

	// 解析 Token，再验一次权
	claim, err := util.ParseToken(req.Token)
	if err != nil {
		klog.Error("handler.handler.CommentAction err:", err.Error())
		return nil, err
	}
	return service.NewPublishService(ctx).PublishAction(req.Title, req.Data, claim.UserID)
}

// GetPublishVideos implements the PublishServiceImpl interface.
func (s *PublishServiceImpl) GetPublishVideos(ctx context.Context, req *publish.DouyinPublishListRequest) (resp *publish.DouyinPublishListResponse, err error) {
	if err := req.IsValid(); err != nil {
		errNo := errno.UserRequestParameterError
		errNo.ErrMsg = err.Error()
		return nil, errNo
	}

	// 解析 Token，再验一次权
	claim, err := util.ParseToken(req.Token)
	if err != nil {
		klog.Error("handler.handler.CommentAction err:", err.Error())
		return nil, err
	}
	return service.NewPublishService(ctx).GetPublishVideos(claim.UserID, uint64(req.UserId))
}
