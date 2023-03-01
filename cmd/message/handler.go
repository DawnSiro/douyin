package main

import (
	"context"

	"douyin/cmd/message/service"
	"douyin/kitex_gen/message"
	"douyin/pkg/constant"
	"douyin/pkg/errno"
	"douyin/pkg/util"

	"github.com/cloudwego/kitex/pkg/klog"
)

// MessageServiceImpl implements the last service interface defined in the IDL.
type MessageServiceImpl struct{}

// SendMessage implements the MessageServiceImpl interface.
func (s *MessageServiceImpl) SendMessage(ctx context.Context, req *message.DouyinMessageActionRequest) (resp *message.DouyinMessageActionResponse, err error) {
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
	if req.ActionType == constant.SendMessageAction {
		return service.NewMessageService(ctx).SendMessage(claim.UserID, uint64(req.ToUserId), req.Content)
	}
	return nil, errno.UserRequestParameterError
}

// GetMessageChat implements the MessageServiceImpl interface.
func (s *MessageServiceImpl) GetMessageChat(ctx context.Context, req *message.DouyinMessageChatRequest) (resp *message.DouyinMessageChatResponse, err error) {
	// 解析 Token，再验一次权
	claim, err := util.ParseToken(req.Token)
	if err != nil {
		klog.Error("handler.handler.CommentAction err:", err.Error())
		return nil, err
	}
	return service.NewMessageService(ctx).GetMessageChat(claim.UserID, uint64(req.ToUserId), *req.PreMsgTime)
}
