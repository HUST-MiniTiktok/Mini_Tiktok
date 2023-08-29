package comment_test

import (
	"context"
	"os"
	"testing"

	comment_dal "github.com/HUST-MiniTiktok/mini_tiktok/cmd/comment/dal"
	comment_service "github.com/HUST-MiniTiktok/mini_tiktok/cmd/comment/service"
	jwt "github.com/HUST-MiniTiktok/mini_tiktok/pkg/mw/jwt"
)

type DemoUserType struct {
	Id       int64
	UserName string
	Password string
	Token    string
}

type DemoVideoType struct {
	Id int64
}

type DemoCommentType struct {
	Id          int64
	UserId      int64
	VideoId     int64
	CommentText string
}

var (
	ctx            = context.Background()
	CommentService *comment_service.CommentService
	Jwt            = jwt.NewJWT()

	DemoUser = DemoUserType{
		Id:       101,
		UserName: "demo@mail.com",
		Password: "demopassword",
	}
	DemoVideo = DemoVideoType{
		Id: 11,
	}
	DemoComment = DemoCommentType{
		Id:          0,
		UserId:      101,
		VideoId:     11,
		CommentText: "test comment",
	}
)

func TestMain(m *testing.M) {
	os.Setenv("GO_ENV", "test")
	comment_dal.Init()

	CommentService = comment_service.NewCommentService(ctx)

	DoLogin()

	code := m.Run()
	os.Exit(code)
}

func TestMainOrder(t *testing.T) {
	t.Run("test_omment_action", TestCommentAction)
	t.Run("test_get_video_comment_count", TestGetVideoCommentCount)
	t.Run("test_comment_list", TestCommentList)
	t.Run("test_comment_action_delete", TestDeleteCommentAction)
	t.Run("test_get_video_comment_count_after_delete", TestGetVideoCommentCount)
}

func DoLogin() {
	token, err := Jwt.CreateToken(jwt.UserClaims{ID: DemoUser.Id})
	if err != nil {
		panic(err)
	}
	DemoUser.Token = token
}
