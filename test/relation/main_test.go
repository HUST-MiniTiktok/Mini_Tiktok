package relation_test

import (
	"context"
	"os"
	"testing"

	relation_dal "github.com/HUST-MiniTiktok/mini_tiktok/cmd/relation/dal"
	relation_service "github.com/HUST-MiniTiktok/mini_tiktok/cmd/relation/service"

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

var (
	ctx             = context.Background()
	RelationService *relation_service.RelationService
	UserService     *user_service.UserService
	DemoUserList    = []DemoUserType{
		{UserName: "demo1@mail.com", Password: "demopass1"},
		{UserName: "demo2@mail.com", Password: "demopass2"},
		{UserName: "demo3@mail.com", Password: "demopass3"},
	}
)

func TestMain(m *testing.M) {
	os.Setenv("GO_ENV", "test")
	relation_dal.Init()
	user_dal.Init()
	RelationService = relation_service.NewRelationService(ctx)
	UserService = user_service.NewUserService(ctx)

	DoRegister()

	code := m.Run()
	os.Exit(code)
}

func DoRegister() {
	for _, u := range DemoUserList {
		resp, err := UserService.Register(&user.UserRegisterRequest{
			Username: u.UserName,
			Password: u.Password,
		})
		if err != nil {
			panic(err)
		}
		if resp == nil {
			panic("resp is nil")
		}
		u.Token = resp.GetToken()
		u.Id = resp.GetUserId()
	}
}
