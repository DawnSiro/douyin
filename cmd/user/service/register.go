package service

import (
	"context"
	"douyin/dal/db"
	"douyin/kitex_gen/user"
	"douyin/pkg/errno"
	"douyin/pkg/util"
)

type RegisterService struct {
	ctx context.Context
}

// NewRegisterService new CreateUserService
func NewRegisterService(ctx context.Context) *RegisterService {
	return &RegisterService{ctx: ctx}
}

// Register create user info.
func (s *RegisterService) Register(req *user.DouyinUserRegisterRequest) (int64, error) {
	users, err := db.SelectUserByName(s.ctx, req.Username)
	if err != nil {
		return 0, err
	}
	if len(users) != 0 {
		return 0, errno.UserAlreadyExistErr
	}

	// 进行加密并存储
	password := util.BcryptHash(req.Password)
	return db.Register(s.ctx, &db.User{
		Username: req.Username,
		Password: password,
	})
}
