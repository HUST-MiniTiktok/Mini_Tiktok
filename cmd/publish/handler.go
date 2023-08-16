package main

import (
	"context"

	publish "github.com/HUST-MiniTiktok/mini_tiktok/cmd/publish/kitex_gen/publish"
	service "github.com/HUST-MiniTiktok/mini_tiktok/cmd/publish/service"
)

// PublishServiceImpl implements the last service interface defined in the IDL.
type PublishServiceImpl struct{}

// PublishAction implements the PublishServiceImpl interface.
func (s *PublishServiceImpl) PublishAction(ctx context.Context, request *publish.PublishActionRequest) (resp *publish.PublishActionResponse, err error) {
	publish_service := service.NewPublishService(ctx)
	resp, err = publish_service.PublishAction(request)
	return
}

// PublishList implements the PublishServiceImpl interface.
func (s *PublishServiceImpl) PublishList(ctx context.Context, request *publish.PublishListRequest) (resp *publish.PublishListResponse, err error) {
	// TODO: Your code here...
	return
}
