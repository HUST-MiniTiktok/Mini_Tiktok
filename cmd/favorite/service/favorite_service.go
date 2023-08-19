package service

import (
	"context"

	db "github.com/HUST-MiniTiktok/mini_tiktok/cmd/favorite/dal/db"
	"github.com/HUST-MiniTiktok/mini_tiktok/cmd/favorite/pack"
	"github.com/HUST-MiniTiktok/mini_tiktok/cmd/favorite/rpc"
	favorite "github.com/HUST-MiniTiktok/mini_tiktok/kitex_gen/favorite"
	publish "github.com/HUST-MiniTiktok/mini_tiktok/kitex_gen/publish"
	"github.com/HUST-MiniTiktok/mini_tiktok/pkg/errno"
	"github.com/HUST-MiniTiktok/mini_tiktok/pkg/mw/jwt"
)

var (
	Jwt *jwt.JWT
)

func init() {
	Jwt = jwt.NewJWT()
}

type FavoriteService struct {
	ctx context.Context
}

func NewFavoriteService(ctx context.Context) *FavoriteService {
	return &FavoriteService{ctx: ctx}
}

func (s *FavoriteService) FavoriteAction(ctx context.Context, req *favorite.FavoriteActionRequest) (resp *favorite.FavoriteActionResponse, err error) {
	claim, err := Jwt.ExtractClaims(req.Token)
	if err != nil {
		return pack.NewFavoriteActionResponse(errno.AuthorizationFailedErr), err
	}

	// 类型是点赞请求
	if req.ActionType == 1 {
		if _, err = db.NewFavorite(ctx, claim.ID, req.VideoId); err != nil {
			return pack.NewFavoriteActionResponse(err), err
		}
	} else if req.ActionType == 2 { //取消点赞
		if _, err = db.CancelFavorite(ctx, claim.ID, req.VideoId); err != nil {
			return pack.NewFavoriteActionResponse(err), err
		}
	} else {
		return pack.NewFavoriteActionResponse(errno.ParamErr), errno.ParamErr
	}

	return pack.NewFavoriteActionResponse(errno.Success), nil
}

func (s *FavoriteService) FavoriteList(ctx context.Context, req *favorite.FavoriteListRequest) (resp *favorite.FavoriteListResponse, err error) {
	_, err = Jwt.ExtractClaims(req.Token)
	if err != nil {
		return pack.NewFavoriteListResponse(errno.AuthorizationFailedErr), err
	}

	_, videoIDList, err := db.GetFavoriteList(ctx, req.UserId)
	if err != nil {
		return pack.NewFavoriteListResponse(err), err
	}

	videosResponse, err := rpc.PublishRPC.GetVideoByIdList(ctx, &publish.GetVideoByIdListRequest{Id: videoIDList, Token: req.Token})
	if err != nil {
		return pack.NewFavoriteListResponse(err), err
	} else {
		resp = pack.NewFavoriteListResponse(errno.Success)
		resp.VideoList = videosResponse.VideoList
		return resp, nil
	}
}

func (s *FavoriteService) CheckIsFavorite(ctx context.Context, req *favorite.CheckIsFavoriteRequest) (resp *favorite.CheckIsFavoriteResponse, err error) {
	is_exist, err := db.CheckFavorite(ctx, req.UserId, req.VideoId)
	if err != nil {
		return pack.NewCheckIsFavoriteResponse(err), err
	}
	resp = pack.NewCheckIsFavoriteResponse(errno.Success)
	resp.IsFavorite = is_exist
	return resp, nil
}

func (s *FavoriteService) GetVideoFavoriteCount(ctx context.Context, req *favorite.GetVideoFavoriteCountRequest) (resp *favorite.GetVideoFavoriteCountResponse, err error) {
	count, err := db.VideoFavoriteCount(ctx, req.VideoId)
	if err != nil {
		return pack.NewGetVideoFavoriteCountResponse(err), err
	}
	resp = pack.NewGetVideoFavoriteCountResponse(errno.Success)
	resp.FavoriteCount = count
	return resp, nil
}
