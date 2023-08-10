package db

import (
	"time"
	"gorm.io/gorm"
)

const VideoTableName = "video"

type Video struct {
	gorm.Model
	AuthorID    int64 		`json:"authorID"`
	PlayURL     string 		`json:"playURL"`
	CoverURL    string 		`json:"coverURL"`
	PublishTime time.Time 	`json:"publishTime"`
	Title       string		`json:"title"`
}

func (Video) TableName() string {
	return VideoTableName
}