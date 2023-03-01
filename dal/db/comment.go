package db

import (
	"context"
	"douyin/pkg/global"
	"errors"
	"time"

	"douyin/pkg/constant"
	"douyin/pkg/errno"

	"gorm.io/gorm"
)

type Comment struct {
	ID          uint64    `json:"id"`
	IsDeleted   uint8     `gorm:"default:0;not null" json:"is_deleted"`
	VideoID     uint64    `gorm:"not null" json:"video_id"`
	UserID      uint64    `gorm:"not null" json:"user_id"`
	Content     string    `gorm:"type:varchar(255);not null" json:"content"`
	CreatedTime time.Time `gorm:"not null" json:"created_time"`
}

func (n *Comment) TableName() string {
	return constant.CommentTableName
}

func CreateComment(ctx context.Context, videoID uint64, content string, userID uint64) (*Comment, error) {
	comment := &Comment{
		VideoID:     videoID,
		UserID:      userID,
		Content:     content,
		CreatedTime: time.Now(),
	}
	// DB 层开事务来保证原子性
	err := global.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 先查询 VideoID 是否存在，然后增加评论数，再创建评论
		video := &Video{
			ID: videoID,
		}
		err := global.DB.WithContext(ctx).First(&video).Error
		if err != nil {
			return err
		}
		// 增加视频评论数
		err = global.DB.WithContext(ctx).Model(&video).Update("comment_count", video.CommentCount+1).Error
		if err != nil {
			return err
		}
		// 创建评论
		return global.DB.WithContext(ctx).Create(comment).Error
	})
	if err != nil {
		return nil, err
	}

	return comment, nil
}

// DeleteCommentByID 通过评论ID 删除评论，默认使用软删除，提高性能
func DeleteCommentByID(ctx context.Context, videoID, commentID uint64) (*Comment, error) {
	comment := &Comment{
		ID: commentID,
	}
	// 先查询是否存在评论
	result := global.DB.WithContext(ctx).Where("is_deleted = ?", constant.DataNotDeleted).Limit(1).Find(comment)
	if result.RowsAffected == 0 {
		return nil, errors.New("delete data failed")
	}

	// DB 层开事务来保证原子性
	err := global.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 减少视频评论数
		video := &Video{
			ID: videoID,
		}
		err := global.DB.WithContext(ctx).First(&video).Error
		if err != nil {
			return err
		}
		err = global.DB.WithContext(ctx).Model(&video).Update("comment_count", video.CommentCount-1).Error
		if err != nil {
			return err
		}
		// 删除评论
		result = global.DB.WithContext(ctx).Model(comment).Update("is_deleted", constant.DataDeleted)
		if result.Error != nil {
			return result.Error
		}
		if result.RowsAffected == 0 {
			return errno.UserRequestParameterError
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return comment, nil
}

func SelectCommentListByVideoID(ctx context.Context, videoID uint64) ([]*Comment, error) {
	res := make([]*Comment, 0)
	err := global.DB.WithContext(ctx).Where("video_id = ? AND is_deleted = ?", videoID, constant.DataNotDeleted).Find(&res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}

func IsCommentCreatedByMyself(ctx context.Context, userID uint64, commentID uint64) bool {
	result := global.DB.WithContext(ctx).Where("id = ? AND user_id = ? AND is_deleted = ?", commentID, userID, constant.DataNotDeleted).
		Find(&Comment{})
	if result.RowsAffected == 0 {
		return false
	}
	return true
}
