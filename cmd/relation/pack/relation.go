package pack

import (
	"douyin/dal/db"
	"douyin/kitex_gen/relation"
	"douyin/pkg/errno"

	"github.com/cloudwego/kitex/pkg/klog"
)

func User(u *db.User, isFollow bool) *relation.User {
	if u == nil {
		klog.Error("pack.user.User err:", errno.ServiceError)
		return nil
	}
	followCount := u.FollowingCount
	followerCount := u.FollowerCount
	return &relation.User{
		Id:            int64(u.ID),
		Name:          u.Username,
		FollowCount:   &followCount,
		FollowerCount: &followerCount,
		IsFollow:      isFollow,
		Avatar:        u.Avatar,
	}
}

func FriendUser(u *db.User, isFollow bool, messageContent string, msgType uint8) *relation.FriendUser {
	if u == nil {
		klog.Error("pack.user.UserInfo err:", errno.ServiceError)
		return nil
	}
	followCount := u.FollowingCount
	followerCount := u.FollowerCount
	return &relation.FriendUser{
		Id:            int64(u.ID),
		Name:          u.Username,
		FollowCount:   &followCount,
		FollowerCount: &followerCount,
		IsFollow:      isFollow,
		Avatar:        u.Avatar,
		Message:       &messageContent,
		MsgType:       int8(msgType),
	}
}
