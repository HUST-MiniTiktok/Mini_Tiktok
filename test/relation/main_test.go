package relation_test

import (
	"context"
	"os"
	"testing"

	dal "github.com/HUST-MiniTiktok/mini_tiktok/cmd/relation/dal"
	service "github.com/HUST-MiniTiktok/mini_tiktok/cmd/relation/service"

	jwt "github.com/HUST-MiniTiktok/mini_tiktok/pkg/mw/jwt"
)

type DemoUserType struct {
	Id       int64
	UserName string
	Password string
	Token    string
}

var (
	ctx             = context.Background()
	Jwt             = jwt.NewJWT()
	RelationService *service.RelationService
	DemoUserList    = []DemoUserType{
		{Id: 11, UserName: "demo1@mail.com", Password: "demopass1"},
		{Id: 12, UserName: "demo2@mail.com", Password: "demopass2"},
		{Id: 13, UserName: "demo3@mail.com", Password: "demopass3"},
	}
)

func TestMain(m *testing.M) {
	os.Setenv("GO_ENV", "test")
	dal.Init()
	RelationService = service.NewRelationService(ctx)

	DoLogin()

	code := m.Run()
	os.Exit(code)
}

func TestMainOrder(t *testing.T) {
	t.Run("relation_client", TestRelationClient)
	t.Run("relation_follow_action", TestRelationFollowAction)
	t.Run("relation_follow_list_after_follow", TestRelationFollowList)
	t.Run("relation_follower_list_after_follow", TestRelationFollowerList)
	t.Run("relation_friend_list_after_follow", TestRelationFriendList)
	t.Run("get_follow_info_by_user_after_follow", TestGetFollowInfoByUser)
	t.Run("relation_unfollow_action", TestRelationUnFollowAction)
	t.Run("relation_follow_list_after_unfollow", TestRelationFollowList)
	t.Run("relation_follower_list_after_unfollow", TestRelationFollowerList)
	t.Run("relation_friend_list_after_unfollow", TestRelationFriendList)
	t.Run("get_follow_info_by_user_after_unfollow", TestGetFollowInfoByUser)
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
