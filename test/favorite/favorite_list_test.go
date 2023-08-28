package favorite_test

import (
	"context"
	"testing"

	"bou.ke/monkey"
	"github.com/HUST-MiniTiktok/mini_tiktok/cmd/feed/dal/db"
	"github.com/HUST-MiniTiktok/mini_tiktok/cmd/feed/pack"
	"github.com/HUST-MiniTiktok/mini_tiktok/kitex_gen/common"
	favorite "github.com/HUST-MiniTiktok/mini_tiktok/kitex_gen/favorite"
	"github.com/HUST-MiniTiktok/mini_tiktok/pkg/mw/oss"
)

// type FavoriteListRequest struct {
// 	UserId int64  `thrift:"user_id,1" frugal:"1,default,i64" json:"user_id"`
// 	Token  string `thrift:"token,2" frugal:"2,default,string" json:"token"`
// }
// type FavoriteListResponse struct {
// 	StatusCode int32           `thrift:"status_code,1" frugal:"1,default,i32" json:"status_code"`
// 	StatusMsg  *string         `thrift:"status_msg,2,optional" frugal:"2,optional,string" json:"status_msg,omitempty"`
// 	VideoList  []*common.Video `thrift:"video_list,3" frugal:"3,default,list<common.Video>" json:"video_list"`
// }

func TestFavoriteList(t *testing.T) {
	monkey.Patch(pack.ToKitexVideo, func(ctx context.Context, curr_user_id int64, curr_user_token string, db_video *db.Video) (*common.Video, error) {
		return &common.Video{
			Id:       db_video.ID,
			PlayUrl:  oss.ToRealURL(ctx, db_video.PlayURL),
			CoverUrl: oss.ToRealURL(ctx, db_video.CoverURL),
			Title:    db_video.Title,
		}, nil
	})
	defer monkey.Unpatch(pack.ToKitexVideo)

	resp, err := FavoriteService.FavoriteList(&favorite.FavoriteListRequest{UserId: id, Token: token})
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
	// video_id = resp.GetVideoList()[0].GetId()
}
