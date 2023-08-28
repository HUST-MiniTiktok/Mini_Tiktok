package user_test

import (
	"testing"

	user "github.com/HUST-MiniTiktok/mini_tiktok/kitex_gen/user"
)

func TestCheckUser(t *testing.T) {
	resp, err := UserService.CheckUserIsExist(&user.CheckUserIsExistRequest{
		UserId: DemoUser.Id,
	})
	if err != nil {
		t.Fatal(err)
	}
	if resp == nil {
		t.Fatal("resp is nil")
	}
	t.Logf("check user response: %v", resp)
	if resp.IsExist != true {
		t.Fatal("check user failed")
	}
}
