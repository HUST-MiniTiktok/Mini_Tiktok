package service

import (
	"context"

	db "github.com/HUST-MiniTiktok/mini_tiktok/cmd/user/dal/db"
	"github.com/HUST-MiniTiktok/mini_tiktok/cmd/user/pack"
	user "github.com/HUST-MiniTiktok/mini_tiktok/kitex_gen/user"
	"github.com/HUST-MiniTiktok/mini_tiktok/pkg/errno"
	"github.com/HUST-MiniTiktok/mini_tiktok/pkg/mw/crypt"
	"github.com/HUST-MiniTiktok/mini_tiktok/pkg/mw/jwt"
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
	db_user, err = db.GetUserById(ctx, request.UserId)
	if err != nil {
		return pack.NewUserResponse(err), err
	}
	if db_user == nil {
		return pack.NewUserResponse(errno.UserIsNotExistErr), errno.UserIsNotExistErr
	}

	resp_user, err := pack.ToKitexUser(ctx, request.Token, db_user)
	if err != nil {
		return pack.NewUserResponse(err), err
	}

	resp = pack.NewUserResponse(errno.Success)
	resp.User = resp_user
	return resp, nil
}

func (s *UserService) Register(ctx context.Context, request *user.UserRegisterRequest) (resp *user.UserRegisterResponse, err error) {
	db_user_ck, err := db.GetUserByUserName(ctx, request.Username)
	if err != nil {
		return pack.NewUserRegisterResponse(err), err
	}

	if db_user_ck != nil {
		return pack.NewUserRegisterResponse(errno.UserAlreadyExistErr), errno.UserAlreadyExistErr
	}

	encrypted_password, err := crypt.HashPassword(request.Password)
	if err != nil {
		return pack.NewUserRegisterResponse(err), err
	}

	db_user_new := db.User{
		UserName: request.Username,
		Password: encrypted_password,
	}

	user_id, err := db.CreateUser(ctx, &db_user_new)
	if err != nil {
		return pack.NewUserRegisterResponse(err), err
	}

	token, err := Jwt.CreateToken(jwt.UserClaims{ID: user_id})
	if err != nil {
		return pack.NewUserRegisterResponse(err), err
	}

	resp = pack.NewUserRegisterResponse(errno.Success)
	resp.UserId = user_id
	resp.Token = token
	return resp, nil
}

func (s *UserService) Login(ctx context.Context, request *user.UserLoginRequest) (resp *user.UserLoginResponse, err error) {
	db_user_ck, err := db.GetUserByUserName(ctx, request.Username)
	if err != nil {
		return pack.NewUserLoginResponse(err), err
	}

	if db_user_ck == nil {
		return pack.NewUserLoginResponse(errno.UserIsNotExistErr), errno.UserIsNotExistErr
	}

	if !crypt.VerifyPassword(request.Password, db_user_ck.Password) {
		return pack.NewUserLoginResponse(errno.UserPasswordErr), errno.UserPasswordErr
	}

	user_id := int64(db_user_ck.ID)
	token, err := Jwt.CreateToken(jwt.UserClaims{ID: user_id})
	if err != nil {
		return pack.NewUserLoginResponse(err), err
	}

	resp = pack.NewUserLoginResponse(errno.Success)
	resp.UserId = user_id
	resp.Token = token
	return resp, nil
}

func (s *UserService) CheckUserIsExist(ctx context.Context, request *user.CheckUserIsExistRequest) (resp *user.CheckUserIsExistResponse, err error) {
	is_exist, err := db.CheckUserById(ctx, request.UserId)

	if err != nil {
		return pack.NewCheckUserIsExistResponse(err), err
	}

	resp = pack.NewCheckUserIsExistResponse(errno.Success)
	resp.IsExist = is_exist
	return resp, nil
}
