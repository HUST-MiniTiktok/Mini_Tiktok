package db

import (
	"context"
	"time"
)

const VideoTableName = "video"

type Video struct {
	ID          int64     `json:"id" gorm:"primaryKey;autoincrement"`
	AuthorID    int64     `json:"author_id" gorm:"index:author_id_idx"`
	PlayURL     string    `json:"play_url"`
	CoverURL    string    `json:"cover_url"`
	PublishTime time.Time `json:"publish_time"`
	Title       string    `json:"title"`
}

func (Video) TableName() string {
	return VideoTableName
}

// CreateVideo: create a new video record
func CreateVideo(ctx context.Context, video *Video) (id int64, err error) {
	err = DB.WithContext(ctx).Create(video).Error
	if err != nil {
		return -1, err
	}
	id = video.ID
	return
}

// CheckVideoExistById: check if a video exists by id
func CheckVideoExistById(ctx context.Context, id int64) (exist bool, err error) {
	var db_video Video
	err = DB.WithContext(ctx).Model(&Video{}).Where("id = ?", id).Limit(1).Find(&db_video).Error
	if err != nil {
		return false, err
	}
	return db_video != Video{}, nil
}

// GetVideoById: get a video by id
func GetVideoById(ctx context.Context, id int64) (video *Video, err error) {
	err = DB.WithContext(ctx).First(&video, id).Error
	if err != nil {
		return nil, err
	}
	return
}

// GetVideoByAuthorId: get videos by author id
func GetVideoByAuthorId(ctx context.Context, authorId int64) (videos []*Video, err error) {
	err = DB.WithContext(ctx).Where("author_id = ?", authorId).Find(&videos).Error
	if err != nil {
		return nil, err
	}
	return
}

// GetVideosByLastPublishTime: get videos by last publish time
func GetVideosByLastPublishTime(ctx context.Context, lastPublishTime time.Time) (videos []*Video, err error) {
	err = DB.WithContext(ctx).Where("publish_time < ?", lastPublishTime).Order("publish_time desc").Limit(30).Find(&videos).Error
	if err != nil {
		return nil, err
	}
	return
}

// GetVideosByIDs: get videos by video ids
func GetVideosByIDs(ctx context.Context, ids []int64) (videos []*Video, err error) {
	err = DB.WithContext(ctx).Where("id in ?", ids).Find(&videos).Error
	if err != nil {
		return nil, err
	}
	return
}
