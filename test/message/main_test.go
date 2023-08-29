package message_test

import (
	"context"
	"os"
	"testing"

	dal "github.com/HUST-MiniTiktok/mini_tiktok/cmd/message/dal"
	service "github.com/HUST-MiniTiktok/mini_tiktok/cmd/message/service"

	jwt "github.com/HUST-MiniTiktok/mini_tiktok/pkg/mw/jwt"
)

type DemoUserType struct {
	Id       int64
	UserName string
	Password string
	Token    string
}

var (
	ctx            = context.Background()
	Jwt            = jwt.NewJWT()
	MessageService *service.MessageService

	DemoUserList = []DemoUserType{
		{Id: 11, UserName: "demo1@mail.com", Password: "demopass1"},
		{Id: 12, UserName: "demo2@mail.com", Password: "demopass2"},
		{Id: 13, UserName: "demo3@mail.com", Password: "demopass3"},
	}
)

func TestMain(m *testing.M) {
	os.Setenv("GO_ENV", "test")
	dal.Init()
	MessageService = service.NewMessageService(ctx)

	DoLogin()

	code := m.Run()
	os.Exit(code)
}

func TestMainOrder(t *testing.T) {
	t.Run("message_client", TestMessageClient)
	t.Run("message_action", TestMessageAction)
	t.Run("message_chat", TestMessageChat)
	t.Run("get_friend_latest_msg", TestGetFriendLatestMsg)
}

func DoLogin() {
	for i, user := range DemoUserList {
		token, err := Jwt.CreateToken(jwt.UserClaims{ID: user.Id})
		if err != nil {
			panic(err)
		}
		DemoUserList[i].Token = token
	}
}
