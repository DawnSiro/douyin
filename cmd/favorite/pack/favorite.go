package pack

import (
	"douyin/dal/db"
	"douyin/kitex_gen/favorite"
	"douyin/pkg/errno"

	"github.com/cloudwego/kitex/pkg/klog"
)

func Video(v *db.Video, u *db.User, isFollow, isFavorite bool) *favorite.Video {
	if v == nil || u == nil {
		klog.Error("pack.video.Video err:", errno.ServiceError)
		return nil
	}
	author := &favorite.UserInfo{
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
	return &favorite.Video{
		Id:            int64(v.ID),
		Author:        author,
		PlayUrl:       v.PlayURL,
		CoverUrl:      v.CoverURL,
		FavoriteCount: v.FavoriteCount,
		CommentCount:  v.CommentCount,
		IsFavorite:    isFavorite,
		Title:         v.Title,
	}
}