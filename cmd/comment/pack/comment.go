package pack

import (
	"douyin/dal/db"
	"douyin/kitex_gen/comment"
	"douyin/pkg/errno"

	"github.com/cloudwego/kitex/pkg/klog"
)

func Comment(dbc *db.Comment, dbu *db.User, isFollow bool) *comment.Comment {
	if dbc == nil || dbu == nil {
		klog.Error("pack.comment.Comment err:", errno.ServiceError)
		return nil
	}

	return &comment.Comment{
		Id:         int64(dbc.ID),
		User:       User(dbu, isFollow),
		Content:    dbc.Content,
		CreateDate: dbc.CreatedTime.Format("01-02"), // 评论发布日期，格式 mm-dd
	}
}

func User(u *db.User, isFollow bool) *comment.User {
	if u == nil {
		klog.Error("pack.user.User err:", errno.ServiceError)
		return nil
	}
	followCount := u.FollowingCount
	followerCount := u.FollowerCount
	return &comment.User{
		Id:            int64(u.ID),
		Name:          u.Username,
		FollowCount:   &followCount,
		FollowerCount: &followerCount,
		IsFollow:      isFollow,
		Avatar:        u.Avatar,
	}
}
