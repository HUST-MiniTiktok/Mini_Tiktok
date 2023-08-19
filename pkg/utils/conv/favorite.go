package conv

import (
	hertz_favorite "github.com/HUST-MiniTiktok/mini_tiktok/cmd/api/biz/model/favorite"
	kitex_favorite "github.com/HUST-MiniTiktok/mini_tiktok/kitex_gen/favorite"
)

func ToKitexFavoriteActionRequest(req *hertz_favorite.FavoriteActionRequest) *kitex_favorite.FavoriteActionRequest {
	return &kitex_favorite.FavoriteActionRequest{
		Token:      req.Token,
		VideoId:    req.VideoID,
		ActionType: req.ActionType,
	}
}

func ToHertzFavoriteActionResponse(resp *kitex_favorite.FavoriteActionResponse) *hertz_favorite.FavoriteActionResponse {
	return &hertz_favorite.FavoriteActionResponse{
		StatusCode: resp.StatusCode,
		StatusMsg:  resp.StatusMsg,
	}
}

func ToKitexFavoriteListRequest(req *hertz_favorite.FavoriteListRequest) *kitex_favorite.FavoriteListRequest {
	return &kitex_favorite.FavoriteListRequest{
		UserId: req.UserID,
		Token:  req.Token,
	}
}

func ToHertzFavoriteListResponse(resp *kitex_favorite.FavoriteListResponse) *hertz_favorite.FavoriteListResponse {
	return &hertz_favorite.FavoriteListResponse{
		StatusCode: resp.StatusCode,
		StatusMsg:  resp.StatusMsg,
		VideoList:  ToHertzVideoList(resp.VideoList),
	}
}
