package favorite_test

import (
	"context"
	"testing"

	"bou.ke/monkey"
	"github.com/HUST-MiniTiktok/mini_tiktok/cmd/favorite/pack"
	favorite "github.com/HUST-MiniTiktok/mini_tiktok/kitex_gen/favorite"
)

func TestGetUserFavoriteInfo(t *testing.T) {

	monkey.Patch(pack.GetPublishInfoByUserId, func(ctx context.Context, UserId int64) (VideoIds []int64, err error) {
		vIds := make([]int64, 0, 1)
		vIds = append(vIds, DemoVideo.Id)
		return vIds, nil
	})
	defer monkey.Unpatch(pack.GetPublishInfoByUserId)

	resp, err := FavoriteService.GetUserFavoriteInfo(&favorite.GetUserFavoriteInfoRequest{UserId: DemoUser.Id})
	if err != nil {
		t.Fatal(err)
	}
	if resp == nil {
		t.Fatal("resp is nil")
	}
	t.Logf("get_user_favorite_info response: %v", resp)
}
