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
		err := tx.WithContext(ctx).Model(&Comment{}).Create(&comment).Error
		if err != nil {
			return err
		}
		return nil
	})

	go RDIncrCommentCount(video_id, 1)

	return comment, err
}

// DelComment deletes a comment from the database.
func DelComment(ctx context.Context, comment_id int64, video_id int64) error {
	if err := DB.WithContext(ctx).Model(&Comment{}).Where(&Comment{ID: comment_id}).Delete(&Comment{}).Error; err != nil {
		return err
	}

	go RDIncrCommentCount(video_id, -1)

	return nil
}

// GetVideoComments returns a list of video comments.
func GetVideoComments(ctx context.Context, video_id int64) ([]*Comment, error) {
	var comments []*Comment
	err := DB.WithContext(ctx).Model(&Comment{}).Where(&Comment{VideoId: video_id}).Find(&comments).Error
	if err != nil {
		return nil, err
	}
	return comments, nil
}

// GetVideoComments returns the number of comments of the inputed video.
func GetVideoCommentCounts(ctx context.Context, video_id int64) (count int64, err error) {

	if RDExistCommentCount(video_id) {
		return RDGetCommentCount(video_id), nil
	}

	err = DB.WithContext(ctx).Model(&Comment{}).Where(&Comment{VideoId: video_id}).Count(&count).Error
	if err != nil {
		return 0, err
	}

	go RDSetCommentCount(video_id, count)

	return count, nil
}
