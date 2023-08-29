package comment_test

import (
	"context"
	"testing"

	"bou.ke/monkey"
	"github.com/HUST-MiniTiktok/mini_tiktok/cmd/comment/dal/db"
	"github.com/HUST-MiniTiktok/mini_tiktok/cmd/comment/pack"
	comment "github.com/HUST-MiniTiktok/mini_tiktok/kitex_gen/comment"
	common "github.com/HUST-MiniTiktok/mini_tiktok/kitex_gen/common"
)

func TestCommentAction(t *testing.T) {
	monkey.Patch(pack.ToKitexComment, func(ctx context.Context, db_comment *db.Comment, curr_user_token string) (*comment.Comment, error) {
		testUser := common.User{
			Id: DemoUser.Id,
		}
		return &comment.Comment{
			Id:         db_comment.ID,
			User:       &testUser,
			Content:    db_comment.CommentText,
			CreateDate: db_comment.CreatedAt.Format("01-02"),
		}, nil
	})
	defer monkey.Unpatch(pack.ToKitexComment)

	resp, err := CommentService.CommentAction(&comment.CommentActionRequest{
		Token:       DemoUser.Token,
		VideoId:     DemoVideo.Id,
		ActionType:  1,
		CommentText: &DemoComment.CommentText,
		CommentId:   nil,
	})
	if err != nil {
		t.Fatal(err)
	}
	if resp == nil {
		t.Fatal("resp is nil")
	}
	DemoComment.Id = resp.Comment.Id
	t.Logf("Comment response: %v", resp)
}

func TestDeleteCommentAction(t *testing.T) {
	resp, err := CommentService.CommentAction(&comment.CommentActionRequest{
		Token:       DemoUser.Token,
		VideoId:     DemoVideo.Id,
		ActionType:  2,
		CommentText: nil,
		CommentId:   &DemoComment.Id,
	})
	if err != nil {
		t.Fatal(err)
	}
	if resp == nil {
		t.Fatal("resp is nil")
	}
	t.Logf("Delete Comment response: %v", resp)
}
