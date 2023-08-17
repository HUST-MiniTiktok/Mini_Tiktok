package service

import (
	"bytes"
	"context"
	"strconv"
	"time"

	"github.com/google/uuid"

	"github.com/HUST-MiniTiktok/mini_tiktok/kitex_gen/favorite"
	db "github.com/HUST-MiniTiktok/mini_tiktok/cmd/publish/dal/db"
	"github.com/HUST-MiniTiktok/mini_tiktok/cmd/publish/rpc"
	"github.com/HUST-MiniTiktok/mini_tiktok/conf"
	common "github.com/HUST-MiniTiktok/mini_tiktok/kitex_gen/common"
	publish "github.com/HUST-MiniTiktok/mini_tiktok/kitex_gen/publish"
	user "github.com/HUST-MiniTiktok/mini_tiktok/kitex_gen/user"
	"github.com/HUST-MiniTiktok/mini_tiktok/mw/ffmpeg"
	"github.com/HUST-MiniTiktok/mini_tiktok/mw/jwt"
	"github.com/HUST-MiniTiktok/mini_tiktok/mw/oss"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/remote/trans/nphttp2/codes"
)

var (
	VideoBucketName string
	ImageBucketName string
	Jwt *jwt.JWT
)

func init() {
	VideoBucketName = conf.GetConf().GetString("oss.videobucket")
	ImageBucketName = conf.GetConf().GetString("oss.imagebucket")
	Jwt = jwt.NewJWT()
}

type PublishService struct {
	ctx context.Context
}

func NewPublishService(ctx context.Context) *PublishService {
	return &PublishService{
		ctx: ctx,
	}
}

func (s *PublishService) PublishAction(request *publish.PublishActionRequest) (resp *publish.PublishActionResponse, err error) {
	klog.Infof("publish_action request")
	user_claims, err := Jwt.ExtractClaims(request.GetToken())
	author_id := user_claims.ID

	if err != nil {
		err_msg := err.Error()
		resp = &publish.PublishActionResponse{
			StatusCode: int32(codes.Unauthenticated),
			StatusMsg:  &err_msg,
		}
	}

	cover_data, err := ffmpeg.GetVideoCover(request.Data)
	if err != nil {
		err_msg := err.Error()
		resp = &publish.PublishActionResponse{StatusCode: int32(codes.Internal), StatusMsg: &err_msg}
		return
	}
	klog.Infof("cover_size=%v", strconv.FormatInt(int64(len(cover_data)), 10))

	video_buf := bytes.NewBuffer(request.Data)
	video_filename := uuid.NewString() + ".mp4"
	video_info, err := oss.PutToBucketWithBuf(s.ctx, VideoBucketName, video_filename, video_buf)
	if err != nil {
		err_msg := err.Error()
		resp = &publish.PublishActionResponse{StatusCode: int32(codes.Internal), StatusMsg: &err_msg}
		return
	}
	klog.Infof("upload_video_size=%v", strconv.FormatInt(video_info.Size, 10))

	cover_buf := bytes.NewBuffer(cover_data)
	cover_filename := uuid.NewString() + ".png"
	cover_info, err := oss.PutToBucketWithBuf(s.ctx, ImageBucketName, cover_filename, cover_buf)
	if err != nil {
		err_msg := err.Error()
		resp = &publish.PublishActionResponse{StatusCode: int32(codes.Internal), StatusMsg: &err_msg}
		return
	}
	klog.Infof("upload_cover_size=%v", strconv.FormatInt(cover_info.Size, 10))

	id, err := db.CreateVideo(s.ctx, &db.Video{
		AuthorID:    author_id,
		PlayURL:     oss.ToDbURL(s.ctx, VideoBucketName, video_filename),
		CoverURL:    oss.ToDbURL(s.ctx, ImageBucketName, cover_filename),
		PublishTime: time.Now(),
		Title:       request.Title,
	})
	if err != nil {
		err_msg := err.Error()
		resp = &publish.PublishActionResponse{StatusCode: int32(codes.Internal), StatusMsg: &err_msg}
		return
	}
	klog.Infof("db create video_id=%v", id)
	resp = &publish.PublishActionResponse{
		StatusCode: int32(codes.OK),
		StatusMsg:  nil,
	}
	return
}

