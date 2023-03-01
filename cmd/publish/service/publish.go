package service

import (
	"bytes"
	"context"
	"errors"
	"io"
	"time"

	"douyin/cmd/publish/pack"
	"douyin/dal/db"
	"douyin/kitex_gen/publish"
	"douyin/pkg/errno"
	"douyin/pkg/util"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/gofrs/uuid"
)

type PublishService struct {
	ctx context.Context
}

func NewPublishService(ctx context.Context) *PublishService {
	return &PublishService{
		ctx: ctx,
	}
}

func (s *PublishService) PublishAction(title string, videoData []byte, userID uint64) (*publish.DouyinPublishActionResponse, error) {
	if userID == 0 {
		err := errors.New("userID error")
		klog.Error("service.publish.PublishAction err:", err.Error())
		return nil, err
	}

	// 上传 Object 需要一个实现了 io.Reader 接口的结构体
	var reader io.Reader = bytes.NewReader(videoData)
	u1, err := uuid.NewV4()
	if err != nil {
		klog.Error("service.publish.PublishAction err:", err.Error())
		return nil, err
	}
	fileName := u1.String() + "." + "mp4"
	klog.Info("service.publish.PublishAction videoName:", fileName)
	// 上传视频并生成封面
	playURL, coverURL, err := util.UploadVideo(&reader, fileName)
	if err != nil {
		klog.Error("service.publish.PublishAction err:", err.Error())
		return nil, err
	}

	err = db.CreateVideo(s.ctx, &db.Video{
		PublishTime:   time.Now(),
		AuthorID:      userID,
		PlayURL:       playURL,
		CoverURL:      coverURL,
		FavoriteCount: 0,
		CommentCount:  0,
		Title:         title,
	})
	if err != nil {
		klog.Error("service.publish.PublishAction err:", err.Error())
		return nil, err
	}

	return &publish.DouyinPublishActionResponse{
		StatusCode: errno.Success.ErrCode,
	}, nil
}

func (s *PublishService) GetPublishVideos(userID uint64) (*publish.DouyinPublishListResponse, error) {
	videoList := make([]*publish.Video, 0)

	videos, err := db.GetVideosByAuthorID(s.ctx, userID)
	if err != nil {
		klog.Error("service.publish.GetPublishVideos err:", err.Error())
		return nil, err
	}

	for i := 0; i < len(videos); i++ {
		u, err := db.SelectUserByID(s.ctx, videos[i].AuthorID)
		if err != nil {
			klog.Error("service.publish.GetPublishVideos err:", err.Error())
			return nil, err
		}

		video := pack.Video(videos[i], u,
			db.IsFollow(s.ctx, userID, u.ID), db.IsFavoriteVideo(s.ctx, userID, videos[i].ID))
		videoList = append(videoList, video)
	}

	return &publish.DouyinPublishListResponse{
		StatusCode: errno.Success.ErrCode,
		VideoList:  videoList,
	}, nil
}
