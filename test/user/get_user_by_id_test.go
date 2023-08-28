package user_test

import (
	"context"
	"testing"

	"bou.ke/monkey"
	"github.com/HUST-MiniTiktok/mini_tiktok/cmd/user/pack"
	user "github.com/HUST-MiniTiktok/mini_tiktok/kitex_gen/user"

	db "github.com/HUST-MiniTiktok/mini_tiktok/cmd/user/dal/db"
	common "github.com/HUST-MiniTiktok/mini_tiktok/kitex_gen/common"
)

func TestGetUser(t *testing.T) {
	monkey.Patch(pack.ToKitexUser, func(ctx context.Context, curr_user_token string, db_user *db.User) (*common.User, error) {
		return &common.User{
			Id:   db_user.ID,
			Name: db_user.UserName,
		}, nil
	})
	defer monkey.Unpatch(pack.ToKitexUser)

	resp, err := UserService.GetUserById(&user.UserRequest{
		UserId: id,
		Token:  token,
	})
	if err != nil {
		t.Fatal(err)
	}
	if resp == nil {
		t.Fatal("resp is nil")
	}
	t.Logf("get user response: %v", resp)
	if resp.User == nil {
		t.Fatal("user is nil")
	}
	if resp.User.GetId() != id || resp.User.GetName() != DemoUserName {
		t.Fatal("get user info not match")
	}
}
