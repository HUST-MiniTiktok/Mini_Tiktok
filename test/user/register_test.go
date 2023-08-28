package user_test

import (
	"testing"

	user "github.com/HUST-MiniTiktok/mini_tiktok/kitex_gen/user"
)

func TestRegister(t *testing.T) {
	resp, err := UserService.Register(&user.UserRegisterRequest{
		Username: DemoUser.UserName,
		Password: DemoUser.Password,
	})
	if err != nil {
		t.Fatal(err)
	}
	if resp == nil {
		t.Fatal("resp is nil")
	}
	t.Logf("register response: %v", resp)
	DemoUser.Token = resp.Token
	DemoUser.Id = resp.UserId
}
