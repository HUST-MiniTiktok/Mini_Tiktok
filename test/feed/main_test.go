package feed_test

import (
	"context"
	"os"
	"testing"

	dal "github.com/HUST-MiniTiktok/mini_tiktok/cmd/feed/dal"
	service "github.com/HUST-MiniTiktok/mini_tiktok/cmd/feed/service"
)

var (
	ctx         = context.Background()
	FeedService *service.FeedService
)

func TestMain(m *testing.M) {
	os.Setenv("GO_ENV", "test")
	dal.Init()
	FeedService = service.NewFeedService(ctx)

	code := m.Run()
	os.Exit(code)
}

func TestMainOrder(t *testing.T) {
	t.Run("feed", TestFeed)
}
