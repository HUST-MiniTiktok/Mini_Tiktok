package pack

import (
	"context"

	"github.com/HUST-MiniTiktok/mini_tiktok/cmd/favorite/client"
	"github.com/HUST-MiniTiktok/mini_tiktok/kitex_gen/common"
	publish "github.com/HUST-MiniTiktok/mini_tiktok/kitex_gen/publish"
)

func GetVideoByIdList(ctx context.Context, videoIDList []int64, token string) (VideoList []*common.Video, err error) {
	videosResponse, err := client.PublishRPC.GetVideoByIdList(ctx, &publish.GetVideoByIdListRequest{Id: videoIDList, Token: token})
	if err != nil {
		return nil, err
	}
	return videosResponse.VideoList, nil
}

func GetPublishInfoByUserId(ctx context.Context, UserId int64) (VideoIds []int64, err error) {
	publishInfoResp, err := client.PublishRPC.GetPublishInfoByUserId(ctx, &publish.GetPublishInfoByUserIdRequest{UserId: UserId})
	if err != nil {
		return nil, err
	}
	return publishInfoResp.VideoIds, err
}
