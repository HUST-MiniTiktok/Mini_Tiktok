package feed_test

import (
	"context"
	"os"
	"testing"

	dal "github.com/HUST-MiniTiktok/mini_tiktok/cmd/feed/dal"
	service "github.com/HUST-MiniTiktok/mini_tiktok/cmd/feed/service"

	jwt "github.com/HUST-MiniTiktok/mini_tiktok/pkg/mw/jwt"
)

type DemoUserType struct {
	Id       int64
	UserName string
	Password string
	Token    string
}

var (
	ctx         = context.Background()
	Jwt         = jwt.NewJWT()
	FeedService *service.FeedService
	DemoUser    = DemoUserType{
		Id:       101,
		UserName: "demo@mail.com",
		Password: "demopassword",
	}
)

func TestMain(m *testing.M) {
	os.Setenv("GO_ENV", "test")
	dal.Init()
	FeedService = service.NewFeedService(ctx)

	code := m.Run()
	os.Exit(code)
}

func TestMainOrder(t *testing.T) {
	t.Run("feed_client", TestFeedClient)
	t.Run("feed", TestFeed)
}

func DoLogin() {
	token, err := Jwt.CreateToken(jwt.UserClaims{ID: DemoUser.Id})
	if err != nil {
		panic(err)
	}
	DemoUser.Token = token
}
