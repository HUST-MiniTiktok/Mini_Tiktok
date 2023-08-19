package service

import (
	"context"

	db "github.com/HUST-MiniTiktok/mini_tiktok/cmd/user/dal/db"
	common "github.com/HUST-MiniTiktok/mini_tiktok/kitex_gen/common"
	user "github.com/HUST-MiniTiktok/mini_tiktok/kitex_gen/user"
	crypt "github.com/HUST-MiniTiktok/mini_tiktok/pkg/mw/crypt"
	jwt "github.com/HUST-MiniTiktok/mini_tiktok/pkg/mw/jwt"
	"github.com/cloudwego/kitex/pkg/remote/trans/nphttp2/codes"
)

var (
	Jwt *jwt.JWT
)

func init() {
	Jwt = jwt.NewJWT()
}

type UserService struct {
	ctx context.Context
}

func NewUserService(ctx context.Context) *UserService {
	return &UserService{ctx: ctx}
}

func (s *UserService) GetUserById(ctx context.Context, request *user.UserRequest) (resp *user.UserResponse, err error) {
	var db_user *db.User
	var resp_user *common.User
	db_user, err = db.GetUserById(ctx, request.UserId)
	if err != nil {
		errMsg := "db user not found"
		resp = &user.UserResponse{StatusCode: int32(codes.NotFound), StatusMsg: &errMsg, User: nil}
		return
	}
	resp_user = &common.User{
		Id:              db_user.ID,
		Name:            db_user.UserName,
		Avatar:          &db_user.Avatar,
		BackgroundImage: &db_user.BackgroundImage,
		Signature:       &db_user.Signature,
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

	token, err := Jwt.CreateToken(jwt.UserClaims{ID: user_id})
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
		errMsg := "db check user name failed"
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
		resp = &user.UserLoginResponse{StatusCode: int32(codes.PermissionDenied), StatusMsg: &errMsg}
		return
	}

	user_id := int64(db_user_ck.ID)
	token, err := Jwt.CreateToken(jwt.UserClaims{ID: user_id})
	if err != nil {
		errMsg := "create token failed"
		resp = &user.UserLoginResponse{StatusCode: int32(codes.Internal), StatusMsg: &errMsg}
		return
	}

	return &user.UserLoginResponse{StatusCode: int32(codes.OK), StatusMsg: nil, UserId: user_id, Token: token}, nil
}

func (s *UserService) CheckUserIsExist(ctx context.Context, request *user.CheckUserIsExistRequest) (resp *user.CheckUserIsExistResponse, err error) {
	is_exist, err := db.CheckUserById(ctx, request.UserId)

	if err != nil {
		errMsg := "db check user id failed"
		resp = &user.CheckUserIsExistResponse{StatusCode: int32(codes.Internal), StatusMsg: &errMsg}
		return
	}

	return &user.CheckUserIsExistResponse{StatusCode: int32(codes.OK), StatusMsg: nil, IsExist: is_exist}, nil
}
