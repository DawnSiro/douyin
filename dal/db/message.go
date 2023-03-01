package db

import (
	"context"
	"douyin/pkg/global"
	"time"

	"douyin/pkg/constant"
)

type Message struct {
	ID         uint64    `json:"id"`
	ToUserID   uint64    `gorm:"not null" json:"to_user_id"`
	FromUserID uint64    `gorm:"not null" json:"from_user_id"`
	Content    string    `gorm:"type:varchar(255);not null" json:"content"`
	CreateTime time.Time `gorm:"not null" json:"create_time" `
}

func (n *Message) TableName() string {
	return constant.MessageTableName
}

type FriendMessageResp struct {
	Content string
	MsgType uint8
}

func CreateMessage(ctx context.Context, fromUserID, toUserID uint64, content string) error {
	return global.DB.WithContext(ctx).Create(&Message{FromUserID: fromUserID, ToUserID: toUserID, Content: content, CreateTime: time.Now()}).Error
}

func GetMessagesByUserIDAndPreMsgTime(ctx context.Context, userID, oppositeID uint64, preMsgTime int64) ([]*Message, error) {
	res := make([]*Message, 0)
	message := &Message{}
	// 使用 Union 来避免使用 or 导致不走索引的问题
	err := global.DB.WithContext(ctx).Raw("? UNION ? ORDER BY create_time ASC",
		global.DB.WithContext(ctx).Where("to_user_id = ? AND from_user_id = ? AND `create_time` > ?",
			userID, oppositeID, time.UnixMilli(preMsgTime)).Model(message),
		global.DB.WithContext(ctx).Where("to_user_id = ? AND from_user_id = ? AND `create_time` > ?",
			oppositeID, userID, time.UnixMilli(preMsgTime)).Model(message),
	).Scan(&res).Error
	if err != nil {
		return nil, err
	}

	return res, nil
}

func GetLatestMsg(ctx context.Context, userID uint64, oppositeID uint64) (*FriendMessageResp, error) {
	message := &Message{}
	// 使用 Union 来避免使用 or 导致不走索引的问题
	err := global.DB.WithContext(ctx).Raw("? UNION ? ORDER BY create_time DESC LIMIT 1",
		global.DB.WithContext(ctx).Where("to_user_id = ? AND from_user_id = ?", userID, oppositeID).Model(message),
		global.DB.WithContext(ctx).Where("to_user_id = ? AND from_user_id = ?", oppositeID, userID).Model(message),
	).Scan(&message).Error
	if err != nil {
		return nil, err
	}

	switch message.ToUserID {
	case oppositeID:
		return &FriendMessageResp{
			Content: message.Content,
			MsgType: constant.SentMessage,
		}, nil
	default: // 默认发给自己
		return &FriendMessageResp{
			Content: message.Content,
			MsgType: constant.ReceivedMessage,
		}, nil
	}
}
