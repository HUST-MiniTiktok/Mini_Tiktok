package main

import (
	"context"
	feed "github.com/HUST-MiniTiktok/mini_tiktok/cmd/feed/kitex_gen/feed"
)

// FeedServiceImpl implements the last service interface defined in the IDL.
type FeedServiceImpl struct{}

// GetFeed implements the FeedServiceImpl interface.
func (s *FeedServiceImpl) GetFeed(ctx context.Context, request *feed.FeedRequest) (resp *feed.FeedResponse, err error) {
	// TODO: Your code here...
	return
}

// GetVideoById implements the FeedServiceImpl interface.
func (s *FeedServiceImpl) GetVideoById(ctx context.Context, request *feed.GetVideoByIdRequest) (resp *feed.GetVideoByIdResponse, err error) {
	// TODO: Your code here...
	return
}
