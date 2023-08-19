package pack

import (
	favorite "github.com/HUST-MiniTiktok/mini_tiktok/kitex_gen/favorite"
	"github.com/HUST-MiniTiktok/mini_tiktok/pkg/utils/conv"
)

func NewFavoriteActionResponse(err error) *favorite.FavoriteActionResponse {
	return conv.ToKitexBaseResponse(err, &favorite.FavoriteActionResponse{}).(*favorite.FavoriteActionResponse)
}

func NewFavoriteListResponse(err error) *favorite.FavoriteListResponse {
	return conv.ToKitexBaseResponse(err, &favorite.FavoriteListResponse{}).(*favorite.FavoriteListResponse)
}

func NewGetVideoFavoriteInfoResponse(err error) *favorite.GetVideoFavoriteInfoResponse {
	return conv.ToKitexBaseResponse(err, &favorite.GetVideoFavoriteInfoResponse{}).(*favorite.GetVideoFavoriteInfoResponse)
}

func NewGetUserFavoriteInfoResponse(err error) *favorite.GetUserFavoriteInfoResponse {
	return conv.ToKitexBaseResponse(err, &favorite.GetUserFavoriteInfoResponse{}).(*favorite.GetUserFavoriteInfoResponse)
}
