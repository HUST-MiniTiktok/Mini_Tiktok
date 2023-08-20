package db

import (
	"context"
	"log"
	"testing"
	"time"
)

func TestInit(t *testing.T) {
	Init()
	// DB.Where(&Favorite{}).Delete(&Favorite{})
}

func TestNewFavorite(t *testing.T) {
	ctx := context.Background()
	Init()
	status, err := NewFavorite(ctx, 1, 1)
	log.Println(status)
	log.Println(err)
	status, err = NewFavorite(ctx, 1, 2)
	log.Println(status)
	log.Println(err)
	status, err = NewFavorite(ctx, 2, 1)
	log.Println(status)
	log.Println(err)
	status, err = NewFavorite(ctx, 2, 2)
	log.Println(status)
	log.Println(err)
}

func TestCancelFavorite(t *testing.T) {
	ctx := context.Background()
	Init()
	status, err := CancelFavorite(ctx, 1, 1)
	// CancelFavorite(ctx, 1, 2)
	log.Println(status)
	log.Println(err)
}

func TestCheckFavorite(t *testing.T) {
	ctx := context.Background()
	Init()
	exist, err := CheckFavorite(ctx, 1, 1)
	log.Println(exist)
	log.Println(err)
	NewFavorite(ctx, 1, 1)
	exist, err = CheckFavorite(ctx, 1, 1)
	log.Println(exist)
	log.Println(err)
	CancelFavorite(ctx, 1, 1)
	exist, err = CheckFavorite(ctx, 1, 1)
	log.Println(exist)
	log.Println(err)
}

func TestVideoFavoriteCount(t *testing.T) {
	ctx := context.Background()
	Init()
	count, err := VideoFavoriteCount(ctx, 1)
	log.Println(count)
	log.Println(err)
	count, err = VideoFavoriteCount(ctx, 2)
	log.Println(count)
	log.Println(err)

	NewFavorite(ctx, 1, 1)
	NewFavorite(ctx, 1, 2)
	NewFavorite(ctx, 2, 1)
	NewFavorite(ctx, 2, 2)
	time.Sleep(time.Second)
	count, err = VideoFavoriteCount(ctx, 1)
	log.Println(count)
	log.Println(err)
	count, err = VideoFavoriteCount(ctx, 2)
	log.Println(count)
	log.Println(err)

	CancelFavorite(ctx, 1, 1)
	CancelFavorite(ctx, 2, 1)
	CancelFavorite(ctx, 1, 2)
	CancelFavorite(ctx, 2, 2)
	time.Sleep(time.Second)
	count, err = VideoFavoriteCount(ctx, 1)
	log.Println(count)
	log.Println(err)
	count, err = VideoFavoriteCount(ctx, 2)
	log.Println(count)
	log.Println(err)
	count, err = VideoFavoriteCount(ctx, 3)
	log.Println(count)
	log.Println(err)
}

func TestGetFavoriteList(t *testing.T) {
	ctx := context.Background()
	Init()
	NewFavorite(ctx, 1, 1)
	NewFavorite(ctx, 1, 2)
	NewFavorite(ctx, 2, 1)
	NewFavorite(ctx, 2, 2)
	status, videoList, err := GetFavoriteList(ctx, 1)
	log.Println(status)
	log.Println(videoList)
	log.Println(err)
	status, videoList, err = GetFavoriteList(ctx, 2)
	log.Println(status)
	log.Println(videoList)
	log.Println(err)
	CancelFavorite(ctx, 1, 1)
	// CancelFavorite(ctx, 1, 2)
	CancelFavorite(ctx, 2, 1)
	CancelFavorite(ctx, 2, 2)
	status, videoList, err = GetFavoriteList(ctx, 1)
	log.Println(status)
	log.Println(videoList)
	log.Println(err)
	status, videoList, err = GetFavoriteList(ctx, 2)
	log.Println(status)
	log.Println(videoList)
	log.Println(err)
	status, videoList, err = GetFavoriteList(ctx, 3)
	log.Println(status)
	log.Println(videoList)
	log.Println(err)
}
