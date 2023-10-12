package main

import (
	"context"

	service "github.com/HUST-MiniTiktok/mini_tiktok/cmd/comment/service"
	comment "github.com/HUST-MiniTiktok/mini_tiktok/kitex_gen/comment"
)

// CommentServiceImpl implements the last service interface defined in the IDL.
type CommentServiceImpl struct{}

// CommentAction implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) CommentAction(ctx context.Context, request *comment.CommentActionRequest) (resp *comment.CommentActionResponse, err error) {
	comment_service := service.NewCommentService(ctx)
	resp, err = comment_service.CommentAction(request)
	return
}

// CommentList implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) CommentList(ctx context.Context, request *comment.CommentListRequest) (resp *comment.CommentListResponse, err error) {
	comment_service := service.NewCommentService(ctx)
	resp, err = comment_service.CommentList(request)
	return
}

// GetVideoCommentCount implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) GetVideoCommentCount(ctx context.Context, request *comment.GetVideoCommentCountRequest) (resp *comment.GetVideoCommentCountResponse, err error) {
	comment_service := service.NewCommentService(ctx)
	resp, err = comment_service.GetVideoCommentCount(request)
	return
}

// GetVideoCommentListCount implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) GetVideoCommentListCount(ctx context.Context, request *comment.GetVideoCommentCountListRequest) (resp *comment.GetVideoCommentCountListResponse, err error) {
	comment_service := service.NewCommentService(ctx)
	resp, err = comment_service.GetVideoCommentCountList(request)
	return
}
