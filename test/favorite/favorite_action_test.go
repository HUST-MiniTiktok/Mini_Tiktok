package favorite_test

import (
	"testing"

	favorite "github.com/HUST-MiniTiktok/mini_tiktok/kitex_gen/favorite"
)

// Action type: 1: 点赞 2: 取消点赞
func TestFavoriteAction(t *testing.T) {
	resp, err := FavoriteService.FavoriteAction(&favorite.FavoriteActionRequest{
		Token:      DemoUser.Token,
		VideoId:    DemoVideo.Id,
		ActionType: 1,
	})
	if err != nil {
		t.Fatal(err)
	}
	if resp == nil {
		t.Fatal("resp is nil")
	}
	t.Logf("Favorite response: %v", resp)
}

func TestCancelFavoriteAction(t *testing.T) {
	resp, err := FavoriteService.FavoriteAction(&favorite.FavoriteActionRequest{
		Token:      DemoUser.Token,
		VideoId:    DemoVideo.Id,
		ActionType: 2,
	})
	if err != nil {
		t.Fatal(err)
	}
	if resp == nil {
		t.Fatal("resp is nil")
	}
	t.Logf("Favorite response: %v", resp)
}
