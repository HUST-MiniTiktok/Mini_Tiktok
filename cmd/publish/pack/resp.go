package pack

import (
	publish "github.com/HUST-MiniTiktok/mini_tiktok/kitex_gen/publish"
	"github.com/HUST-MiniTiktok/mini_tiktok/pkg/utils/conv"
)

func NewPublishActionResponse(err error) *publish.PublishActionResponse {
	return conv.ToKitexBaseResponse(err, &publish.PublishActionResponse{}).(*publish.PublishActionResponse)
}

func NewPublishListResponse(err error) *publish.PublishListResponse {
	return conv.ToKitexBaseResponse(err, &publish.PublishListResponse{}).(*publish.PublishListResponse)
}

func NewGetVideoByIdResponse(err error) *publish.GetVideoByIdResponse {
	return conv.ToKitexBaseResponse(err, &publish.GetVideoByIdResponse{}).(*publish.GetVideoByIdResponse)
}

func NewGetVideoByIdListResponse(err error) *publish.GetVideoByIdListResponse {
	return conv.ToKitexBaseResponse(err, &publish.GetVideoByIdListResponse{}).(*publish.GetVideoByIdListResponse)
}
