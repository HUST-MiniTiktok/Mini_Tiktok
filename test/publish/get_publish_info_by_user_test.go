package publish_test

import (
	"testing"

	publish "github.com/HUST-MiniTiktok/mini_tiktok/kitex_gen/publish"
)

func TestGetPublishInfoByUser(t *testing.T) {
	resp, err := PublishService.GetPublishInfoByUserId(&publish.GetPublishInfoByUserIdRequest{UserId: DemoUser.Id})
	if err != nil {
		t.Fatal(err)
	}
	if resp == nil {
		t.Fatal("resp is nil")
	}
	if resp.GetVideoIds() == nil || len(resp.GetVideoIds()) == 0 {
		t.Fatal("video_ids is empty")
	}
	if resp.WorkCount == 0 {
		t.Fatal("work_count is 0")
	}
	t.Logf("get_info_by_user response: %v", resp)
}
