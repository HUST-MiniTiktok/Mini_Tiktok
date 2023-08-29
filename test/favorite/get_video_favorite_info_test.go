package favorite_test

import (
	"testing"

	favorite "github.com/HUST-MiniTiktok/mini_tiktok/kitex_gen/favorite"
)

func TestGetVideoFavoriteInfo(t *testing.T) {
	resp, err := FavoriteService.GetVideoFavoriteInfo(&favorite.GetVideoFavoriteInfoRequest{UserId: DemoUser.Id, VideoId: DemoVideo.Id})
	if err != nil {
		t.Fatal(err)
	}
	if resp == nil {
		t.Fatal("resp is nil")
	}
	t.Logf("get_video_favorite_info response: %v", resp)
}
