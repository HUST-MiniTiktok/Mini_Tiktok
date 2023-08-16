package service

import (
	"context"

	db "github.com/HUST-MiniTiktok/mini_tiktok/cmd/user/dal/db"
	user "github.com/HUST-MiniTiktok/mini_tiktok/cmd/user/kitex_gen/user"
	crypt "github.com/HUST-MiniTiktok/mini_tiktok/mw/crypt"
	jwt "github.com/HUST-MiniTiktok/mini_tiktok/mw/jwt"
	"github.com/cloudwego/kitex/pkg/remote/trans/nphttp2/codes"
)

type UserService struct {
	ctx context.Context
}

func NewUserService(ctx context.Context) *UserService {
	return &UserService{ctx: ctx}
}

func (s *UserService) GetUserById(ctx context.Context, request *user.UserRequest) (resp *user.UserResponse, err error) {
	var db_user *db.User
	var resp_user *user.User
	db_user, err = db.GetUserById(ctx, request.UserId)
	if err != nil {
		errMsg := "db user not found"
		resp = &user.UserResponse{StatusCode: int32(codes.NotFound), StatusMsg: &errMsg, User: nil}
		return
	}
	resp_user = &user.User{
		Id:              int64(db_user.ID),
		Name:            db_user.UserName,
		Avatar:          &db_user.Avatar,
		BackgroundImage: &db_user.BackgroundImage,
		Signature:       &db_user.Signature,
		FavoriteCount:   new(int64),
		TotalFavorited:  new(int64),
		FollowCount:     new(int64),
		FollowerCount:   new(int64),
		IsFollow:        false,
	}
	return &user.UserResponse{StatusCode: int32(codes.OK), StatusMsg: nil, User: resp_user}, nil
}

func (s *UserService) Register(ctx context.Context, request *user.UserRegisterRequest) (resp *user.UserRegisterResponse, err error) {
	db_user_ck, err := db.GetUserByUserName(ctx, request.Username)
	if err != nil {
		errMsg := "db check username failed"
		resp = &user.UserRegisterResponse{StatusCode: int32(codes.Internal), StatusMsg: &errMsg}
		return
	}

	if db_user_ck != nil {
		errMsg := "username already exists"
		resp = &user.UserRegisterResponse{StatusCode: int32(codes.AlreadyExists), StatusMsg: &errMsg}
		return
	}

	encrypted_password, err := crypt.HashPassword(request.Password)
	if err != nil {
		errMsg := "crypt password failed"
		resp = &user.UserRegisterResponse{StatusCode: int32(codes.Internal), StatusMsg: &errMsg}
	}

	db_user_new := db.User{
		UserName: request.Username,
		Password: encrypted_password,
	}

	user_id, err := db.CreateUser(ctx, &db_user_new)
	if err != nil {
		errMsg := "db create user failed"
		resp = &user.UserRegisterResponse{StatusCode: int32(codes.Internal), StatusMsg: &errMsg}
		return
	}

	token, err := jwt.GetJwt().CreateToken(jwt.UserClaims{ID: user_id})
	if err != nil {
		errMsg := "create token failed"
		resp = &user.UserRegisterResponse{StatusCode: int32(codes.Internal), StatusMsg: &errMsg}
		return
	}

	return &user.UserRegisterResponse{StatusCode: int32(codes.OK), StatusMsg: nil, UserId: user_id, Token: token}, nil
}

func (s *UserService) Login(ctx context.Context, request *user.UserLoginRequest) (resp *user.UserLoginResponse, err error) {
	db_user_ck, err := db.GetUserByUserName(ctx, request.Username)
	if err != nil {
		errMsg := "db check username failed"
		resp = &user.UserLoginResponse{StatusCode: int32(codes.Internal), StatusMsg: &errMsg}
		return
	}

	if db_user_ck == nil {
		errMsg := "username not exists"
		resp = &user.UserLoginResponse{StatusCode: int32(codes.NotFound), StatusMsg: &errMsg}
		return
	}

	if !crypt.VerifyPassword(request.Password, db_user_ck.Password) {
		errMsg := "password incorrect"
		resp = &user.UserLoginResponse{StatusCode: int32(codes.Unauthenticated), StatusMsg: &errMsg}
		return
	}

	user_id := int64(db_user_ck.ID)
	token, err := jwt.GetJwt().CreateToken(jwt.UserClaims{ID: user_id})
	if err != nil {
		errMsg := "create token failed"
		resp = &user.UserLoginResponse{StatusCode: int32(codes.Internal), StatusMsg: &errMsg}
		return
	}

	return &user.UserLoginResponse{StatusCode: int32(codes.OK), StatusMsg: nil, UserId: user_id, Token: token}, nil
}
