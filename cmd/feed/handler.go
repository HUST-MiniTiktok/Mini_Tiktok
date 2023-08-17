package main

import (
	"context"

	feed "github.com/HUST-MiniTiktok/mini_tiktok/kitex_gen/feed"
	service "github.com/HUST-MiniTiktok/mini_tiktok/cmd/feed/service"
)

// FeedServiceImpl implements the last service interface defined in the IDL.
type FeedServiceImpl struct{}

// GetFeed implements the FeedServiceImpl interface.
func (s *FeedServiceImpl) GetFeed(ctx context.Context, request *feed.FeedRequest) (resp *feed.FeedResponse, err error) {
	feed_service := service.NewFeedService(ctx)
	resp, err = feed_service.GetFeed(request)
	return
}
