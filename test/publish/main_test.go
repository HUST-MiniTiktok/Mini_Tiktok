package publish_test

import (
	"context"
	"os"
	"testing"

	publish_dal "github.com/HUST-MiniTiktok/mini_tiktok/cmd/publish/dal"
	publish_service "github.com/HUST-MiniTiktok/mini_tiktok/cmd/publish/service"
	user_dal "github.com/HUST-MiniTiktok/mini_tiktok/cmd/user/dal"
	user_service "github.com/HUST-MiniTiktok/mini_tiktok/cmd/user/service"
	user "github.com/HUST-MiniTiktok/mini_tiktok/kitex_gen/user"
)

var (
	ctx            = context.Background()
	PublishService *publish_service.PublishService
	UserService    *user_service.UserService
	DemoUserName   = "demo@gmail.com"
	DemoPassword   = "demo!Password"
	DemoVideoData  []byte
	DemoVideoFile  = "bear.mp4"
	DemoVideoTitle = "bear"
	token          string
	id             int64
	video_id       int64
)

func TestMain(m *testing.M) {
	os.Setenv("GO_ENV", "test")
	publish_dal.Init()
	user_dal.Init()
	PublishService = publish_service.NewPublishService(ctx)
	UserService = user_service.NewUserService(ctx)
	DoLogin()
	DoLoadVideo()

	code := m.Run()
	os.Exit(code)
}

func TestMainOrder(t *testing.T) {
	t.Run("publish_action", TestPublishAction)
	t.Run("publish_list", TestPublishList)
	t.Run("get_video_by_id", TestGetVideoById)
	t.Run("get_info_by_user", TestGetInfoByUser)
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

func DoLoadVideo() {
	var err error
	DemoVideoData, err = os.ReadFile(DemoVideoFile)
	if err != nil {
		panic(err)
	}
}
