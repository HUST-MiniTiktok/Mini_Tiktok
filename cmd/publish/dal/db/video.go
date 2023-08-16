package db

import (
	"context"
	"sync"
	"time"

)

const VideoTableName = "video"

type Video struct {
	ID          int64     `json:"id"`
	AuthorID    int64     `json:"author_id"`
	PlayURL     string    `json:"play_url"`
	CoverURL    string    `json:"cover_url"`
	PublishTime time.Time `json:"publish_time"`
	Title       string    `json:"title"`
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
	err_chan := make(chan error)
	video_chan := make(chan *Video, len(ids))
	var wg sync.WaitGroup
	for _, id := range ids {
		wg.Add(1)
		go func(id int64) {
			defer wg.Done()
			video, err := GetVideoById(ctx, id)
			err_chan <- err
			video_chan <- video
		}(id)
	}
	for i := 0; i < len(ids); i++ {
		err = <- err_chan
		if err != nil {
			return nil, err
		}
		videos = append(videos, <-video_chan)
	}
	return
}
