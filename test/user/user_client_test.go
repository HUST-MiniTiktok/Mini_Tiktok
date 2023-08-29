package user_test

import (
	"testing"

	rpc "github.com/HUST-MiniTiktok/mini_tiktok/pkg/rpc"
)

func TestUserClient(t *testing.T) {
	client := rpc.NewUserClient()
	if client == nil {
		t.Errorf("new user client failed")
		return
	}
	t.Logf("new user client success")
}
