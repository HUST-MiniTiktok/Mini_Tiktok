package service

import (
	"bytes"
	"context"
	"strconv"
	"time"

	"github.com/google/uuid"

	db "github.com/HUST-MiniTiktok/mini_tiktok/cmd/publish/dal/db"
	"github.com/HUST-MiniTiktok/mini_tiktok/cmd/publish/pack"
	publish "github.com/HUST-MiniTiktok/mini_tiktok/kitex_gen/publish"
	"github.com/HUST-MiniTiktok/mini_tiktok/pkg/conf"
	"github.com/HUST-MiniTiktok/mini_tiktok/pkg/errno"
	"github.com/HUST-MiniTiktok/mini_tiktok/pkg/mw/ffmpeg"
	"github.com/HUST-MiniTiktok/mini_tiktok/pkg/mw/jwt"
	"github.com/HUST-MiniTiktok/mini_tiktok/pkg/mw/oss"
	"github.com/cloudwego/kitex/pkg/klog"
)

var (
	VideoBucketName string
	ImageBucketName string
	Jwt             *jwt.JWT
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

// PublishAction: publish a video
func (s *PublishService) PublishAction(request *publish.PublishActionRequest) (resp *publish.PublishActionResponse, err error) {
	claim, err := Jwt.ExtractClaims(request.Token)
	author_id := claim.ID

	if err != nil || claim.ID == 0 {
		return pack.NewPublishActionResponse(errno.AuthorizationFailedErr), err
	}
	// generate filenames for video and cover
	cover_filename := uuid.NewString() + ".png"
	video_filename := uuid.NewString() + ".mp4"
	// concurrent upload video and cover
	err_chan := make(chan error)
	ok := make(chan bool)
	go func() {
		// get cover data
		cover_data, err := ffmpeg.GetVideoCover(request.Data)
		if err != nil {
			err_chan <- err
			return
		}
		klog.Infof("cover_size=%v", strconv.FormatInt(int64(len(cover_data)), 10))
		// upload cover
		cover_buf := bytes.NewBuffer(cover_data)
		cover_info, err := oss.PutToBucketWithBuf(s.ctx, ImageBucketName, cover_filename, cover_buf)
		if err != nil {
			err_chan <- err
			return
		}
		klog.Infof("upload_cover_size=%v", strconv.FormatInt(cover_info.Size, 10))
		ok <- true
	}()
	go func() {
		// upload video
		video_buf := bytes.NewBuffer(request.Data)
		video_info, err := oss.PutToBucketWithBuf(s.ctx, VideoBucketName, video_filename, video_buf)
		if err != nil {
			err_chan <- err
			return
		}
		klog.Infof("upload_video_size=%v", strconv.FormatInt(video_info.Size, 10))
		ok <- true
	}()

	for i := 0; i < 2; i++ {
		select {
		case err := <-err_chan:
			return pack.NewPublishActionResponse(err), err
		case <-ok:
		}
	}

	_, err = db.CreateVideo(s.ctx, &db.Video{
		AuthorID:    author_id,
		PlayURL:     oss.ToDbURL(VideoBucketName, video_filename),
		CoverURL:    oss.ToDbURL(ImageBucketName, cover_filename),
		PublishTime: time.Now(),
		Title:       request.Title,
	})
	if err != nil {
		return pack.NewPublishActionResponse(err), err
	}

	return pack.NewPublishActionResponse(errno.Success), nil
}

// PublishList: get a list of videos published by a user
func (s *PublishService) PublishList(request *publish.PublishListRequest) (resp *publish.PublishListResponse, err error) {
	user_claims, err := Jwt.ExtractClaims(request.Token)
	query_user_id := request.UserId
	var curr_user_id int64
	if err != nil {
		curr_user_id = 0
	} else {
		curr_user_id = user_claims.ID
	}

	db_videos, err := db.GetVideoByAuthorId(s.ctx, query_user_id)
	if err != nil {
		return pack.NewPublishListResponse(err), err
	}

	kitex_videos, err := pack.ToKitexVideoList(s.ctx, curr_user_id, request.Token, db_videos)
	if err != nil {
		return pack.NewPublishListResponse(err), err
	}

	resp = pack.NewPublishListResponse(errno.Success)
	resp.VideoList = kitex_videos
	return
}

func (s *PublishService) GetVideoById(request *publish.GetVideoByIdRequest) (resp *publish.GetVideoByIdResponse, err error) {
	user_claims, err := Jwt.ExtractClaims(request.Token)
	var curr_user_id int64
	if err != nil {
		curr_user_id = 0
	} else {
		curr_user_id = user_claims.ID
	}

	db_video, err := db.GetVideoById(s.ctx, request.Id)
	if err != nil {
		return pack.NewGetVideoByIdResponse(err), err
	}

	kitex_video, err := pack.ToKitexVideo(s.ctx, curr_user_id, request.Token, db_video)
	if err != nil {
		return pack.NewGetVideoByIdResponse(err), err
	}

	resp = pack.NewGetVideoByIdResponse(errno.Success)
	resp.Video = kitex_video
	return resp, nil
}

func (s *PublishService) GetVideoByIdList(request *publish.GetVideoByIdListRequest) (resp *publish.GetVideoByIdListResponse, err error) {
	user_claims, err := Jwt.ExtractClaims(request.Token)
	var curr_user_id int64
	if err != nil {
		curr_user_id = 0
	} else {
		curr_user_id = user_claims.ID
	}

	db_videos, err := db.GetVideosByIdList(s.ctx, request.Id)
	if err != nil {
		return pack.NewGetVideoByIdListResponse(err), err
	}

	kitex_videos, err := pack.ToKitexVideoList(s.ctx, curr_user_id, request.Token, db_videos)
	if err != nil {
		return pack.NewGetVideoByIdListResponse(err), err
	}

	resp = pack.NewGetVideoByIdListResponse(errno.Success)
	resp.VideoList = kitex_videos
	return resp, nil
}

// GetPublishInfoByUserId: get user work count and video id list
func (s *PublishService) GetPublishInfoByUserId(request *publish.GetPublishInfoByUserIdRequest) (resp *publish.GetPublishInfoByUserIdResponse, err error) {
	video_ids, err := db.GetVideoIdListByAuthorId(s.ctx, request.UserId)
	if err != nil {
		return pack.NewGetPublishInfoByUserIdResponse(err), err
	}

	resp = pack.NewGetPublishInfoByUserIdResponse(errno.Success)
	resp.WorkCount = int64(len(video_ids))
	resp.VideoIds = video_ids
	return
}
