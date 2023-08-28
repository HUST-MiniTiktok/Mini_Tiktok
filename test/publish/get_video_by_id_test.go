package publish_test

import (
	"context"
	"testing"

	"bou.ke/monkey"
	"github.com/HUST-MiniTiktok/mini_tiktok/cmd/publish/pack"
	publish "github.com/HUST-MiniTiktok/mini_tiktok/kitex_gen/publish"

	db "github.com/HUST-MiniTiktok/mini_tiktok/cmd/publish/dal/db"
	common "github.com/HUST-MiniTiktok/mini_tiktok/kitex_gen/common"
	"github.com/HUST-MiniTiktok/mini_tiktok/pkg/mw/oss"
)

func TestGetVideoById(t *testing.T) {
	monkey.Patch(pack.ToKitexVideo, func(ctx context.Context, curr_user_id int64, curr_user_token string, db_video *db.Video) (*common.Video, error) {
		return &common.Video{
			Id:       db_video.ID,
			PlayUrl:  oss.ToRealURL(ctx, db_video.PlayURL),
			CoverUrl: oss.ToRealURL(ctx, db_video.CoverURL),
			Title:    db_video.Title,
		}, nil
	})
	defer monkey.Unpatch(pack.ToKitexVideo)

	resp, err := PublishService.GetVideoById(&publish.GetVideoByIdRequest{Token: token, Id: video_id})
	if err != nil {
		t.Fatal(err)
	}
	if resp == nil {
		t.Fatal("resp is nil")
	}
	if resp.GetVideo() == nil {
		t.Fatal("video is nil")
	}
	t.Logf("get_video response: %v", resp)
}
