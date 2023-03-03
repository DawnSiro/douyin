package pack

import (
	"douyin/dal/db"
	"douyin/kitex_gen/feed"
	"douyin/pkg/errno"

	"github.com/cloudwego/kitex/pkg/klog"
)

func Video(v *db.Video, u *db.User, isFollow, isFavorite bool) *feed.Video {
	if v == nil || u == nil {
		klog.Error("pack.video.Video err:", errno.ServiceError)
		return nil
	}
	author := &feed.UserInfo{
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
	return &feed.Video{
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

func VideoData(data *db.VideoData) *feed.Video {
	if data == nil {
		return nil
	}
	followCount := data.FollowCount
	followerCount := data.FollowerCount
	author := &feed.UserInfo{
		Id:              int64(data.UID),
		Name:            data.Username,
		FollowCount:     followCount,
		FollowerCount:   followerCount,
		IsFollow:        data.IsFollow,
		Avatar:          data.Avatar,
		BackgroundImage: data.BackgroundImage,
		Signature:       data.Signature,
		TotalFavorited:  data.TotalFavorited,
		WorkCount:       data.WorkCount,
		FavoriteCount:   data.UserFavoriteCount,
	}
	return &feed.Video{
		Id:            int64(data.VID),
		Author:        author,
		PlayUrl:       data.PlayURL,
		CoverUrl:      data.CoverURL,
		FavoriteCount: data.FavoriteCount,
		CommentCount:  data.CommentCount,
		IsFavorite:    data.IsFavorite,
		Title:         data.Title,
	}
}

func VideoDataList(dataList []*db.VideoData) []*feed.Video {
	res := make([]*feed.Video, 0, len(dataList))
	for i := 0; i < len(dataList); i++ {
		res = append(res, VideoData(dataList[i]))
	}
	return res
}
