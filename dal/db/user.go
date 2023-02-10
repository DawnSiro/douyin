package db

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"

	"douyin/pkg/constant"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username       string `gorm:"column:username;index:idx_username,unique;type:varchar(40);not null" json:"username"`
	Password       string `gorm:"type:varchar(256);not null" json:"password"`
	FollowingCount int64  `gorm:"default:0" json:"following_count"`
	FollowerCount  int64  `gorm:"default:0" json:"follower_count"`
	Avatar         string `gorm:"type:varchar(256)" json:"avatar"`
}

func (u *User) TableName() string {
	return constant.UserTableName
}

// Register create user
func Register(ctx context.Context, user *User) (userID int64, err error) {
	if err := DB.WithContext(ctx).Create(user).Error; err != nil {
		return 0, err
	}
	return int64(user.ID), nil
}

// SelectUserByName query list of user info
func SelectUserByName(ctx context.Context, userName string) ([]*User, error) {
	res := make([]*User, 0)
	if err := DB.WithContext(ctx).Where("username = ?", userName).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func SelectUserByID(ctx context.Context, userID int64) (res *User, err error) {
	res = new(User)
	if err = DB.WithContext(ctx).Where("id = ?", userID).Take(res).Error; err != nil {
		klog.Info("select error")
		return nil, err
	}
	return res, nil
}
