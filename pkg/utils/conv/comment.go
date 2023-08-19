package conv

import (
	hertz_comment "github.com/HUST-MiniTiktok/mini_tiktok/cmd/api/biz/model/comment"
	kitex_comment "github.com/HUST-MiniTiktok/mini_tiktok/kitex_gen/comment"
)

func ToHertzComment(comment *kitex_comment.Comment) *hertz_comment.Comment {
	return &hertz_comment.Comment{
		ID:         comment.Id,
		User:       ToHertzUser(comment.User),
		Content:    comment.Content,
		CreateDate: comment.CreateDate,
	}
}

func ToHertzCommentList(commentList []*kitex_comment.Comment) []*hertz_comment.Comment {
	hertzCommentList := make([]*hertz_comment.Comment, 0, len(commentList))
	for _, comment := range commentList {
		hertzCommentList = append(hertzCommentList, ToHertzComment(comment))
	}
	return hertzCommentList
}

func ToKitexCommentActionRequest(req *hertz_comment.CommentActionRequest) *kitex_comment.CommentActionRequest {
	return &kitex_comment.CommentActionRequest{
		Token:       req.Token,
		VideoId:     req.VideoID,
		ActionType:  req.ActionType,
		CommentText: req.CommentText,
		CommentId:   req.CommentID,
	}
}

func ToHertzCommentActionResponse(resp *kitex_comment.CommentActionResponse) *hertz_comment.CommentActionResponse {
	r := &hertz_comment.CommentActionResponse{
		StatusCode: resp.StatusCode,
		StatusMsg:  resp.StatusMsg,
	}
	if resp.Comment != nil {
		r.Comment = ToHertzComment(resp.Comment)
	}
	return r
}

func ToKitexCommentListRequest(req *hertz_comment.CommentListRequest) *kitex_comment.CommentListRequest {
	return &kitex_comment.CommentListRequest{
		Token:   req.Token,
		VideoId: req.VideoID,
	}
}

func ToHertzCommentListResponse(resp *kitex_comment.CommentListResponse) *hertz_comment.CommentListResponse {
	return &hertz_comment.CommentListResponse{
		StatusCode:  resp.StatusCode,
		StatusMsg:   resp.StatusMsg,
		CommentList: ToHertzCommentList(resp.CommentList),
	}
}
