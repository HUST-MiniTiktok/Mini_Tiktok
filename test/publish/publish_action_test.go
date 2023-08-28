package publish_test

import (
	"testing"

	publish "github.com/HUST-MiniTiktok/mini_tiktok/kitex_gen/publish"
)

func TestPublishAction(t *testing.T) {
	resp, err := PublishService.PublishAction(&publish.PublishActionRequest{
		Token: token,
		Data:  DemoVideoData,
		Title: DemoVideoTitle,
	})
	if err != nil {
		t.Fatal(err)
	}
	if resp == nil {
		t.Fatal("resp is nil")
	}
	t.Logf("publish response: %v", resp)
}
