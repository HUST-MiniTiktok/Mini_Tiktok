package dal

import (
	"github.com/HUST-MiniTiktok/mini_tiktok/cmd/publish/dal/db"
	"github.com/HUST-MiniTiktok/mini_tiktok/cmd/publish/dal/oss"
)

// Init Mysql DB
func Init() {
	db.Init()
	oss.Init()
}