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

// GetFavoriteInfo implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) GetVideoFavoriteInfo(ctx context.Context, request *favorite.GetVideoFavoriteInfoRequest) (resp *favorite.GetVideoFavoriteInfoResponse, err error) {
	favorite_service := service.NewFavoriteService(ctx)
	resp, err = favorite_service.GetVideoFavoriteInfo(ctx, request)
	return
}

// GetUserFavoriteInfo implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) GetUserFavoriteInfo(ctx context.Context, request *favorite.GetUserFavoriteInfoRequest) (resp *favorite.GetUserFavoriteInfoResponse, err error) {
	favorite_service := service.NewFavoriteService(ctx)
	resp, err = favorite_service.GetUserFavoriteInfo(ctx, request)
	return
}