func (s *PublishService) PublishList(request *publish.PublishListRequest) (resp *publish.PublishListResponse, err error) {
	klog.Infof("publish_list request")
	user_claims, err := Jwt.ExtractClaims(request.GetToken())
	query_user_id := request.UserId
	var curr_user_id int64
	if err != nil {
		curr_user_id = 0
	} else {
		curr_user_id = user_claims.ID
	}
	
	db_videos, err := db.GetVideoByAuthorId(s.ctx, query_user_id)
	if err != nil {
		err_msg := err.Error()
		resp = &publish.PublishListResponse{StatusCode: int32(codes.Internal), StatusMsg: &err_msg}
		return
	}

	err_chan := make(chan error)
	defer close(err_chan)
	video_chan := make(chan *common.Video)
	defer close(video_chan)

	var kitex_videos []*common.Video

	for _, db_video := range db_videos {
		go func(db_video *db.Video) {
			author, err := rpc.UserRPC.User(s.ctx, &user.UserRequest{UserId: db_video.AuthorID})
			if err != nil {
				err_chan <- err
				return
			} 
			favorite_count, err := rpc.FavoriteRPC.GetVideoFavoriteCount(s.ctx, &favorite.GetVideoFavoriteCountRequest{VideoId: db_video.ID})
			if err != nil {
				err_chan <- err
				return
			}
			is_favorite, err := rpc.FavoriteRPC.CheckIsFavorite(s.ctx, &favorite.CheckIsFavoriteRequest{VideoId: db_video.ID, UserId: curr_user_id})
			if err != nil {
				err_chan <- err
				return
			}

			klog.Infof("real_play_url=%v", oss.ToRealURL(s.ctx, db_video.PlayURL))
			klog.Infof("real_cover_url=%v", oss.ToRealURL(s.ctx, db_video.CoverURL))

			video_chan <- &common.Video{
				Id:       db_video.ID,
				Author:   author.User,
				PlayUrl:  oss.ToRealURL(s.ctx, db_video.PlayURL),
				CoverUrl: oss.ToRealURL(s.ctx, db_video.CoverURL),
				Title:    db_video.Title,
				FavoriteCount: favorite_count.FavoriteCount,
				IsFavorite: is_favorite.IsFavorite,
			}
		}(db_video)
	}

	for i := 0; i < len(db_videos); i++ {
		select {
		case err := <-err_chan:
			err_msg := err.Error()
			resp = &publish.PublishListResponse{StatusCode: int32(codes.Internal), StatusMsg: &err_msg}
			return resp, err
		case video := <-video_chan:
			kitex_videos = append(kitex_videos, video)
		}
	}

	resp = &publish.PublishListResponse{
		StatusCode: int32(codes.OK),
		StatusMsg:  nil,
		VideoList:  kitex_videos,
	}
	return
}

func (s *PublishService) GetVideoById(request *publish.GetVideoByIdRequest) (resp *publish.GetVideoByIdResponse, err error) {
	klog.Infof("get_video_by_id request")
	db_video, err := db.GetVideoById(s.ctx, request.Id)
	if err != nil {
		err_msg := err.Error()
		resp = &publish.GetVideoByIdResponse{StatusCode: int32(codes.Internal), StatusMsg: &err_msg}
		return
	}
	author, err := rpc.UserRPC.User(s.ctx, &user.UserRequest{UserId: db_video.AuthorID})
	if err != nil {
		err_msg := err.Error()
		resp = &publish.GetVideoByIdResponse{StatusCode: int32(codes.Internal), StatusMsg: &err_msg}
		return
	}

	kitex_video := &common.Video{
		Id:       db_video.ID,
		Author:   author.User,
		PlayUrl:  oss.ToRealURL(s.ctx, db_video.PlayURL),
		CoverUrl: oss.ToRealURL(s.ctx, db_video.CoverURL),
		Title:    db_video.Title,
	}
	resp = &publish.GetVideoByIdResponse{
		StatusCode: int32(codes.OK),
		StatusMsg:  nil,
		Video:      kitex_video,
	}
	return
}

func (s *PublishService) GetVideoByIdList(request *publish.GetVideoByIdListRequest) (resp *publish.GetVideoByIdListResponse, err error) {
	db_videos, err := db.GetVideosByIDs(s.ctx, request.Id)
	if err != nil {
		err_msg := err.Error()
		resp = &publish.GetVideoByIdListResponse{StatusCode: int32(codes.Internal), StatusMsg: &err_msg}
		return
	}
	err_chan := make(chan error)
	defer close(err_chan)
	video_chan := make(chan *common.Video)
	defer close(video_chan)
	var kitex_videos []*common.Video

	for _, db_video := range db_videos {
		go func(db_video *db.Video) {
			author, err := rpc.UserRPC.User(s.ctx, &user.UserRequest{UserId: db_video.AuthorID})
			kitex_author := author.User
			if err != nil {
				err_chan <- err
			} else {
				video_chan <- &common.Video{
					Id:       db_video.ID,
					Author:   kitex_author,
					PlayUrl:  oss.ToRealURL(s.ctx, db_video.PlayURL),
					CoverUrl: oss.ToRealURL(s.ctx, db_video.CoverURL),
					Title:    db_video.Title,
				}
			}
		}(db_video)
	}

	for i := 0; i < len(db_videos); i++ {
		select {
		case err := <-err_chan:
			err_msg := err.Error()
			resp = &publish.GetVideoByIdListResponse{StatusCode: int32(codes.Internal), StatusMsg: &err_msg}
			return resp, err
		case video := <-video_chan:
			kitex_videos = append(kitex_videos, video)
		}
	}

	resp = &publish.GetVideoByIdListResponse{
		StatusCode: int32(codes.OK),
		StatusMsg:  nil,
		VideoList:  kitex_videos,
	}
	return
}
