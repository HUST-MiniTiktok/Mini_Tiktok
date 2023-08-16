package service

import (
	"bytes"
	"context"
	"strconv"
	"time"

	"github.com/google/uuid"

	db "github.com/HUST-MiniTiktok/mini_tiktok/cmd/publish/dal/db"
	oss "github.com/HUST-MiniTiktok/mini_tiktok/cmd/publish/dal/oss"
	publish "github.com/HUST-MiniTiktok/mini_tiktok/cmd/publish/kitex_gen/publish"
	user "github.com/HUST-MiniTiktok/mini_tiktok/cmd/user/kitex_gen/user"
	"github.com/HUST-MiniTiktok/mini_tiktok/conf"
	"github.com/HUST-MiniTiktok/mini_tiktok/mw/ffmpeg"
	"github.com/HUST-MiniTiktok/mini_tiktok/mw/jwt"
	rpc "github.com/HUST-MiniTiktok/mini_tiktok/rpc"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/remote/trans/nphttp2/codes"
)

var (
	VideoBucketName string
	ImageBucketName string
)

func init() {
	VideoBucketName = conf.GetConf().GetString("oss.videobucket")
	ImageBucketName = conf.GetConf().GetString("oss.imagebucket")
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
	user_claims, err := jwt.Jwt.ExtractClaims(request.GetToken())
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
		PlayURL:     VideoBucketName + "/" + video_filename,
		CoverURL:    ImageBucketName + "/" + cover_filename,
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
	user_claims, err := jwt.Jwt.ExtractClaims(request.GetToken())
	curr_user_id := user_claims.ID
	query_user_id := request.UserId
	if err != nil {
		curr_user_id = 0
	}
	_ = curr_user_id
	db_videos, err := db.GetVideoByAuthorId(s.ctx, query_user_id)
	if err != nil {
		err_msg := err.Error()
		resp = &publish.PublishListResponse{StatusCode: int32(codes.Internal), StatusMsg: &err_msg}
		return
	}

	err_chan := make(chan error)
	video_chan := make(chan *publish.Video)
	var kitex_videos []*publish.Video

	for _, db_video := range db_videos {
		go func(db_video *db.Video) {
			author, err := rpc.User(s.ctx, &user.UserRequest{UserId: db_video.AuthorID})
			kitex_author := author.User
			if err != nil {
				err_chan <- err
			} else {
				video_chan <- &publish.Video{
					Id:       db_video.ID,
					Author:   (*publish.User)(kitex_author),
					PlayUrl:  db_video.PlayURL,
					CoverUrl: db_video.CoverURL,
					Title:    db_video.Title,
				}
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
	db_video, err := db.GetVideoById(s.ctx, request.Id)
	if err != nil {
		err_msg := err.Error()
		resp = &publish.GetVideoByIdResponse{StatusCode: int32(codes.Internal), StatusMsg: &err_msg}
		return
	}
	author, err := rpc.User(s.ctx, &user.UserRequest{UserId: db_video.AuthorID})
	kitex_author := author.User
	if err != nil {
		err_msg := err.Error()
		resp = &publish.GetVideoByIdResponse{StatusCode: int32(codes.Internal), StatusMsg: &err_msg}
		return
	}

	kitex_video := &publish.Video{
		Id:       db_video.ID,
		Author:   (*publish.User)(kitex_author),
		PlayUrl:  db_video.PlayURL,
		CoverUrl: db_video.CoverURL,
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
	video_chan := make(chan *publish.Video)
	var kitex_videos []*publish.Video

	for _, db_video := range db_videos {
		go func(db_video *db.Video) {
			author, err := rpc.User(s.ctx, &user.UserRequest{UserId: db_video.AuthorID})
			kitex_author := author.User
			if err != nil {
				err_chan <- err
			} else {
				video_chan <- &publish.Video{
					Id:       db_video.ID,
					Author:   (*publish.User)(kitex_author),
					PlayUrl:  db_video.PlayURL,
					CoverUrl: db_video.CoverURL,
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
