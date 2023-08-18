package service

import (
	"bytes"
	"context"
	"strconv"
	"time"

	"github.com/google/uuid"

	db "github.com/HUST-MiniTiktok/mini_tiktok/cmd/publish/dal/db"
	"github.com/HUST-MiniTiktok/mini_tiktok/cmd/publish/rpc"
	"github.com/HUST-MiniTiktok/mini_tiktok/conf"
	comment "github.com/HUST-MiniTiktok/mini_tiktok/kitex_gen/comment"
	common "github.com/HUST-MiniTiktok/mini_tiktok/kitex_gen/common"
	favorite "github.com/HUST-MiniTiktok/mini_tiktok/kitex_gen/favorite"
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

func (s *PublishService) PublishAction(request *publish.PublishActionRequest) (resp *publish.PublishActionResponse, err error) {
	klog.Infof("publish_action request")
	user_claims, err := Jwt.ExtractClaims(request.Token)
	author_id := user_claims.ID

	if err != nil {
		err_msg := err.Error()
		return &publish.PublishActionResponse{StatusCode: int32(codes.PermissionDenied), StatusMsg: &err_msg}, err
	}

	cover_filename := uuid.NewString() + ".png"
	video_filename := uuid.NewString() + ".mp4"
	err_chan := make(chan error)
	ok := make(chan bool)
	go func() {
		cover_data, err := ffmpeg.GetVideoCover(request.Data)
		if err != nil {
			err_chan <- err
			return
		}
		klog.Infof("cover_size=%v", strconv.FormatInt(int64(len(cover_data)), 10))

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
			err_msg := err.Error()
			return &publish.PublishActionResponse{StatusCode: int32(codes.Internal), StatusMsg: &err_msg}, err
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
		err_msg := err.Error()
		return &publish.PublishActionResponse{StatusCode: int32(codes.Internal), StatusMsg: &err_msg}, err
	}
	return &publish.PublishActionResponse{StatusCode: int32(codes.OK),StatusMsg: nil}, nil
}

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
		err_msg := err.Error()
		return &publish.PublishListResponse{StatusCode: int32(codes.Internal), StatusMsg: &err_msg}, err
	}

	err_chan := make(chan error)
	video_chan := make(chan *common.Video)
	kitex_videos := make([]*common.Video, 0, len(db_videos))

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
			comment_count, err := rpc.CommentRPC.GetVideoCommentCount(s.ctx, &comment.GetVideoCommentCountRequest{VideoId: db_video.ID})
			if err != nil {
				err_chan <- err
				return
			}

			video_chan <- &common.Video{
				Id:            db_video.ID,
				Author:        author.User,
				PlayUrl:       oss.ToRealURL(s.ctx, db_video.PlayURL),
				CoverUrl:      oss.ToRealURL(s.ctx, db_video.CoverURL),
				Title:         db_video.Title,
				FavoriteCount: favorite_count.FavoriteCount,
				IsFavorite:    is_favorite.IsFavorite,
				CommentCount:  comment_count.CommentCount,
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

	resp = &publish.PublishListResponse {
		StatusCode: int32(codes.OK),
		StatusMsg:  nil,
		VideoList:  kitex_videos,
	}
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
		errMsg := err.Error()
		return &publish.GetVideoByIdResponse{
			StatusCode: int32(codes.Internal),
			StatusMsg:  &errMsg,
		}, err
	}

	authorChan := make(chan *user.UserResponse)
	favoriteCountChan := make(chan *favorite.GetVideoFavoriteCountResponse)
	isFavoriteChan := make(chan *favorite.CheckIsFavoriteResponse)
	commentCountChan := make(chan *comment.GetVideoCommentCountResponse)

	go func() {
		author, err := rpc.UserRPC.User(s.ctx, &user.UserRequest{UserId: db_video.AuthorID})
		if err != nil {
			klog.Errorf("RPC User Error: %v", err)
			return
		}
		authorChan <- author
	}()

	go func() {
		favorite_count, err := rpc.FavoriteRPC.GetVideoFavoriteCount(s.ctx, &favorite.GetVideoFavoriteCountRequest{VideoId: db_video.ID})
		if err != nil {
			klog.Errorf("RPC GetVideoFavoriteCount Error: %v", err)
			return
		}
		favoriteCountChan <- favorite_count
	}()

	go func() {
		is_favorite, err := rpc.FavoriteRPC.CheckIsFavorite(s.ctx, &favorite.CheckIsFavoriteRequest{VideoId: db_video.ID, UserId: curr_user_id})
		if err != nil {
			klog.Errorf("RPC CheckIsFavorite Error: %v", err)
			return
		}
		isFavoriteChan <- is_favorite
	}()

	go func() {
		comment_count, err := rpc.CommentRPC.GetVideoCommentCount(s.ctx, &comment.GetVideoCommentCountRequest{VideoId: db_video.ID})
		if err != nil {
			klog.Errorf("RPC GetVideoCommentCount Error: %v", err)
			return
		}
		commentCountChan <- comment_count
	}()

	author := <-authorChan
	favorite_count := <-favoriteCountChan
	is_favorite := <-isFavoriteChan
	comment_count := <-commentCountChan

	kitex_video := &common.Video{
		Id:            db_video.ID,
		Author:        author.User,
		PlayUrl:       oss.ToRealURL(s.ctx, db_video.PlayURL),
		CoverUrl:      oss.ToRealURL(s.ctx, db_video.CoverURL),
		Title:         db_video.Title,
		FavoriteCount: favorite_count.FavoriteCount,
		IsFavorite:    is_favorite.IsFavorite,
		CommentCount:  comment_count.CommentCount,
	}

	return &publish.GetVideoByIdResponse{
		StatusCode: int32(codes.OK),
		StatusMsg:  nil,
		Video:      kitex_video,
	}, nil
}

