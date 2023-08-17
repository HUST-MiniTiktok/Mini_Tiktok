package conv

import (
	hertz_feed "github.com/HUST-MiniTiktok/mini_tiktok/cmd/api/biz/model/feed"
	kitex_feed "github.com/HUST-MiniTiktok/mini_tiktok/kitex_gen/feed"
)

func ToKitexFeedRequest(req *hertz_feed.FeedRequest) *kitex_feed.FeedRequest {
	return &kitex_feed.FeedRequest{
		LatestTimestamp: req.LatestTimestamp,
		Token: req.Token,
	}
}

func ToHertzFeedResponse(resp *kitex_feed.FeedResponse) *hertz_feed.FeedResponse {
	return &hertz_feed.FeedResponse{
		StatusCode: resp.StatusCode,
		StatusMsg:  resp.StatusMsg,
		VideoList: ToHertzVideoList(resp.VideoList),
		NextTime: resp.NextTime,
	}
}