package main

import (
	"context"

	"douyin/cmd/favorite/service"
	"douyin/kitex_gen/favorite"
	"douyin/pkg/constant"
	"douyin/pkg/errno"
	"douyin/pkg/util"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/kitex/pkg/klog"
)

// FavoriteServiceImpl implements the last service interface defined in the IDL.
type FavoriteServiceImpl struct{}

// FavoriteVideo implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) FavoriteVideo(ctx context.Context, req *favorite.DouyinFavoriteActionRequest) (resp *favorite.DouyinFavoriteActionResponse, err error) {
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

	if req.ActionType == constant.Favorite {
		resp, err = service.NewFavoriteService(ctx).FavoriteVideo(claim.UserID, uint64(req.VideoId))
	} else if req.ActionType == constant.CancelFavorite {
		resp, err = service.NewFavoriteService(ctx).CancelFavoriteVideo(claim.UserID, uint64(req.VideoId))
	} else {
		err = errno.UserRequestParameterError
		hlog.Error("handler.favorite_service.FavoriteVideo err:", err.Error())
	}
	return
}

// GetFavoriteList implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) GetFavoriteList(ctx context.Context, req *favorite.DouyinFavoriteListRequest) (resp *favorite.DouyinFavoriteListResponse, err error) {
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
	return service.NewFavoriteService(ctx).GetFavoriteList(claim.UserID, uint64(req.UserId))
}
