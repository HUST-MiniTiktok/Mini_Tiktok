package service

import (
	"context"
	"time"

	"github.com/HUST-MiniTiktok/mini_tiktok/cmd/feed/client"
	db "github.com/HUST-MiniTiktok/mini_tiktok/cmd/feed/dal/db"
	"github.com/HUST-MiniTiktok/mini_tiktok/cmd/feed/pack"
	comment "github.com/HUST-MiniTiktok/mini_tiktok/kitex_gen/comment"
	common "github.com/HUST-MiniTiktok/mini_tiktok/kitex_gen/common"
	favorite "github.com/HUST-MiniTiktok/mini_tiktok/kitex_gen/favorite"
	feed "github.com/HUST-MiniTiktok/mini_tiktok/kitex_gen/feed"
	user "github.com/HUST-MiniTiktok/mini_tiktok/kitex_gen/user"
	"github.com/HUST-MiniTiktok/mini_tiktok/pkg/conf"
	"github.com/HUST-MiniTiktok/mini_tiktok/pkg/errno"
	"github.com/HUST-MiniTiktok/mini_tiktok/pkg/mw/jwt"
	"github.com/HUST-MiniTiktok/mini_tiktok/pkg/mw/oss"
	"github.com/HUST-MiniTiktok/mini_tiktok/pkg/utils"
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

type FeedService struct {
	ctx context.Context
}

func NewFeedService(ctx context.Context) *FeedService {
	return &FeedService{ctx: ctx}
}

func (s *FeedService) GetFeed(request *feed.FeedRequest) (resp *feed.FeedResponse, err error) {
	claim, err := Jwt.ExtractClaims(request.GetToken())
	var curr_user_id int64
	if err != nil {
		curr_user_id = 0
	} else {
		curr_user_id = claim.ID
	}

	var last_time time.Time
	if request.LatestTimestamp != nil {
		last_time = utils.MillTimeStampToTime(*request.LatestTimestamp)
	} else {
		last_time = time.Now()
	}

	var db_videos []*db.Video
	db_videos, err = db.GetVideosByLastPublishTime(s.ctx, last_time)
	if err != nil {
		return pack.NewFeedResponse(err), err
	}

	err_chan := make(chan error)
	video_chan := make(chan *common.Video)
	kitex_videos := make([]*common.Video, 0, len(db_videos))

	for _, db_video := range db_videos {
		go func(db_video *db.Video) {
			author, err := client.UserRPC.User(s.ctx, &user.UserRequest{UserId: db_video.AuthorID, Token: request.GetToken()})
			if err != nil {
				err_chan <- err
				return
			}
			favorite_info, err := client.FavoriteRPC.GetVideoFavoriteInfo(s.ctx, &favorite.GetVideoFavoriteInfoRequest{UserId: curr_user_id, VideoId: db_video.ID})
			if err != nil {
				err_chan <- err
				return
			}
			comment_count, err := client.CommentRPC.GetVideoCommentCount(s.ctx, &comment.GetVideoCommentCountRequest{VideoId: db_video.ID})
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
				FavoriteCount: favorite_info.FavoriteCount,
				IsFavorite:    favorite_info.IsFavorite,
				CommentCount:  comment_count.CommentCount,
			}
		}(db_video)
	}

	for i := 0; i < len(db_videos); i++ {
		select {
		case err := <-err_chan:
			return pack.NewFeedResponse(err), err
		case video := <-video_chan:
			kitex_videos = append(kitex_videos, video)
		}
	}

	resp = pack.NewFeedResponse(errno.Success)
	resp.VideoList = kitex_videos
	if len(db_videos) > 0 {
		next_time := db_videos[len(db_videos)-1].PublishTime.Unix()
		resp.NextTime = &next_time
	}
	return resp, nil
}
