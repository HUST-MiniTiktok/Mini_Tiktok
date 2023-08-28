package favorite_test

import (
	"context"
	"testing"

	"bou.ke/monkey"
	"github.com/HUST-MiniTiktok/mini_tiktok/cmd/favorite/pack"
	favorite "github.com/HUST-MiniTiktok/mini_tiktok/kitex_gen/favorite"
)

// type GetUserFavoriteInfoRequest struct {
// 	UserId int64 `thrift:"user_id,1" frugal:"1,default,i64" json:"user_id"`
// }

// type GetUserFavoriteInfoResponse struct {
// 	StatusCode     int32   `thrift:"status_code,1" frugal:"1,default,i32" json:"status_code"`
// 	StatusMsg      *string `thrift:"status_msg,2,optional" frugal:"2,optional,string" json:"status_msg,omitempty"`
// 	TotalFavorited int64   `thrift:"total_favorited,3" frugal:"3,default,i64" json:"total_favorited"`
// 	FavoriteCount  int64   `thrift:"favorite_count,4" frugal:"4,default,i64" json:"favorite_count"`
// }

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
