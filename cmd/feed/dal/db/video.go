package db

import (
	"context"
	"gorm.io/gorm"
	"time"
)

const VideoTableName = "video"

type Video struct {
	gorm.Model
	AuthorID    int64
	PlayURL     string
	CoverURL    string
	PublishTime time.Time
	Title       string
}

func (Video) TableName() string {
	return VideoTableName
}

func CreateVideo(ctx context.Context, video *Video) (id int64, err error) {
	err = DB.WithContext(ctx).Create(video).Error
	if err != nil {
		return -1, err
	}
	id = int64(video.ID)
	return
}

func CheckVideoExistById(ctx context.Context, id int64) (exist bool, err error) {
	var count int64
	err = DB.WithContext(ctx).Model(&Video{}).Where("id = ?", id).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count != 0, nil
}

func GetVideoById(ctx context.Context, id int64) (video *Video, err error) {
	err = DB.WithContext(ctx).First(&video, id).Error
	if err != nil {
		return nil, err
	}
	return
}

func GetVideoByAuthorId(ctx context.Context, authorId int64) (videos []*Video, err error) {
	err = DB.WithContext(ctx).Where("author_id = ?", authorId).Find(&videos).Error
	if err != nil {
		return nil, err
	}
	return
}

func GetVideosByLastPublishTime(ctx context.Context, lastPublishTime time.Time) (videos []*Video, err error) {
	err = DB.WithContext(ctx).Where("publish_time < ?", lastPublishTime).Order("publish_time desc").Limit(30).Find(&videos).Error
	if err != nil {
		return nil, err
	}
	return
}

func GetVideosByIDs(ctx context.Context, ids []int64) (videos []*Video, err error) {
	for _, id := range ids {
		video, err := GetVideoById(ctx, id)
		if err != nil {
			return nil, err
		}
		videos = append(videos, video)
	}
	return
}
