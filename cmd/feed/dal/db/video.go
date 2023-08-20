package db

import (
	"context"
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
	id = video.ID
	return
}

func CheckVideoExistById(ctx context.Context, id int64) (exist bool, err error) {
	var db_video Video
	err = DB.WithContext(ctx).Model(&Video{}).Where("id = ?", id).Limit(1).Find(&db_video).Error
	if err != nil {
		return false, err
	}
	return db_video != Video{}, nil
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

	for _, id := range ids {
		go func(id int64) {
			video, err := GetVideoById(ctx, id)
			err_chan <- err
			video_chan <- video
		}(id)
	}
	for i := 0; i < len(ids); i++ {
		err = <-err_chan
		if err != nil {
			return nil, err
		}
		videos = append(videos, <-video_chan)
	}
	return
}
