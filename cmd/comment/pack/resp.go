package pack

import (
	comment "github.com/HUST-MiniTiktok/mini_tiktok/kitex_gen/comment"
	"github.com/HUST-MiniTiktok/mini_tiktok/pkg/utils/conv"
)

func NewCommentActionResponse(err error) *comment.CommentActionResponse {
	return conv.ToKitexBaseResponse(err, &comment.CommentActionResponse{}).(*comment.CommentActionResponse)
}

func NewCommentListResponse(err error) *comment.CommentListResponse {
	return conv.ToKitexBaseResponse(err, &comment.CommentListResponse{}).(*comment.CommentListResponse)
}

func NewGetVideoCommentCountResponse(err error) *comment.GetVideoCommentCountResponse {
	return conv.ToKitexBaseResponse(err, &comment.GetVideoCommentCountResponse{}).(*comment.GetVideoCommentCountResponse)
}
