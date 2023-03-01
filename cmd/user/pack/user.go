package pack

import (
	"douyin/dal/db"
	"douyin/kitex_gen/user"
	"douyin/pkg/errno"

	"github.com/cloudwego/kitex/pkg/klog"
)

func UserInfo(u *db.User, isFollow bool) *user.UserInfo {
	if u == nil {
		klog.Error("pack.user.UserInfo err:", errno.ServiceError)
		return nil
	}
	return &user.UserInfo{
		Id:              int64(u.ID),
		Name:            u.Username,
		FollowCount:     u.FollowingCount,
		FollowerCount:   u.FollowerCount,
		IsFollow:        isFollow,
		Avatar:          u.Avatar,
		BackgroundImage: u.BackgroundImage,
		Signature:       u.Signature,
		TotalFavorited:  u.TotalFavorited,
		WorkCount:       u.WorkCount,
		FavoriteCount:   u.FavoriteCount,
	}
}
