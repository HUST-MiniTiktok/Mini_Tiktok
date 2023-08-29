package publish_test

import (
	"context"
	"os"
	"testing"

	dal "github.com/HUST-MiniTiktok/mini_tiktok/cmd/publish/dal"
	service "github.com/HUST-MiniTiktok/mini_tiktok/cmd/publish/service"

	jwt "github.com/HUST-MiniTiktok/mini_tiktok/pkg/mw/jwt"
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
	Jwt            = jwt.NewJWT()
	PublishService *service.PublishService

	DemoUser = DemoUserType{
		Id:       101,
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
	dal.Init()
	PublishService = service.NewPublishService(ctx)

	DoLogin()
	DoLoadVideo()

	code := m.Run()
	os.Exit(code)
}

func TestMainOrder(t *testing.T) {
	t.Run("publish_client", TestPublishClient)
	t.Run("publish_action", TestPublishAction)
	t.Run("publish_list", TestPublishList)
	t.Run("get_video_by_id", TestGetVideoById)
	t.Run("get_publish_info_by_user", TestGetPublishInfoByUser)
}

func DoLogin() {
	token, err := Jwt.CreateToken(jwt.UserClaims{ID: DemoUser.Id})
	if err != nil {
		panic(err)
	}
	DemoUser.Token = token
}

func DoLoadVideo() {
	var err error
	data, err := os.ReadFile(DemoVideo.Path)
	if err != nil {
		panic(err)
	}
	DemoVideo.Data = data
}
