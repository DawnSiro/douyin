package service

import (
	"context"
	"douyin/dal/db"
	"douyin/kitex_gen/user"
	"douyin/pkg/errno"
	"douyin/pkg/util"
)

type LoginService struct {
	ctx context.Context
}

// NewLoginService new LoginService
func NewLoginService(ctx context.Context) *LoginService {
	return &LoginService{
		ctx: ctx,
	}
}

// Login check user info
func (s *LoginService) Login(req *user.DouyinUserLoginRequest) (userID int64, err error) {
	users, err := db.SelectUserByName(s.ctx, req.Username)
	if err != nil {
		return 0, err
	}
	if len(users) == 0 {
		return 0, errno.AuthorizationFailedErr
	}

	u := users[0]
	if !util.BcryptCheck(req.Password, u.Password) {
		return 0, errno.AuthorizationFailedErr
	}
	return int64(u.ID), nil
}
