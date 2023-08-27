package user_test

import (
	"context"
	"os"
	"testing"

	dal "github.com/HUST-MiniTiktok/mini_tiktok/cmd/user/dal"
	service "github.com/HUST-MiniTiktok/mini_tiktok/cmd/user/service"
)

var (
	ctx          = context.Background()
	UserService  *service.UserService
	DemoUserName = "demo@gmail.com"
	DemoPassword = "demo!Password"
	token        string
	id           int64
)

func TestMain(m *testing.M) {
	os.Setenv("GO_ENV", "test")
	dal.Init()
	UserService = service.NewUserService(ctx)

	code := m.Run()
	os.Exit(code)
}

func TestMainOrder(t *testing.T) {
	t.Run("register", TestRegister)
	t.Run("login", TestLogin)
	t.Run("check", TestCheckUser)
	t.Run("get", TestGetUser)
}
