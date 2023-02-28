package main

import (
	"context"

	"douyin/cmd/comment/service"
	"douyin/kitex_gen/comment"
	"douyin/pkg/constant"
	"douyin/pkg/errno"
	"douyin/pkg/util"

	"github.com/cloudwego/kitex/pkg/klog"
)

// CommentServiceImpl implements the last service interface defined in the IDL.
type CommentServiceImpl struct{}

// CommentAction implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) CommentAction(ctx context.Context, req *comment.DouyinCommentActionRequest) (resp *comment.DouyinCommentActionResponse, err error) {
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

	// 这里注意走 ActionType 对应的逻辑的时候要注意判断相关字段是否为空
	if req.ActionType == constant.PostComment && req.CommentText != nil {
		resp, err = service.NewCommentService(ctx).PostComment(claim.UserID, uint64(req.VideoId), *req.CommentText)
	} else if req.ActionType == constant.DeleteComment && req.CommentId != nil {
		resp, err = service.NewCommentService(ctx).DeleteComment(claim.UserID, uint64(req.VideoId), uint64(*req.CommentId))
	} else {
		err = errno.UserRequestParameterError
		klog.Error("handler.handler.CommentAction err:", err.Error())
	}
	return
}

// GetCommentList implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) GetCommentList(ctx context.Context, req *comment.DouyinCommentListRequest) (resp *comment.DouyinCommentListResponse, err error) {
	if err := req.IsValid(); err != nil {
		errNo := errno.UserRequestParameterError
		errNo.ErrMsg = err.Error()
		return nil, errNo
	}

	// 解析 Token，再验一次权
	claim, err := util.ParseToken(req.Token)
	if err != nil {
		klog.Error("handler.handler.GetCommentList err:", err.Error())
		return nil, err
	}

	resp, err = service.NewCommentService(ctx).GetCommentList(claim.UserID, uint64(req.VideoId))
	if err != nil {
		klog.Error("handler.handler.GetCommentList err:", err.Error())
		return nil, err
	}
	return
}
