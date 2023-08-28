package service

import (
	"context"

	db "github.com/HUST-MiniTiktok/mini_tiktok/cmd/comment/dal/db"
	"github.com/HUST-MiniTiktok/mini_tiktok/cmd/comment/pack"
	comment "github.com/HUST-MiniTiktok/mini_tiktok/kitex_gen/comment"
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

// CommentAction: publish a comment or delete a comment
func (s *CommentService) CommentAction(request *comment.CommentActionRequest) (resp *comment.CommentActionResponse, err error) {
	claim, err := Jwt.ExtractClaims(request.Token)
	if err != nil || claim.ID == 0 {
		return pack.NewCommentActionResponse(errno.AuthorizationFailedErr), err
	}
	// 1-发布评论，2-删除评论
	if request.ActionType == 1 {
		newcomment, err := db.NewComment(s.ctx, claim.ID, request.VideoId, *request.CommentText)
		if err != nil {
			return pack.NewCommentActionResponse(err), err
		}

		kitex_comment, err := pack.ToKitexComment(s.ctx, newcomment, request.Token)
		if err != nil {
			return pack.NewCommentActionResponse(err), err
		}

		resp = pack.NewCommentActionResponse(errno.Success)
		resp.Comment = kitex_comment
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

// CommentList: get comments of a video
func (s *CommentService) CommentList(request *comment.CommentListRequest) (resp *comment.CommentListResponse, err error) {

	claim, err := Jwt.ExtractClaims(request.Token)
	if err != nil || claim.ID == 0 {
		return pack.NewCommentListResponse(errno.AuthorizationFailedErr), err
	}

	db_comments, err := db.GetVideoComments(s.ctx, request.VideoId)
	if err != nil {
		return pack.NewCommentListResponse(err), err
	}

	kitex_comments, err := pack.ToKitexCommentList(s.ctx, db_comments, request.Token)
	if err != nil {
		return pack.NewCommentListResponse(err), err
	}

	resp = pack.NewCommentListResponse(errno.Success)
	resp.CommentList = kitex_comments
	return resp, nil
}

// GetVideoCommentCount implements the CommentServiceImpl interface.
func (s *CommentService) GetVideoCommentCount(request *comment.GetVideoCommentCountRequest) (resp *comment.GetVideoCommentCountResponse, er error) {
	count, err := db.GetVideoCommentCounts(s.ctx, request.VideoId)
	if err != nil {
		return pack.NewGetVideoCommentCountResponse(err), err
	}
	resp = pack.NewGetVideoCommentCountResponse(errno.Success)
	resp.CommentCount = count
	return resp, nil
}
