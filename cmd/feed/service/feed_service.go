package service

import (
	"context"
	"time"

	db "github.com/HUST-MiniTiktok/mini_tiktok/cmd/feed/dal/db"
	"github.com/HUST-MiniTiktok/mini_tiktok/cmd/feed/pack"
	feed "github.com/HUST-MiniTiktok/mini_tiktok/kitex_gen/feed"
	"github.com/HUST-MiniTiktok/mini_tiktok/pkg/conf"
	"github.com/HUST-MiniTiktok/mini_tiktok/pkg/errno"
	"github.com/HUST-MiniTiktok/mini_tiktok/pkg/mw/jwt"
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

// GetFeed: get feed videos
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

	kitex_videos, err := pack.ToKitexVideoList(s.ctx, curr_user_id, request.GetToken(), db_videos)
	if err != nil {
		return pack.NewFeedResponse(err), err
	}

	resp = pack.NewFeedResponse(errno.Success)
	resp.VideoList = kitex_videos
	if len(db_videos) > 0 {
		next_time := db_videos[len(db_videos)-1].PublishTime.Unix()
		resp.NextTime = &next_time
	}
	return resp, nil
}
