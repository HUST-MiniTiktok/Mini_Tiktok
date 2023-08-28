package favorite_test

import (
	"context"
	"testing"

	"bou.ke/monkey"
	"github.com/HUST-MiniTiktok/mini_tiktok/cmd/favorite/pack"
	"github.com/HUST-MiniTiktok/mini_tiktok/kitex_gen/common"
	favorite "github.com/HUST-MiniTiktok/mini_tiktok/kitex_gen/favorite"
)

func TestFavoriteList(t *testing.T) {
	monkey.Patch(pack.GetVideoByIdList, func(ctx context.Context, videoIDList []int64, token string) (VideoList []*common.Video, err error) {
		vlist := make([]*common.Video, 0, 1)
		v := &common.Video{
			Id: DemoVideo.Id,
			// PlayUrl:  DemoVideo.PlayUrl,
			// CoverUrl: DemoVideo.CoverUrl,
			// Title:    DemoVideo.Title,
		}
		vlist = append(vlist, v)
		return vlist, nil
	})
	defer monkey.Unpatch(pack.GetVideoByIdList)

	resp, err := FavoriteService.FavoriteList(&favorite.FavoriteListRequest{UserId: DemoUser.Id, Token: DemoUser.Token})
	if err != nil {
		t.Fatal(err)
	}
	if resp == nil {
		t.Fatal("resp is nil")
	}
	if resp.GetVideoList() == nil || len(resp.GetVideoList()) == 0 {
		t.Fatal("video_list is empty")
	}
	t.Logf("Favorite_list response: %v", resp)
}
