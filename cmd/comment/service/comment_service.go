package service

import (
	"context"

	db "github.com/HUST-MiniTiktok/mini_tiktok/cmd/comment/dal/db"
	"github.com/HUST-MiniTiktok/mini_tiktok/cmd/comment/pack"
	"github.com/HUST-MiniTiktok/mini_tiktok/cmd/comment/rpc"
	comment "github.com/HUST-MiniTiktok/mini_tiktok/kitex_gen/comment"
	user "github.com/HUST-MiniTiktok/mini_tiktok/kitex_gen/user"
	"github.com/HUST-MiniTiktok/mini_tiktok/pkg/errno"
	"github.com/HUST-MiniTiktok/mini_tiktok/pkg/mw/jwt"
)

var (
	Jwt *jwt.JWT
)

func init() {
	Jwt = jwt.NewJWT()
}

type CommentService struct {
	ctx context.Context
}

func NewCommentService(ctx context.Context) *CommentService {
	return &CommentService{ctx: ctx}
}

// CommentAction implements the CommentServiceImpl interface.
func (s *CommentService) CommentAction(ctx context.Context, request *comment.CommentActionRequest) (resp *comment.CommentActionResponse, err error) {
	claim, err := Jwt.ExtractClaims(request.Token)
	if err != nil || claim.ID == 0 {
		return pack.NewCommentActionResponse(errno.AuthorizationFailedErr), err
	}

	if request.ActionType == 1 { // 1-发布评论，2-删除评论
		newcomment, err := db.NewComment(s.ctx, claim.ID, request.VideoId, *request.CommentText)
		if err != nil {
			return pack.NewCommentActionResponse(err), err
		}

		// 创建成功需要返回 comment类型的newcomment
		author, err := rpc.UserRPC.User(s.ctx, &user.UserRequest{UserId: newcomment.UserId})
		if err != nil {
			return pack.NewCommentActionResponse(err), err
		}

		comComment := &comment.Comment{
			Id:         newcomment.ID,
			User:       author.User,
			Content:    newcomment.CommentText,
			CreateDate: newcomment.CreatedAt.Format("01-02"),
		}

		resp = pack.NewCommentActionResponse(errno.Success)
		resp.Comment = comComment
		return resp, nil

	} else if request.ActionType == 2 {
		if err := db.DelComment(s.ctx, *request.CommentId, request.VideoId); err != nil {
			return pack.NewCommentActionResponse(err), err
		}
		// 删除成功
		return pack.NewCommentActionResponse(errno.Success), nil

	} else {
		// ActionType参数错误
		return pack.NewCommentActionResponse(errno.ParamErr), errno.ParamErr
	}
}

// CommentList implements the CommentServiceImpl interface.
func (s *CommentService) CommentList(ctx context.Context, request *comment.CommentListRequest) (resp *comment.CommentListResponse, err error) {

	claim, err := Jwt.ExtractClaims(request.Token)
	if err != nil || claim.ID == 0 {
		return pack.NewCommentListResponse(errno.AuthorizationFailedErr), err
	}

	db_comments, err := db.GetVideoComments(ctx, request.VideoId)
	if err != nil {
		return pack.NewCommentListResponse(err), err
	}

	kitex_comments := make([]*comment.Comment, 0, len(db_comments))
	err_chan := make(chan error)
	comment_chan := make(chan *comment.Comment)

	for _, db_comment := range db_comments {
		go func(db_comment *db.Comment) {
			author, err := rpc.UserRPC.User(s.ctx, &user.UserRequest{UserId: db_comment.UserId})
			if err != nil {
				err_chan <- err
			} else {
				comment_chan <- &comment.Comment{
					Id:         db_comment.ID,
					User:       author.User,
					Content:    db_comment.CommentText,
					CreateDate: db_comment.CreatedAt.Format("01-02"),
				}
			}
		}(db_comment)
	}

	for i := 0; i < len(db_comments); i++ {
		select {
		case err := <-err_chan:
			return pack.NewCommentListResponse(err), err
		case comComment := <-comment_chan:
			kitex_comments = append(kitex_comments, comComment)
		}
	}

	resp = pack.NewCommentListResponse(errno.Success)
	resp.CommentList = kitex_comments
	return resp, nil
}

func (s *CommentService) GetVideoCommentCount(ctx context.Context, request *comment.GetVideoCommentCountRequest) (resp *comment.GetVideoCommentCountResponse, er error) {
	count, err := db.GetVideoCommentCounts(s.ctx, request.VideoId)
	if err != nil {
		return pack.NewGetVideoCommentCountResponse(err), err
	}
	resp = pack.NewGetVideoCommentCountResponse(errno.Success)
	resp.CommentCount = count
	return resp, nil
}
