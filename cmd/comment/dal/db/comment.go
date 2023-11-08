package db

import (
	"context"
	"strconv"
	"time"

	"github.com/HUST-MiniTiktok/mini_tiktok/pkg/mw/redis"
	"github.com/cloudwego/kitex/pkg/klog"

	"gorm.io/gorm"
)

const (
	CommentTableName = "comment"
)

type Comment struct {
	ID          int64          `json:"id" gorm:"primaryKey;autoincrement"`
	UserId      int64          `json:"user_id"`
	VideoId     int64          `json:"video_id" gorm:"index:comment_video_idx"`
	CommentText string         `json:"comment_text"`
	CreatedAt   time.Time      `json:"created_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

func (Comment) TableName() string {
	return CommentTableName
}

// NewComment creates a new Comment
func NewComment(ctx context.Context, user_id int64, video_id int64, comment_text string) (comment *Comment, err error) {
	comment = &Comment{
		UserId:      user_id,
		VideoId:     video_id,
		CommentText: comment_text,
	}
	err = DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		err := tx.WithContext(ctx).Model(&Comment{}).Create(comment).Error
		if err != nil {
			return err
		}
		return nil
	})

	go RDIncrCommentCount(video_id, 1)
	go Filter.AddToBloomFilter(strconv.Itoa(int(video_id)))

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
func GetVideoCommentCount(ctx context.Context, video_id int64) (count int64, err error) {

	if RDExistCommentCount(video_id) {
		return RDGetCommentCount(video_id), nil
	}

	// cache do not exist，try to query from mysql
	if RDClient.AcquireLock(video_id, CommentCountField) {
		defer RDClient.ReleaseLock(video_id, CommentCountField)

		exist := Filter.TestBloom(strconv.Itoa(int(video_id)))
		if !exist { // video not exist
			return 0, nil
		}

		err = DB.WithContext(ctx).Model(&Comment{}).Where(&Comment{VideoId: video_id}).Count(&count).Error
		if err != nil {
			return 0, err
		}
		go RDSetCommentCount(video_id, count)
		return count, nil
	}

	time.Sleep(redis.RetryTime) // delay and retry
	return GetVideoCommentCount(ctx, video_id)
}

// GetVideoComments returns the number of comments of the inputed video.
//
//	func GetVideoCommentCountList(ctx context.Context, video_id_list []int64) (count_list []int64, err error) {
//		size := len(video_id_list)
//		for i := 0; i < size; i++ {
//			count, err := GetVideoCommentCount(ctx, video_id_list[i])
//			if err != nil {
//				return nil, err
//			}
//			count_list[i] = count
//		}
//		return count_list, nil
//	}
func GetVideoCommentCountList(ctx context.Context, video_id_list []int64) (count_list []int64, err error) {
	size := len(video_id_list)
	var sql_video_id_list []int64
	var sql_video_id_order []int
	for i, j := 0, 0; i < size; i++ {
		if RDExistCommentCount(video_id_list[i]) {
			count_list[i] = RDGetCommentCount(video_id_list[i])
		} else {
			sql_video_id_order[j] = i
			sql_video_id_list[j] = video_id_list[i]
			j++
		}
	}
	if len(sql_video_id_list) == 0 {
		return count_list, nil
	}
	// to do
	// 批量查询
	// err = DB.WithContext(ctx).Model(&Comment{}).Group("VideoId").Count(&count_list).Error
	// if err != nil {
	// 	return nil, err
	// }

	return count_list, nil
}

func LoadCommentVideoIDToBloomFilter(ctx context.Context) error {
	var videoIdList []string

	err := DB.WithContext(ctx).Model(&Comment{}).Pluck("VideoId", &videoIdList).Error
	if err != nil {
		klog.Errorf("Load Comment VideoID To BloomFilter Failed: %v", err)
		return err
	}

	for _, video_id := range videoIdList {
		Filter.AddToBloomFilter(video_id)
	}
	return nil
}
