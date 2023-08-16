package service

import (
	"context"

	db "github.com/HUST-MiniTiktok/mini_tiktok/cmd/favorite/dal/db"
	favorite "github.com/HUST-MiniTiktok/mini_tiktok/cmd/favorite/kitex_gen/favorite"
	publish "github.com/HUST-MiniTiktok/mini_tiktok/cmd/publish/kitex_gen/publish"
	"github.com/HUST-MiniTiktok/mini_tiktok/mw/jwt"
	rpc "github.com/HUST-MiniTiktok/mini_tiktok/rpc"
	"github.com/cloudwego/kitex/pkg/remote/trans/nphttp2/codes"
)

type FavoriteService struct {
	ctx context.Context
}

func NewFavoriteService(ctx context.Context) *FavoriteService {
	return &FavoriteService{ctx: ctx}
}

func (s *FavoriteService) FavoriteAction(ctx context.Context, req *favorite.FavoriteActionRequest) (resp *favorite.FavoriteActionResponse, err error) {

	claim, err := jwt.Jwt.ExtractClaims(req.Token)
	if err != nil {
		msg := err.Error()
		return &favorite.FavoriteActionResponse{
			// status_code = 7 表示鉴权失败
			StatusCode: int32(codes.PermissionDenied),
			StatusMsg:  &msg}, nil
	}

	// 类型是点赞请求
	if req.ActionType == 1 {
		if _, err = db.NewFavorite(ctx, claim.ID, req.VideoId); err != nil {
			msg := err.Error()
			return &favorite.FavoriteActionResponse{
				//点赞失败
				StatusCode: int32(codes.Internal),
				StatusMsg:  &msg}, nil
		}
	} else if req.ActionType == 2 { //取消点赞
		if _, err = db.CancelFavorite(ctx, claim.ID, req.VideoId); err != nil {
			msg := err.Error()
			return &favorite.FavoriteActionResponse{
				//取消点赞失败
				StatusCode: int32(codes.Internal),
				StatusMsg:  &msg}, nil
		}
	} else {
		msg := "action_type error"
		return &favorite.FavoriteActionResponse{
			// status_code = 3 表示参数错误
			StatusCode: int32(codes.InvalidArgument),
			StatusMsg:  &msg}, nil
	}

	return &favorite.FavoriteActionResponse{
		// status_code = 0 表示操作成功
		StatusCode: int32(codes.OK),
		StatusMsg:  nil}, nil
}

func (s *FavoriteService) FavoriteList(ctx context.Context, req *favorite.FavoriteListRequest) (resp *favorite.FavoriteListResponse, err error) {

	claim, err := jwt.Jwt.ExtractClaims(req.Token)
	if err != nil {
		msg := err.Error()
		return &favorite.FavoriteListResponse{
			// status_code = 7 表示鉴权失败
			StatusCode: int32(codes.PermissionDenied),
			StatusMsg:  &msg,
			VideoList:  nil}, nil
	}

	if req.UserId == 0 || claim.ID == 0 {
		msg := "request ID = 0"
		return &favorite.FavoriteListResponse{
			//参数错误
			StatusCode: int32(codes.InvalidArgument),
			StatusMsg:  &msg,
			VideoList:  nil}, nil
	}

	_, videoIDList, err := db.GetFavoriteList(ctx, req.UserId)
	if err != nil {
		msg := "GetFavoriteList Failed"
		return &favorite.FavoriteListResponse{
			// 获取失败
			StatusCode: int32(codes.Internal),
			StatusMsg:  &msg,
			VideoList:  nil}, nil
	}

	err_chan := make(chan error)
	video_chan := make(chan []*publish.Video)
	var videoList []*favorite.Video

	videosResponse, err := rpc.GetVideoByIdList(ctx, &publish.GetVideoByIdListRequest{Id: videoIDList})
	if err != nil {
		err_chan <- err
	} else {
		video_chan <- videosResponse.VideoList
	}

	select {
	case err := <-err_chan:
		err_msg := err.Error()
		resp = &favorite.FavoriteListResponse{
			StatusCode: int32(codes.Internal),
			StatusMsg:  &err_msg,
			VideoList:  nil}
		return resp, err
	case PBvideoList := <-video_chan:
		for _, video := range PBvideoList {
			videoList = append(videoList, PubVideoToFavVideo(video))
		}

	}

	return &favorite.FavoriteListResponse{
		StatusCode: int32(codes.OK),
		StatusMsg:  nil,
		VideoList:  videoList}, nil
}

func PubVideoToFavVideo(video *publish.Video) *favorite.Video {
	return &favorite.Video{
		Id:            video.Id,
		Author:        PubUserToFavUser(video.Author),
		PlayUrl:       video.PlayUrl,
		CoverUrl:      video.CoverUrl,
		FavoriteCount: video.FavoriteCount,
		CommentCount:  video.CommentCount,
		IsFavorite:    video.IsFavorite,
		Title:         video.Title,
	}
}

func PubUserToFavUser(user *publish.User) *favorite.User {
	return &favorite.User{
		Id:              user.Id,
		Name:            user.Name,
		FollowCount:     user.FollowCount,
		FollowerCount:   user.FollowerCount,
		IsFollow:        user.IsFollow,
		Avatar:          user.Avatar,
		BackgroundImage: user.BackgroundImage,
		Signature:       user.Signature,
		TotalFavorited:  user.TotalFavorited,
		WorkCount:       user.WorkCount,
		FavoriteCount:   user.FavoriteCount,
	}
}

// type FavoriteListRequest struct {
// 	UserId int64  `thrift:"user_id,1" frugal:"1,default,i64" json:"user_id"`
// 	Token  string `thrift:"token,2" frugal:"2,default,string" json:"token"`
// }

// type FavoriteListResponse struct {
// 	StatusCode int32    `thrift:"status_code,1" frugal:"1,default,i32" json:"status_code"`
// 	StatusMsg  *string  `thrift:"status_msg,2,optional" frugal:"2,optional,string" json:"status_msg,omitempty"`
// 	VideoList  []*Video `thrift:"video_list,3" frugal:"3,default,list<Video>" json:"video_list"`
// }

// type FavoriteActionRequest struct {
// 	Token      string `thrift:"token,1" frugal:"1,default,string" json:"token"`
// 	VideoId    int64  `thrift:"video_id,2" frugal:"2,default,i64" json:"video_id"`
// 	ActionType int32  `thrift:"action_type,3" frugal:"3,default,i32" json:"action_type"`
// }

// type FavoriteActionResponse struct {
// 	StatusCode int32   `thrift:"status_code,1" frugal:"1,default,i32" json:"status_code"`
// 	StatusMsg  *string `thrift:"status_msg,2,optional" frugal:"2,optional,string" json:"status_msg,omitempty"`
// }

// type GetVideoByIdListRequest struct {
// 	Id []int64 `thrift:"id,1" frugal:"1,default,list<i64>" json:"id"`
// }

// type GetVideoByIdListResponse struct {
// 	StatusCode int32    `thrift:"status_code,1" frugal:"1,default,i32" json:"status_code"`
// 	StatusMsg  *string  `thrift:"status_msg,2,optional" frugal:"2,optional,string" json:"status_msg,omitempty"`
// 	VideoList  []*Video `thrift:"video_list,3" frugal:"3,default,list<Video>" json:"video_list"`
// }
