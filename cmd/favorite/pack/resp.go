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

func NewCheckIsFavoriteResponse(err error) *favorite.CheckIsFavoriteResponse {
	return conv.ToKitexBaseResponse(err, &favorite.CheckIsFavoriteResponse{}).(*favorite.CheckIsFavoriteResponse)
}

func NewGetVideoFavoriteCountResponse(err error) *favorite.GetVideoFavoriteCountResponse {
	return conv.ToKitexBaseResponse(err, &favorite.GetVideoFavoriteCountResponse{}).(*favorite.GetVideoFavoriteCountResponse)
}
