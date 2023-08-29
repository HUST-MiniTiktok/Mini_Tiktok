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

func TestCommentList(t *testing.T) {
	monkey.Patch(pack.ToKitexCommentList, func(ctx context.Context, db_comments []*db.Comment, curr_user_token string) ([]*comment.Comment, error) {
		clist := make([]*comment.Comment, 0, 1)
		testUser := common.User{
			Id: DemoUser.Id,
		}
		c := &comment.Comment{
			Id:         db_comments[0].ID,
			User:       &testUser,
			Content:    db_comments[0].CommentText,
			CreateDate: db_comments[0].CreatedAt.Format("01-02"),
		}
		clist = append(clist, c)
		return clist, nil
	})
	defer monkey.Unpatch(pack.ToKitexCommentList)

	resp, err := CommentService.CommentList(&comment.CommentListRequest{Token: DemoUser.Token, VideoId: DemoVideo.Id})
	if err != nil {
		t.Fatal(err)
	}
	if resp == nil {
		t.Fatal("resp is nil")
	}
	if resp.GetCommentList() == nil || len(resp.GetCommentList()) == 0 {
		t.Fatal("comment_list is empty")
	}
	t.Logf("Comment_list response: %v", resp)
}
