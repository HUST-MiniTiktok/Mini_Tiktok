package db

import (
	"context"
	"time"

	"gorm.io/gorm"
)

const CommentTableName = "comment"

type Comment struct {
	ID          int64          `json:"id"`
	UserId      int64          `json:"user_id"`
	VideoId     int64          `json:"video_id"`
	CommentText string         `json:"comment_text"`
	CreatedAt   time.Time      `json:"created_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

func (Comment) TableName() string {
	return CommentTableName
}

// NewComment creates a new Comment
func NewComment(ctx context.Context, user_id int64, video_id int64, comment_text string) (comment Comment, err error) {
	comment = Comment{
		UserId:      user_id,
		VideoId:     video_id,
		CommentText: comment_text,
	}
	err = DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 1.新增评论数据
		err := tx.WithContext(ctx).Model(&Comment{}).Create(&comment).Error
		if err != nil {
			return err
		}
		return nil
	})
	return comment, err
}

// DelComment deletes a comment from the database.
func DelComment(ctx context.Context, commentID int64, vid int64) error {
	err := DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		comment := new(Comment)
		if err := tx.WithContext(ctx).First(&comment, commentID).Error; err != nil { //主键查找
			return err
		}
		// tx.Unscoped().Delete()将永久删除记录
		err := tx.WithContext(ctx).Unscoped().Delete(&comment).Error
		if err != nil {
			return err
		}
		return nil
	})
	return err
}

// GetVideoComments returns a list of video comments.
func GetVideoComments(ctx context.Context, vid int64) ([]*Comment, error) {
	var comments []*Comment
	err := DB.WithContext(ctx).Model(&Comment{}).Where(&Comment{VideoId: vid}).Find(&comments).Error
	if err != nil {
		return nil, err
	}
	return comments, nil
}

// GetVideoComments returns the number of comments of the inputed video.
func GetVideoCommentCounts(ctx context.Context, vid int64) (count int64, err error) {

	err = DB.WithContext(ctx).Model(&Comment{}).Where(&Comment{VideoId: vid}).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}
