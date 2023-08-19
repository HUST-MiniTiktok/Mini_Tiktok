package pack

import (
	feed "github.com/HUST-MiniTiktok/mini_tiktok/kitex_gen/feed"
	"github.com/HUST-MiniTiktok/mini_tiktok/pkg/utils/conv"
)

func NewFeedResponse(err error) *feed.FeedResponse {
	return conv.ToKitexBaseResponse(err, &feed.FeedResponse{}).(*feed.FeedResponse)
}
