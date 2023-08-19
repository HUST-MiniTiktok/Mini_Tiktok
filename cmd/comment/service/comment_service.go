package service

import (
	"context"

	db "github.com/HUST-MiniTiktok/mini_tiktok/cmd/comment/dal/db"
	"github.com/HUST-MiniTiktok/mini_tiktok/cmd/publish/rpc"
	comment "github.com/HUST-MiniTiktok/mini_tiktok/kitex_gen/comment"
	user "github.com/HUST-MiniTiktok/mini_tiktok/kitex_gen/user"
	"github.com/HUST-MiniTiktok/mini_tiktok/pkg/mw/jwt"
	"github.com/cloudwego/kitex/pkg/remote/trans/nphttp2/codes"
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
		msg := err.Error()
		return &comment.CommentActionResponse{
			// status_code = 7 表示鉴权失败
			StatusCode: int32(codes.PermissionDenied),
			StatusMsg:  &msg,
			Comment:    nil}, nil
	}

	if request.ActionType == 1 { // 1-发布评论，2-删除评论
		newcomment, err := db.NewComment(s.ctx, claim.ID, request.VideoId, *request.CommentText)
		if err != nil {
			msg := err.Error() //创建失败
			return &comment.CommentActionResponse{
				StatusCode: int32(codes.Internal),
				StatusMsg:  &msg,
				Comment:    nil}, nil
		}
		// 创建成功需要返回 comment类型的newcomment
		err_chan := make(chan error)
		comment_chan := make(chan *comment.Comment)
		author, err := rpc.UserRPC.User(s.ctx, &user.UserRequest{UserId: newcomment.UserId})
		if err != nil {
			err_chan <- err
		} else {
			comment_chan <- &comment.Comment{
				Id:         newcomment.ID,
				User:       author.User,
				Content:    newcomment.CommentText,
				CreateDate: newcomment.CreatedAt.Format("01-02"),
			}
		}

		select {
		case err := <-err_chan:
			msg := err.Error()
			resp = &comment.CommentActionResponse{
				StatusCode: int32(codes.Internal),
				StatusMsg:  &msg,
				Comment:    nil}
			return resp, err
		case comComment := <-comment_chan:
			return &comment.CommentActionResponse{
				StatusCode: int32(codes.OK),
				StatusMsg:  nil,
				Comment:    comComment}, nil
		}

	} else if request.ActionType == 2 {
		if err := db.DelComment(s.ctx, *request.CommentId, request.VideoId); err != nil {
			msg := err.Error()
			return &comment.CommentActionResponse{
				StatusCode: int32(codes.Internal),
				StatusMsg:  &msg,
				Comment:    nil}, nil
		}
		// 删除成功
		return &comment.CommentActionResponse{
			StatusCode: int32(codes.OK),
			StatusMsg:  nil,
			Comment:    nil}, nil

	} else {
		msg := "action_type error"
		return &comment.CommentActionResponse{
			StatusCode: int32(codes.InvalidArgument),
			StatusMsg:  &msg,
			Comment:    nil}, nil
	}
}

// CommentList implements the CommentServiceImpl interface.
func (s *CommentService) CommentList(ctx context.Context, request *comment.CommentListRequest) (resp *comment.CommentListResponse, err error) {

	claim, err := Jwt.ExtractClaims(request.Token)
	if err != nil || claim.ID == 0 {
		msg := err.Error()
		return &comment.CommentListResponse{
			// status_code = 7 表示鉴权失败
			StatusCode:  int32(codes.PermissionDenied),
			StatusMsg:   &msg,
			CommentList: nil}, err
	}

	Comments, err := db.GetVideoComments(ctx, request.VideoId)
	if err != nil {
		msg := err.Error()
		return &comment.CommentListResponse{
			StatusCode:  int32(codes.Internal),
			StatusMsg:   &msg,
			CommentList: nil}, err
	}

	if len(Comments) == 0 {
		return &comment.CommentListResponse{
			StatusCode:  int32(codes.OK),
			StatusMsg:   nil,
			CommentList: nil}, nil
	}

	var comComments []*comment.Comment
	err_chan := make(chan error)
	comment_chan := make(chan *comment.Comment)

	for _, db_comment := range Comments {
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

	for i := 0; i < len(Comments); i++ {
		select {
		case err := <-err_chan:
			msg := err.Error()
			resp = &comment.CommentListResponse{
				StatusCode:  int32(codes.OK),
				StatusMsg:   &msg,
				CommentList: nil}
			return resp, err
		case comComment := <-comment_chan:
			comComments = append(comComments, comComment)
		}
	}

	return &comment.CommentListResponse{
		StatusCode:  int32(codes.OK),
		StatusMsg:   nil,
		CommentList: comComments}, nil
}

func (s *CommentService) GetVideoCommentCount(ctx context.Context, request *comment.GetVideoCommentCountRequest) (resp *comment.GetVideoCommentCountResponse, er error) {
	count, err := db.GetVideoCommentCounts(s.ctx, request.VideoId)
	if err != nil {
		msg := "Get Video Comments Count Failed"
		return &comment.GetVideoCommentCountResponse{
			StatusCode:   int32(codes.Internal),
			StatusMsg:    &msg,
			CommentCount: 0}, err
	}
	return &comment.GetVideoCommentCountResponse{
		StatusCode:   int32(codes.OK),
		StatusMsg:    nil,
		CommentCount: count}, nil
}
