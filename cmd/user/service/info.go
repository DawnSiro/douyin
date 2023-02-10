package service

import (
	"context"
	"douyin/dal/db"
	"douyin/kitex_gen/user"
	"github.com/cloudwego/kitex/pkg/klog"
)

type InfoService struct {
	ctx context.Context
}

// NewInfoService new InfoService
func NewInfoService(ctx context.Context) *InfoService {
	return &InfoService{
		ctx: ctx,
	}
}

func (s *InfoService) GetUserInfo(req *user.DouyinUserRequest) (*user.User, error) {
	klog.Info("db before")
	u, err := db.SelectUserByID(s.ctx, req.UserId)
	if err != nil {
		klog.Info("data error")
		return nil, err
	}
	klog.Info("db after")

	// find is follow
	followCount := u.FollowingCount
	followerCount := u.FollowerCount

	// pack
	return &user.User{
		Id:            int64(u.ID),
		Name:          u.Username,
		FollowCount:   &followCount,
		FollowerCount: &followerCount,
		IsFollow:      false,
		Avatar:        u.Avatar,
	}, nil
}
