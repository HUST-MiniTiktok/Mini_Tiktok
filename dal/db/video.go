package db

import (
	"context"
	"gorm.io/gorm"
	"time"
)

const VideoTableName = "video"

type Video struct {
	gorm.Model
	AuthorID    int64     `json:"authorID"`
	PlayURL     string    `json:"playURL"`
	CoverURL    string    `json:"coverURL"`
	PublishTime time.Time `json:"publishTime"`
	Title       string    `json:"title"`
}

func (Video) TableName() string {
	return VideoTableName
}

func GetVideoFeed(ctx context.Context, latest_time int64) ([]*Video, error) {
	result := make([]*Video, 0)
	if err := DB.WithContext(ctx).
		Where("created_at < ?", latest_time).
		Order("created_at desc").
		Find(&result).
		Error; err != nil {
		return nil, err
	}
	return result, nil
}

//TODO：其他的数据库操作，比如查询、更新等
