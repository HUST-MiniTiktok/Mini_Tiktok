package service

import (
	"context"

	"github.com/HUST-MiniTiktok/mini_tiktok/cmd/favorite/client"
	db "github.com/HUST-MiniTiktok/mini_tiktok/cmd/favorite/dal/db"
	"github.com/HUST-MiniTiktok/mini_tiktok/cmd/favorite/pack"
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

// FavoriteAction: favorite a video or cancel favorite
func (s *FavoriteService) FavoriteAction(req *favorite.FavoriteActionRequest) (resp *favorite.FavoriteActionResponse, err error) {
	claim, err := Jwt.ExtractClaims(req.Token)
	if err != nil {
		return pack.NewFavoriteActionResponse(errno.AuthorizationFailedErr), err
	}

	// Action type: 1: 点赞 2: 取消点赞
	if req.ActionType == 1 {
		if _, err = db.NewFavorite(s.ctx, claim.ID, req.VideoId); err != nil {
			return pack.NewFavoriteActionResponse(err), err
		}
	} else if req.ActionType == 2 {
		if _, err = db.CancelFavorite(s.ctx, claim.ID, req.VideoId); err != nil {
			return pack.NewFavoriteActionResponse(err), err
		}
	} else {
		return pack.NewFavoriteActionResponse(errno.ParamErr), errno.ParamErr
	}

	return pack.NewFavoriteActionResponse(errno.Success), nil
}

// FavoriteList: get favorite list
func (s *FavoriteService) FavoriteList(req *favorite.FavoriteListRequest) (resp *favorite.FavoriteListResponse, err error) {
	_, err = Jwt.ExtractClaims(req.Token)
	if err != nil {
		return pack.NewFavoriteListResponse(errno.AuthorizationFailedErr), err
	}

	_, videoIDList, err := db.GetFavoriteList(s.ctx, req.UserId)
	if err != nil {
		return pack.NewFavoriteListResponse(err), err
	}

	videosResponse, err := client.PublishRPC.GetVideoByIdList(s.ctx, &publish.GetVideoByIdListRequest{Id: videoIDList, Token: req.Token})
	if err != nil {
		return pack.NewFavoriteListResponse(err), err
	} else {
		resp = pack.NewFavoriteListResponse(errno.Success)
		resp.VideoList = videosResponse.VideoList
		return resp, nil
	}
}

// GetVideoFavoriteInfo: get favoriteCount and isFavorite of a video
func (s *FavoriteService) GetVideoFavoriteInfo(req *favorite.GetVideoFavoriteInfoRequest) (resp *favorite.GetVideoFavoriteInfoResponse, err error) {
	count, err := db.VideoFavoriteCount(s.ctx, req.VideoId)
	if err != nil {
		return pack.NewGetVideoFavoriteInfoResponse(err), err
	}

	is_favorite, err := db.CheckFavorite(s.ctx, req.UserId, req.VideoId)
	if err != nil {
		return pack.NewGetVideoFavoriteInfoResponse(err), err
	}

	resp = pack.NewGetVideoFavoriteInfoResponse(errno.Success)
	resp.FavoriteCount = count
	resp.IsFavorite = is_favorite
	return resp, nil
}

// GetUserFavoriteInfo: get totalFavorited and favoriteCount of a user
func (s *FavoriteService) GetUserFavoriteInfo(req *favorite.GetUserFavoriteInfoRequest) (resp *favorite.GetUserFavoriteInfoResponse, err error) {
	publishInfoResp, err := client.PublishRPC.GetPublishInfoByUserId(s.ctx, &publish.GetPublishInfoByUserIdRequest{UserId: req.UserId})
	if err != nil {
		return pack.NewGetUserFavoriteInfoResponse(err), err
	}
	user_work_ids := publishInfoResp.VideoIds
	var favorite_count int64 = 0
	var favorited_count int64 = 0
	// Get favorited count, concurrently call db.VideoFavoriteCount
	errChan := make(chan error)
	countChan := make(chan int64)
	for _, video_id := range user_work_ids {
		go func(video_id int64) {
			count, err := db.VideoFavoriteCount(s.ctx, video_id)
			if err != nil {
				errChan <- err
			} else {
				countChan <- count
			}
		}(video_id)
	}
	for i := 0; i < len(user_work_ids); i++ {
		select {
		case err := <-errChan:
			return pack.NewGetUserFavoriteInfoResponse(err), err
		case count := <-countChan:
			favorited_count += count
		}
	}
	// get favorite count
	favorite_count, err = db.UserFavoriteCount(s.ctx, req.UserId)
	if err != nil {
		return pack.NewGetUserFavoriteInfoResponse(err), err
	}
	resp = pack.NewGetUserFavoriteInfoResponse(errno.Success)
	resp.TotalFavorited = favorited_count
	resp.FavoriteCount = favorite_count
	return resp, nil
}
