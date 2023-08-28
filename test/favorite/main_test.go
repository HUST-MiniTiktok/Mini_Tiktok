package favorite_test

import (
	"context"
	"os"
	"testing"

	favorite_dal "github.com/HUST-MiniTiktok/mini_tiktok/cmd/favorite/dal"
	favorite_service "github.com/HUST-MiniTiktok/mini_tiktok/cmd/favorite/service"
	user_dal "github.com/HUST-MiniTiktok/mini_tiktok/cmd/user/dal"
	user_service "github.com/HUST-MiniTiktok/mini_tiktok/cmd/user/service"
	user "github.com/HUST-MiniTiktok/mini_tiktok/kitex_gen/user"
)

var (
	ctx             = context.Background()
	FavoriteService *favorite_service.FavoriteService
	UserService     *user_service.UserService
	DemoUserName    = "demo@gmail.com"
	DemoPassword    = "demo!Password"
	token           string
	id              int64
	video_id        int64
)

func TestMain(m *testing.M) {
	os.Setenv("GO_ENV", "test")
	favorite_dal.Init()
	user_dal.Init()
	FavoriteService = favorite_service.NewFavoriteService(ctx)
	UserService = user_service.NewUserService(ctx)
	// DoRegister()
	DoLogin()

	code := m.Run()
	os.Exit(code)
}

func TestMainOrder(t *testing.T) {
	t.Run("favorite_action_1", TestFavoriteAction1)
	t.Run("favorite_list", TestFavoriteList)
	t.Run("get_video_favorite_info", TestGetVideoFavoriteInfo)
	t.Run("get_user_favorite_info", TestGetUserFavoriteInfo)
	t.Run("favorite_action_2", TestFavoriteAction2)
	t.Run("favorite_list", TestFavoriteList)
}

func DoLogin() {
	resp, err := UserService.Login(&user.UserLoginRequest{
		Username: DemoUserName,
		Password: DemoPassword,
	})
	if err != nil {
		panic(err)
	}
	token = resp.GetToken()
	id = resp.GetUserId()
}

func DoRegister() {
	resp, err := UserService.Register(&user.UserRegisterRequest{
		Username: DemoUserName,
		Password: DemoPassword,
	})
	if err != nil {
		panic(err)
	}
	token = resp.Token
	id = resp.UserId
}
