package service

import (
	"context"

	db "github.com/HUST-MiniTiktok/mini_tiktok/cmd/favorite/dal/db"
	"github.com/HUST-MiniTiktok/mini_tiktok/cmd/favorite/rpc"
	common "github.com/HUST-MiniTiktok/mini_tiktok/kitex_gen/common"
	favorite "github.com/HUST-MiniTiktok/mini_tiktok/kitex_gen/favorite"
	publish "github.com/HUST-MiniTiktok/mini_tiktok/kitex_gen/publish"
	"github.com/HUST-MiniTiktok/mini_tiktok/mw/jwt"
	"github.com/cloudwego/kitex/pkg/remote/trans/nphttp2/codes"
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
		msg := err.Error()
		return &favorite.FavoriteActionResponse{
			// status_code = 7 表示鉴权失败
			StatusCode: int32(codes.Unauthenticated),
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

	claim, err := Jwt.ExtractClaims(req.Token)
	if err != nil {
		msg := err.Error()
		return &favorite.FavoriteListResponse{
			// status_code = 7 表示鉴权失败
			StatusCode: int32(codes.PermissionDenied),
			StatusMsg:  &msg,
			VideoList:  nil}, err
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
			VideoList:  nil}, err
	}

	err_chan := make(chan error)
	defer close(err_chan)
	video_chan := make(chan []*common.Video)
	defer close(video_chan)

	videosResponse, err := rpc.PublishRPC.GetVideoByIdList(ctx, &publish.GetVideoByIdListRequest{Id: videoIDList})
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
		return &favorite.FavoriteListResponse{
			StatusCode: int32(codes.OK),
			StatusMsg:  nil,
			VideoList:  PBvideoList}, nil
	}

}

func (s *FavoriteService) CheckIsFavorite(ctx context.Context, req *favorite.CheckIsFavoriteRequest) (resp *favorite.CheckIsFavoriteResponse, err error) {

	exist, err := db.CheckFavorite(ctx, req.UserId, req.VideoId)
	if err != nil {
		msg := "check favorite Failed"
		return &favorite.CheckIsFavoriteResponse{
			// 查找失败
			StatusCode: int32(codes.Internal),
			StatusMsg:  &msg,
			IsFavorite: false}, err
	}
	return &favorite.CheckIsFavoriteResponse{
		StatusCode: int32(codes.OK),
		StatusMsg:  nil,
		IsFavorite: exist}, nil
}

func (s *FavoriteService) GetVideoFavoriteCount(ctx context.Context, req *favorite.GetVideoFavoriteCountRequest) (resp *favorite.GetVideoFavoriteCountResponse, err error) {

	count, err := db.VideoFavoriteCount(ctx, req.VideoId)
	if err != nil {
		msg := "check video favorite Failed"
		return &favorite.GetVideoFavoriteCountResponse{
			// 查找失败
			StatusCode:    int32(codes.Internal),
			StatusMsg:     &msg,
			FavoriteCount: 0}, err
	}
	return &favorite.GetVideoFavoriteCountResponse{
		StatusCode:    int32(codes.OK),
		StatusMsg:     nil,
		FavoriteCount: count}, nil
}
