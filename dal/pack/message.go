package pack

import (
	"douyin/cmd/api/biz/model/api"
	"douyin/dal/db"
	"douyin/pkg/errno"

	"github.com/cloudwego/hertz/pkg/common/hlog"
)

func Messages(ms []*db.Message) []*api.Message {
	if ms == nil {
		hlog.Error("pack.message.Messages err:", errno.ServiceError)
		return nil
	}
	res := make([]*api.Message, 0)
	for i := 0; i < len(ms); i++ {
		res = append(res, Message(ms[i]))
	}
	return res
}

func Message(m *db.Message) *api.Message {
	if m == nil {
		hlog.Error("pack.message.Messages err:", errno.ServiceError)
		return nil
	}
	createTime := m.CreateTime.UnixMilli()
	return &api.Message{
		ID:         int64(m.ID),
		ToUserID:   int64(m.ToUserID),
		FromUserID: int64(m.FromUserID),
		Content:    m.Content,
		CreateTime: &createTime,
	}
}
