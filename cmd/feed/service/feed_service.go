package service

import (
	"time"
	"context"

	db "github.com/HUST-MiniTiktok/mini_tiktok/cmd/feed/dal/db"
	feed "github.com/HUST-MiniTiktok/mini_tiktok/kitex_gen/feed"
	"github.com/HUST-MiniTiktok/mini_tiktok/conf"
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

type FeedService struct {
	ctx context.Context
}

func NewFeedService(ctx context.Context) *FeedService {
	return &FeedService{ctx: ctx}
}

func (s *FeedService) GetFeed(request *feed.FeedRequest) (resp *feed.FeedResponse, err error) {
	var db_videos []*db.Video

	var last_time time.Time
	if request.LatestTimestamp != nil {
		last_time = time.Unix((*request.LatestTimestamp)/1000, 0)
	} else {
		last_time = time.Now()
	}

	db_videos, err = db.GetVideosByLastPublishTime(s.ctx, last_time)
	if err != nil {
		err_msg := err.Error()
		resp = &feed.FeedResponse{StatusCode: int32(codes.Internal), StatusMsg: &err_msg}
		return
	}

	// TODO: 从数据库中获取db.Video，转换为feed.Video
	_ = db_videos
	return
}