package publish_test

import (
	"testing"

	rpc "github.com/HUST-MiniTiktok/mini_tiktok/pkg/rpc"
)

func TestPublishClient(t *testing.T) {
	client := rpc.NewPublishClient()
	if client == nil {
		t.Errorf("new publish client failed")
		return
	}
	t.Logf("new publish client success")
}
