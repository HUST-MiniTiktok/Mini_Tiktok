package user_test

import (
	"testing"

	user "github.com/HUST-MiniTiktok/mini_tiktok/kitex_gen/user"
)

func TestRegister(t *testing.T) {
	resp, err := UserService.Register(ctx, &user.UserRegisterRequest{
		Username: DemoUserName,
		Password: DemoPassword,
	})
	if err != nil {
		t.Fatal(err)
	}
	if resp == nil {
		t.Fatal("resp is nil")
	}
	t.Logf("register response: %v", resp)
	token = resp.Token
	id = resp.UserId
}
