package favorite_test

import (
	"testing"

	favorite "github.com/HUST-MiniTiktok/mini_tiktok/kitex_gen/favorite"
)

// type GetVideoFavoriteInfoRequest struct {
// 	UserId  int64 `thrift:"user_id,1" frugal:"1,default,i64" json:"user_id"`
// 	VideoId int64 `thrift:"video_id,2" frugal:"2,default,i64" json:"video_id"`
// }

// type GetVideoFavoriteInfoResponse struct {
// 	StatusCode    int32   `thrift:"status_code,1" frugal:"1,default,i32" json:"status_code"`
// 	StatusMsg     *string `thrift:"status_msg,2,optional" frugal:"2,optional,string" json:"status_msg,omitempty"`
// 	FavoriteCount int64   `thrift:"favorite_count,3" frugal:"3,default,i64" json:"favorite_count"`
// 	IsFavorite    bool    `thrift:"is_favorite,4" frugal:"4,default,bool" json:"is_favorite"`
// }

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
