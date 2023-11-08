package db

import (
	"context"
	"time"
)

const VideoTableName = "video"
const feedNum = 5

type Video struct {
	ID          int64     `json:"id" gorm:"primaryKey;autoincrement"`
	AuthorID    int64     `json:"author_id" gorm:"index:video_author_idx"`
	PlayURL     string    `json:"play_url"`
	CoverURL    string    `json:"cover_url"`
	PublishTime time.Time `json:"publish_time"`
	Title       string    `json:"title" gorm:"type:varchar(255)"`
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

func GetVideoIdListByAuthorId(ctx context.Context, authorId int64) (video_ids []int64, err error) {
	err = DB.WithContext(ctx).Model(&Video{}).Where("author_id = ?", authorId).Pluck("id", &video_ids).Error
	if err != nil {
		return nil, err
	}
	return
}

func GetVideosByLastPublishTime(ctx context.Context, lastPublishTime time.Time) (videos []*Video, err error) {
	err = DB.WithContext(ctx).Where("publish_time < ?", lastPublishTime).Order("publish_time desc").Limit(feedNum).Find(&videos).Error
	if err != nil {
		return nil, err
	}
	return
}

func GetVideosByIdList(ctx context.Context, ids []int64) (videos []*Video, err error) {
	err = DB.WithContext(ctx).Where("id in ?", ids).Find(&videos).Error
	if err != nil {
		return nil, err
	}
	return
}
