package message_test

import (
	"testing"

	rpc "github.com/HUST-MiniTiktok/mini_tiktok/pkg/rpc"
)

func TestMessageClient(t *testing.T) {
	client := rpc.NewMessageClient()
	if client == nil {
		t.Errorf("new message client failed")
		return
	}
	t.Logf("new message client success")
}
