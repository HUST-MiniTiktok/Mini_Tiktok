package pack

import (
	"context"

	"github.com/HUST-MiniTiktok/mini_tiktok/cmd/comment/client"
	"github.com/HUST-MiniTiktok/mini_tiktok/cmd/comment/dal/db"
	comment "github.com/HUST-MiniTiktok/mini_tiktok/kitex_gen/comment"
	user "github.com/HUST-MiniTiktok/mini_tiktok/kitex_gen/user"
)

// ToKitexComment converts a db.Comment to a kitex Comment.
func ToKitexComment(ctx context.Context, db_comment *db.Comment, curr_user_token string) (*comment.Comment, error) {
	author, err := client.UserRPC.User(ctx, &user.UserRequest{UserId: db_comment.UserId, Token: curr_user_token})
	if err != nil {
		return nil, err
	}
	return &comment.Comment{
		Id:         db_comment.ID,
		User:       author.User,
		Content:    db_comment.CommentText,
		CreateDate: db_comment.CreatedAt.Format("01-02"),
	}, nil
}

// ToKitexCommentList converts a slice of db.Comment to a slice of kitex Comment.
func ToKitexCommentList(ctx context.Context, db_comments []*db.Comment, curr_user_token string) ([]*comment.Comment, error) {
	kitex_comments := make([]*comment.Comment, 0, len(db_comments))
	err_chan := make(chan error)
	comment_chan := make(chan *comment.Comment)
	for _, db_comment := range db_comments {
		go func(db_comment *db.Comment) {
			kitex_comment, err := ToKitexComment(ctx, db_comment, curr_user_token)
			if err != nil {
				err_chan <- err
			} else {
				comment_chan <- kitex_comment
			}
		}(db_comment)
	}
	for i := 0; i < len(db_comments); i++ {
		select {
		case err := <-err_chan:
			return nil, err
		case kitex_comment := <-comment_chan:
			kitex_comments = append(kitex_comments, kitex_comment)
		}
	}
	return kitex_comments, nil
}
