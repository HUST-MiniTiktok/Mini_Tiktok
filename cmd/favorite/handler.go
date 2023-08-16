package main

import (
	"context"

	favorite "github.com/HUST-MiniTiktok/mini_tiktok/cmd/favorite/kitex_gen/favorite"
	service "github.com/HUST-MiniTiktok/mini_tiktok/cmd/favorite/service"
)

// FavoriteServiceImpl implements the last service interface defined in the IDL.
type FavoriteServiceImpl struct{}

// FavoriteAction implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) FavoriteAction(ctx context.Context, request *favorite.FavoriteActionRequest) (resp *favorite.FavoriteActionResponse, err error) {
	user_service := service.NewFavoriteService(ctx)
	resp, err = user_service.FavoriteAction(ctx, request)
	return
}

// FavoriteList implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) FavoriteList(ctx context.Context, request *favorite.FavoriteListRequest) (resp *favorite.FavoriteListResponse, err error) {
	user_service := service.NewFavoriteService(ctx)
	resp, err = user_service.FavoriteList(ctx, request)
	return
}
