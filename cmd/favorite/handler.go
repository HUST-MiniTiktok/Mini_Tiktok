package main

import (
	"context"

	service "github.com/HUST-MiniTiktok/mini_tiktok/cmd/favorite/service"
	favorite "github.com/HUST-MiniTiktok/mini_tiktok/kitex_gen/favorite"
)

// FavoriteServiceImpl implements the last service interface defined in the IDL.
type FavoriteServiceImpl struct{}

// FavoriteAction implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) FavoriteAction(ctx context.Context, request *favorite.FavoriteActionRequest) (resp *favorite.FavoriteActionResponse, err error) {
	favorite_service := service.NewFavoriteService(ctx)
	resp, err = favorite_service.FavoriteAction(ctx, request)
	return
}

// FavoriteList implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) FavoriteList(ctx context.Context, request *favorite.FavoriteListRequest) (resp *favorite.FavoriteListResponse, err error) {
	favorite_service := service.NewFavoriteService(ctx)
	resp, err = favorite_service.FavoriteList(ctx, request)
	return
}

// GetVideoFavoriteCount implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) GetVideoFavoriteCount(ctx context.Context, request *favorite.GetVideoFavoriteCountRequest) (resp *favorite.GetVideoFavoriteCountResponse, err error) {
	favorite_service := service.NewFavoriteService(ctx)
	resp, err = favorite_service.GetVideoFavoriteCount(ctx, request)
	return
}

// CheckIsFavorite implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) CheckIsFavorite(ctx context.Context, request *favorite.CheckIsFavoriteRequest) (resp *favorite.CheckIsFavoriteResponse, err error) {
	favorite_service := service.NewFavoriteService(ctx)
	resp, err = favorite_service.CheckIsFavorite(ctx, request)
	return
}