func (s *PublishService) GetVideoByIdList(request *publish.GetVideoByIdListRequest) (resp *publish.GetVideoByIdListResponse, err error) {
	user_claims, err := Jwt.ExtractClaims(request.Token)
	var curr_user_id int64
	if err != nil {
		curr_user_id = 0
	} else {
		curr_user_id = user_claims.ID
	}

	db_videos, err := db.GetVideosByIDs(s.ctx, request.Id)
	if err != nil {
		err_msg := err.Error()
		return &publish.GetVideoByIdListResponse{StatusCode: int32(codes.Internal), StatusMsg: &err_msg}, err
	}
	
	err_chan := make(chan error)
	video_chan := make(chan *common.Video)
	kitex_videos := make([]*common.Video, 0, len(db_videos))

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
			comment_count, err := rpc.CommentRPC.GetVideoCommentCount(s.ctx, &comment.GetVideoCommentCountRequest{VideoId: db_video.ID})
			if err != nil {
				err_chan <- err
				return
			}

			video_chan <- &common.Video{
				Id:       db_video.ID,
				Author:   author.User,
				PlayUrl:  oss.ToRealURL(s.ctx, db_video.PlayURL),
				CoverUrl: oss.ToRealURL(s.ctx, db_video.CoverURL),
				Title:    db_video.Title,
				FavoriteCount: favorite_count.FavoriteCount,
				IsFavorite:    is_favorite.IsFavorite,
				CommentCount:  comment_count.CommentCount,
			}
		}(db_video)
	}

	for i := 0; i < len(db_videos); i++ {
		select {
		case err := <-err_chan:
			err_msg := err.Error()
			return &publish.GetVideoByIdListResponse{StatusCode: int32(codes.Internal), StatusMsg: &err_msg}, err
		case video := <-video_chan:
			kitex_videos = append(kitex_videos, video)
		}
	}

	return &publish.GetVideoByIdListResponse{
		StatusCode: int32(codes.OK),
		StatusMsg:  nil,
		VideoList:  kitex_videos,
	}, nil
}
