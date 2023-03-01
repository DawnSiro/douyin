package main

import (
	"context"

	"douyin/cmd/relation/service"
	"douyin/kitex_gen/relation"
	"douyin/pkg/constant"
	"douyin/pkg/errno"
	"douyin/pkg/util"

	"github.com/cloudwego/kitex/pkg/klog"
)

// RelationServiceImpl implements the last service interface defined in the IDL.
type RelationServiceImpl struct{}

// Follow implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) Follow(ctx context.Context, req *relation.DouyinRelationActionRequest) (resp *relation.DouyinRelationActionResponse, err error) {
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
	if req.ActionType == constant.Follow {
		resp, err = service.NewRelationService(ctx).Follow(claim.UserID, uint64(req.ToUserId))
	} else if req.ActionType == constant.CancelFollow {
		resp, err = service.NewRelationService(ctx).CancelFollow(claim.UserID, uint64(req.ToUserId))
	}
	err = errno.UserRequestParameterError
	klog.Error("handler.relation_service.Follow err:", err.Error())
	return
}

// GetFollowList implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) GetFollowList(ctx context.Context, req *relation.DouyinRelationFollowListRequest) (resp *relation.DouyinRelationFollowListResponse, err error) {
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
	return service.NewRelationService(ctx).GetFollowList(claim.UserID, uint64(req.UserId))
}

// GetFollowerList implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) GetFollowerList(ctx context.Context, req *relation.DouyinRelationFollowerListRequest) (resp *relation.DouyinRelationFollowerListResponse, err error) {
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
	return service.NewRelationService(ctx).GetFollowerList(claim.UserID, uint64(req.UserId))
}

// GetFriendList implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) GetFriendList(ctx context.Context, req *relation.DouyinRelationFriendListRequest) (resp *relation.DouyinRelationFriendListResponse, err error) {
	if err := req.IsValid(); err != nil {
		errNo := errno.UserRequestParameterError
		errNo.ErrMsg = err.Error()
		return nil, errNo
	}

	// 解析 Token，再验一次权
	claim, err := util.ParseToken(req.Token)
	if err != nil {
		klog.Error("handler.handler.GetFriendList err:", err.Error())
		return nil, err
	}
	if claim.UserID == uint64(req.UserId) {
		return nil, errno.UserRequestParameterError
	}
	return service.NewRelationService(ctx).GetFriendList(claim.UserID)
}
