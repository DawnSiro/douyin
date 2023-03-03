package mw

import (
	"context"

	"douyin/pkg/errno"
	"douyin/pkg/global"
	"douyin/pkg/util"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func JWT() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		token := c.Query("token")
		if token == "" {
			hlog.Error("mw.jwt.ParseToken err:", errno.UserIdentityVerificationFailedError)
			c.JSON(consts.StatusBadRequest, utils.H{
				"status_code": errno.UserIdentityVerificationFailedError.ErrCode,
				"status_msg":  "Token 为空",
			})
			c.Abort()
			return
		}
		claim, err := util.ParseToken(token)
		if err != nil {
			hlog.Error("mw.jwt.ParseToken err:", err.Error())
			c.JSON(consts.StatusBadRequest, utils.H{
				"status_code": errno.UserIdentityVerificationFailedError.ErrCode,
				"status_msg":  errno.UserIdentityVerificationFailedError.ErrMsg,
			})
			c.Abort()
			return
		}
		hlog.Info("mw.jwt.ParseToken userID:", claim.UserID)
		c.Set(global.Config.JWTConfig.IdentityKey, claim.UserID)
		c.Next(ctx)
	}
}

// ParseToken 如果 Token 存在，会试着解析 Token，不存在也会放行。主要用于某些登录和未登录都能使用的接口
func ParseToken() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		token := c.Query("token")
		if token == "" {
			return
		}
		claim, err := util.ParseToken(token)
		if err != nil {
			hlog.Info("mw.jwt.ParseToken err:", err.Error())
			return
		}
		hlog.Info("mw.jwt.ParseToken userID:", claim.UserID)
		c.Set(global.Config.JWTConfig.IdentityKey, claim.UserID)
		return
	}
}
