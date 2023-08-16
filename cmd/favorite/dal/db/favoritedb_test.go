package db

import (
	"context"
	"log"
	"testing"
)

func TestInit(t *testing.T) {
	Init()
}

func TestNewFavorite(t *testing.T) {
	ctx := context.Background()
	Init()
	NewFavorite(ctx, 1, 1)
}

func TestCancelFavorite(t *testing.T) {
	ctx := context.Background()
	Init()
	CancelFavorite(ctx, 1, 1)
}

func TestGetFavoriteList(t *testing.T) {
	ctx := context.Background()
	Init()
	// NewFavorite(ctx, 1, 3)
	// NewFavorite(ctx, 2, 2)
	status, videoList, err := GetFavoriteList(ctx, 1)
	log.Println(status)
	log.Println(videoList)
	log.Println(err)
}
