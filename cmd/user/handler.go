package main

import (
	"context"
	"douyin/cmd/user/service"
	"douyin/dal/pack"
	"douyin/kitex_gen/user"
	"douyin/pkg/errno"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// Register implements the UserServiceImpl interface.
func (s *UserServiceImpl) Register(ctx context.Context, req *user.DouyinUserRegisterRequest) (resp *user.DouyinUserRegisterResponse, err error) {
	resp = new(user.DouyinUserRegisterResponse)

	if err = req.IsValid(); err != nil {
		resp.StatusCode, resp.StatusMsg = pack.BuildCodeAndMsg(errno.ParamErr)
		return resp, nil
	}

	userID, err := service.NewRegisterService(ctx).Register(req)
	if err != nil {
		resp.StatusCode, resp.StatusMsg = pack.BuildCodeAndMsg(err)
		return resp, nil
	}

	resp.StatusCode, resp.StatusMsg = pack.BuildCodeAndMsg(errno.Success)
	resp.UserId = userID
	return resp, nil
}

// Login implements the UserServiceImpl interface.
func (s *UserServiceImpl) Login(ctx context.Context, req *user.DouyinUserLoginRequest) (resp *user.DouyinUserLoginResponse, err error) {
	resp = new(user.DouyinUserLoginResponse)

	if err = req.IsValid(); err != nil {
		resp.StatusCode, resp.StatusMsg = pack.BuildCodeAndMsg(errno.ParamErr)
		return resp, nil
	}

	userID, err := service.NewLoginService(ctx).Login(req)
	if err != nil {
		resp.StatusCode, resp.StatusMsg = pack.BuildCodeAndMsg(err)
		return resp, nil
	}

	resp.StatusCode, resp.StatusMsg = pack.BuildCodeAndMsg(errno.Success)
	resp.UserId = userID
	return resp, nil
}

// GetUserInfo implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetUserInfo(ctx context.Context, req *user.DouyinUserRequest) (resp *user.DouyinUserResponse, err error) {
	resp = new(user.DouyinUserResponse)

	if err = req.IsValid(); err != nil {
		resp.StatusCode, resp.StatusMsg = pack.BuildCodeAndMsg(errno.ParamErr)
		return resp, nil
	}

	userInfo, err := service.NewInfoService(ctx).GetUserInfo(req)
	if err != nil {
		resp.StatusCode, resp.StatusMsg = pack.BuildCodeAndMsg(err)
		return resp, nil
	}

	resp.StatusCode, resp.StatusMsg = pack.BuildCodeAndMsg(errno.Success)
	resp.User = userInfo
	return resp, nil
}
