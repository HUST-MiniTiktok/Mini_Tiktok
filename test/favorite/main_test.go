package favorite_test

import (
	"context"
	"os"
	"testing"

	favorite_dal "github.com/HUST-MiniTiktok/mini_tiktok/cmd/favorite/dal"
	favorite_service "github.com/HUST-MiniTiktok/mini_tiktok/cmd/favorite/service"
	jwt "github.com/HUST-MiniTiktok/mini_tiktok/pkg/mw/jwt"
)

type DemoUserType struct {
	Id       int64
	UserName string
	Password string
	Token    string
}

type DemoVideoType struct {
	Id       int64
	Title    string
	Author   string
	PlayUrl  string
	CoverUrl string
	Data     []byte
	Path     string
}

var (
	ctx             = context.Background()
	FavoriteService *favorite_service.FavoriteService
	Jwt             = jwt.NewJWT()

	DemoUser = DemoUserType{
		Id:       101,
		UserName: "demo@mail.com",
		Password: "demopassword",
	}
	DemoVideo = DemoVideoType{
		Id:       11,
		PlayUrl:  "play",
		CoverUrl: "cober",
		Title:    "bear",
		Path:     "bear.mp4",
	}
)

func TestMain(m *testing.M) {
	os.Setenv("GO_ENV", "test")
	favorite_dal.Init()

	FavoriteService = favorite_service.NewFavoriteService(ctx)

	// DoRegister()
	DoLogin()

	code := m.Run()
	os.Exit(code)
}

func TestMainOrder(t *testing.T) {
	t.Run("favorite_action_1", TestFavoriteAction1)
	t.Run("favorite_list_1", TestFavoriteList)
	t.Run("get_video_favorite_info", TestGetVideoFavoriteInfo)
	t.Run("get_user_favorite_info", TestGetUserFavoriteInfo)
	t.Run("favorite_action_2", TestFavoriteAction2)
	t.Run("favorite_list_2", TestFavoriteList)
}

func DoLogin() {
	token, err := Jwt.CreateToken(jwt.UserClaims{ID: DemoUser.Id})
	if err != nil {
		panic(err)
	}
	DemoUser.Token = token
}
