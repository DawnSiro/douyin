package service

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"

	"douyin/cmd/message/pack"
	"douyin/dal/db"
	"douyin/kitex_gen/message"
	"douyin/pkg/errno"
)

type MessageService struct {
	ctx context.Context
}

func NewMessageService(ctx context.Context) *MessageService {
	return &MessageService{
		ctx: ctx,
	}
}

func (s *MessageService) SendMessage(fromUserID, toUserID uint64, content string) (*message.DouyinMessageActionResponse, error) {
	isFriend := db.IsFriend(s.ctx, fromUserID, toUserID)
	if !isFriend {
		errNo := errno.UserRequestParameterError
		errNo.ErrMsg = "不能给非好友发消息"
		klog.Error("service.message.SendMessage err:", errNo.Error())
		return nil, errNo
	}
	err := db.CreateMessage(s.ctx, fromUserID, toUserID, content)
	if err != nil {
		klog.Error("service.message.SendMessage err:", err.Error())
		return nil, err
	}
	return &message.DouyinMessageActionResponse{
		StatusCode: errno.Success.ErrCode,
	}, nil
}

func (s *MessageService) GetMessageChat(userID, oppositeID uint64, preMsgTime int64) (*message.DouyinMessageChatResponse, error) {
	if userID == oppositeID {
		return nil, errno.UserRequestParameterError
	}
	messages, err := db.GetMessagesByUserIDAndPreMsgTime(s.ctx, userID, oppositeID, preMsgTime)
	if err != nil {
		klog.Error("service.message.GetMessageChat err:", err.Error())
		return nil, err
	}
	return &message.DouyinMessageChatResponse{
		StatusCode:  errno.Success.ErrCode,
		MessageList: pack.Messages(messages),
	}, nil
}
