package conv

import (
	kitex_user "github.com/HUST-MiniTiktok/mini_tiktok/cmd/user/kitex_gen/user"
	hertz_publish "github.com/HUST-MiniTiktok/mini_tiktok/cmd/api/biz/model/publish"
	kitex_publish "github.com/HUST-MiniTiktok/mini_tiktok/cmd/publish/kitex_gen/publish"
)

func ToHertzVideo(video *kitex_publish.Video) *hertz_publish.Video {
	return &hertz_publish.Video{
		ID: video.Id,
		Author: (*hertz_publish.User)(ToHertzUser((*kitex_user.User) (video.Author))),
		PlayURL: video.PlayUrl,
		CoverURL: video.CoverUrl,
		Title: video.Title,
	}
}

func ToHertzVideoList(videoList []*kitex_publish.Video) []*hertz_publish.Video {
	var hertzVideoList []*hertz_publish.Video
	for _, video := range videoList {
		hertzVideoList = append(hertzVideoList, ToHertzVideo(video))
	}
	return hertzVideoList
}


func ToKitexPublishActionRequest(req *hertz_publish.PublishActionRequest) *kitex_publish.PublishActionRequest {
	return &kitex_publish.PublishActionRequest{
		Token:  req.Token,
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
		VideoList: ToHertzVideoList(resp.VideoList),
	}
}