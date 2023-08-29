package user_test

import (
	"context"
	"os"
	"testing"

	dal "github.com/HUST-MiniTiktok/mini_tiktok/cmd/user/dal"
	service "github.com/HUST-MiniTiktok/mini_tiktok/cmd/user/service"
)

type DemoUserType struct {
	Id       int64
	UserName string
	Password string
	Token    string
}

var (
	ctx         = context.Background()
	UserService *service.UserService
	DemoUser    = DemoUserType{
		UserName: "demo@mail.com",
		Password: "demopassword",
	}
)

func TestMain(m *testing.M) {
	os.Setenv("GO_ENV", "test")
	dal.Init()
	UserService = service.NewUserService(ctx)

	code := m.Run()
	os.Exit(code)
}

func TestMainOrder(t *testing.T) {
	t.Run("user_client", TestUserClient)
	t.Run("register", TestRegister)
	t.Run("login", TestLogin)
	t.Run("check", TestCheck)
	t.Run("get_user_by_id", TestGetUser)
}
