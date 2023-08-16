package service

import (
	"time"
	"bytes"
	"context"
	"strconv"
	"github.com/google/uuid"

	db "github.com/HUST-MiniTiktok/mini_tiktok/cmd/publish/dal/db"
	oss "github.com/HUST-MiniTiktok/mini_tiktok/cmd/publish/dal/oss"
	publish "github.com/HUST-MiniTiktok/mini_tiktok/cmd/publish/kitex_gen/publish"
	"github.com/HUST-MiniTiktok/mini_tiktok/conf"
	"github.com/HUST-MiniTiktok/mini_tiktok/mw/ffmpeg"
	"github.com/HUST-MiniTiktok/mini_tiktok/mw/jwt"
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
