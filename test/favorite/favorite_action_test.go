package favorite_test

import (
	"testing"

	favorite "github.com/HUST-MiniTiktok/mini_tiktok/kitex_gen/favorite"
)

// type FavoriteActionRequest struct {
// 	Token      string `thrift:"token,1" frugal:"1,default,string" json:"token"`
// 	VideoId    int64  `thrift:"video_id,2" frugal:"2,default,i64" json:"video_id"`
// 	ActionType int32  `thrift:"action_type,3" frugal:"3,default,i32" json:"action_type"`
// }

// Action type: 1: 点赞 2: 取消点赞
func TestFavoriteAction1(t *testing.T) {
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

func TestFavoriteAction2(t *testing.T) {
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
