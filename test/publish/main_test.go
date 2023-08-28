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

type DemoUserType struct {
	Id       int64
	UserName string
	Password string
	Token    string
}

type DemoVideoType struct {
	Id    int64
	Title string
	Data  []byte
	Path  string
}

var (
	ctx            = context.Background()
	PublishService *publish_service.PublishService
	UserService    *user_service.UserService
	DemoUser       = DemoUserType{
		UserName: "demo@mail.com",
		Password: "demopassword",
	}
	DemoVideo = DemoVideoType{
		Title: "bear",
		Path:  "bear.mp4",
	}
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
		Username: DemoUser.UserName,
		Password: DemoUser.Password,
	})
	if err != nil {
		panic(err)
	}
	DemoUser.Token = resp.GetToken()
	DemoUser.Id = resp.GetUserId()
}

func DoLoadVideo() {
	var err error
	data, err := os.ReadFile(DemoVideo.Path)
	if err != nil {
		panic(err)
	}
	DemoVideo.Data = data
}
