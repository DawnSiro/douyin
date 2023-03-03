// Code generated by hertz generator.

package api

import (
	"context"

	"douyin/cmd/api/biz/model/api"
	"douyin/cmd/api/biz/rpc"
	"douyin/kitex_gen/feed"
	"douyin/pkg/errno"
	"douyin/pkg/global"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// GetFeed .
// @router /douyin/feed/ [GET]
func GetFeed(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.DouyinFeedRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.JSON(consts.StatusOK, &api.DouyinResponse{
			StatusCode: errno.UserRequestParameterError.ErrCode,
			StatusMsg:  err.Error(),
		})
		return
	}

	hlog.Infof("handler.feed_service.GetFeed Request: %#v", req)
	userID := c.GetUint64(global.Config.JWTConfig.IdentityKey)
	hlog.Info("handler.feed_service.GetFeed GetUserID:", userID)
	resp, err := rpc.GetFeed(context.Background(), &feed.DouyinFeedRequest{
		LatestTime: req.LatestTime,
		Token:      req.Token,
	})
	if err != nil {
		errNo := errno.ConvertErr(err)
		c.JSON(consts.StatusOK, &api.DouyinFeedResponse{
			StatusCode: errNo.ErrCode,
			StatusMsg:  &errNo.ErrMsg,
		})
		return
	}

	c.JSON(consts.StatusOK, resp)
}
