package service

import (
	"context"

	"douyin/cmd/user/pack"
	"douyin/dal/db"
	"douyin/kitex_gen/user"
	"douyin/pkg/errno"
	"douyin/pkg/util"

	"github.com/cloudwego/kitex/pkg/klog"
)

type UserService struct {
	ctx context.Context
}

func NewUserService(ctx context.Context) *UserService {
	return &UserService{
		ctx: ctx,
	}
}

func (s *UserService) Register(username, password string) (*user.DouyinUserRegisterResponse, error) {
	users, err := db.SelectUserByName(s.ctx, username)
	if err != nil {
		klog.Error("service.user.Register err:", err.Error())
		return nil, err
	}
	if len(users) != 0 {
		klog.Error("service.user.Register err:", errno.UsernameAlreadyExistsError.Error())
		return nil, errno.UsernameAlreadyExistsError
	}

	// 进行加密并存储
	encryptedPassword := util.BcryptHash(password)
	userID, err := db.CreateUser(s.ctx, &db.User{
		Username: username,
		Password: encryptedPassword,
	})
	if err != nil {
		klog.Error("service.user.Register err:", err.Error())
		return nil, err
	}
	token, err := util.SignToken(userID)
	if err != nil {
		klog.Error("service.user.Register err:", err.Error())
		return nil, err
	}
	return &user.DouyinUserRegisterResponse{
		StatusCode: 0,
		UserId:     int64(userID),
		Token:      token,
	}, nil
}

func (s *UserService) Login(username, password string) (*user.DouyinUserLoginResponse, error) {
	users, err := db.SelectUserByName(s.ctx, username)
	if err != nil {
		klog.Error("service.user.Login err:", err.Error())
		return nil, err
	}
	if len(users) == 0 {
		klog.Error("service.user.Login err:", errno.UserAccountDoesNotExistError.Error())
		return nil, errno.UserAccountDoesNotExistError
	}

	u := users[0]
	if !util.BcryptCheck(password, u.Password) {
		klog.Error("service.user.Login err:", errno.UserPasswordError.Error())
		return nil, errno.UserPasswordError
	}
	token, err := util.SignToken(u.ID)
	return &user.DouyinUserLoginResponse{
		StatusCode: 0,
		UserId:     int64(u.ID),
		Token:      token,
	}, nil
}

func (s *UserService) GetUserInfo(userID, infoUserID uint64) (*user.DouyinUserResponse, error) {
	u, err := db.SelectUserByID(s.ctx, infoUserID)
	if err != nil {
		klog.Error("service.user.GetUserInfo err:", err.Error())
		return nil, err
	}

	// TODO 使用 Redis Hash 来对用户数据进行缓存

	// pack
	userInfo := pack.UserInfo(u, db.IsFollow(s.ctx, userID, infoUserID))
	return &user.DouyinUserResponse{
		StatusCode: errno.Success.ErrCode,
		User:       userInfo,
	}, nil
}
