package conv

import (
	hertz_common "github.com/HUST-MiniTiktok/mini_tiktok/cmd/api/biz/model/common"
	hertz_publish "github.com/HUST-MiniTiktok/mini_tiktok/cmd/api/biz/model/publish"
	kitex_common "github.com/HUST-MiniTiktok/mini_tiktok/kitex_gen/common"
	kitex_publish "github.com/HUST-MiniTiktok/mini_tiktok/kitex_gen/publish"
)

func ToHertzVideo(video *kitex_common.Video) *hertz_common.Video {
	return &hertz_common.Video{
		ID:       video.Id,
		Author:   ToHertzUser(video.Author),
		PlayURL:  video.PlayUrl,
		CoverURL: video.CoverUrl,
		Title:    video.Title,
	}
}

func ToHertzVideoList(videoList []*kitex_common.Video) []*hertz_common.Video {
	hertzVideoList := make([]*hertz_common.Video, 0, len(videoList))
	for _, video := range videoList {
		hertzVideoList = append(hertzVideoList, ToHertzVideo(video))
	}
	return hertzVideoList
}

func ToKitexPublishActionRequest(req *hertz_publish.PublishActionRequest) *kitex_publish.PublishActionRequest {
	return &kitex_publish.PublishActionRequest{
		Token: req.Token,
		Data:  req.Data,
		Title: req.Title,
	}
}

func ToHertzPublishActionResponse(resp *kitex_publish.PublishActionResponse) *hertz_publish.PublishActionResponse {
	return &hertz_publish.PublishActionResponse{
		StatusCode: resp.StatusCode,
		StatusMsg:  resp.StatusMsg,
	}
}

func ToKitexPublishListRequest(req *hertz_publish.PublishListRequest) *kitex_publish.PublishListRequest {
	return &kitex_publish.PublishListRequest{
		UserId: req.UserID,
		Token:  req.Token,
	}
}

func ToHertzPublishListResponse(resp *kitex_publish.PublishListResponse) *hertz_publish.PublishListResponse {
	return &hertz_publish.PublishListResponse{
		StatusCode: resp.StatusCode,
		StatusMsg:  resp.StatusMsg,
		VideoList:  ToHertzVideoList(resp.VideoList),
	}
}
