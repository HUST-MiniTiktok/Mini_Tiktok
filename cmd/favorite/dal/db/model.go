package db

import (
	"time"

	"gorm.io/gorm"
)

type Favorite struct {
	ID        int64          `json:"id"`
	UserId    int64          `json:"user_id"`
	VideoId   int64          `json:"video_id"`
	CreatedAt time.Time      `json:"create_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"delete_at"`
}

// type Video struct {
// 	ID          int64     `json:"id"`
// 	AuthorID    int64     `json:"author_id"`
// 	PlayURL     string    `json:"play_url"`
// 	CoverURL    string    `json:"cover_url"`
// 	PublishTime time.Time `json:"publish_time"`
// 	Title       string    `json:"title"`
// }

// type Message struct {
// 	ID         int64     `json:"id"`
// 	ToUserId   int64     `json:"to_user_id"`
// 	FromUserId int64     `json:"from_user_id"`
// 	Content    string    `json:"content"`
// 	CreatedAt  time.Time `json:"created_at"`
// }

// type Relation struct {
// 	ID         int64          `json:"id"`
// 	UserId     int64          `json:"user_id"`
// 	FollowerId int64          `json:"follower_id"`
// 	CreatedAt  time.Time      `json:"create_at"`
// 	DeletedAt  gorm.DeletedAt `gorm:"index" json:"delete_at"`
// }

// type User struct {
// 	ID              int64  `json:"id"`               // 用户ID
// 	UserName        string `json:"user_name"`        // 用户名
// 	Password        string `json:"password"`         // 密码
// 	Avatar          string `json:"avatar"`           // 头像路径
// 	BackgroundImage string `json:"background_image"` // 背景图片
// 	Signature       string `json:"signature"`        // 签名
// }

// type Comment struct {
// 	ID          int64          `json:"id"`
// 	UserId      int64          `json:"user_id"`
// 	VideoId     int64          `json:"video_id"`
// 	CommentText string         `json:"comment_text"`
// 	CreatedAt   time.Time      `json:"created_at"`
// 	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
// }
