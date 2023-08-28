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

// GetUserById: get a user by user id
func (s *UserService) GetUserById(request *user.UserRequest) (resp *user.UserResponse, err error) {
	var db_user *db.User
	db_user, err = db.GetUserById(s.ctx, request.UserId)
	if err != nil {
		return pack.NewUserResponse(err), err
	}
	if db_user == nil {
		return pack.NewUserResponse(errno.UserIsNotExistErr), errno.UserIsNotExistErr
	}

	resp_user, err := pack.ToKitexUser(s.ctx, request.Token, db_user)
	if err != nil {
		return pack.NewUserResponse(err), err
	}

	resp = pack.NewUserResponse(errno.Success)
	resp.User = resp_user
	return resp, nil
}

// Register: register a new user
func (s *UserService) Register(request *user.UserRegisterRequest) (resp *user.UserRegisterResponse, err error) {
	// check user name is already exist
	db_user_ck, err := db.GetUserByUserName(s.ctx, request.Username)
	if err != nil {
		return pack.NewUserRegisterResponse(err), err
	}
	if db_user_ck != nil {
		return pack.NewUserRegisterResponse(errno.UserAlreadyExistErr), errno.UserAlreadyExistErr
	}

	// encrypt password
	encrypted_password, err := crypt.HashPassword(request.Password)
	if err != nil {
		return pack.NewUserRegisterResponse(err), err
	}

	db_user_new := db.User{
		UserName: request.Username,
		Password: encrypted_password,
	}

	user_id, err := db.CreateUser(s.ctx, &db_user_new)
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

// Login: login a user
func (s *UserService) Login(request *user.UserLoginRequest) (resp *user.UserLoginResponse, err error) {
	db_user_ck, err := db.GetUserByUserName(s.ctx, request.Username)
	if err != nil {
		return pack.NewUserLoginResponse(err), err
	}

	if db_user_ck == nil {
		return pack.NewUserLoginResponse(errno.UserIsNotExistErr), errno.UserIsNotExistErr
	}
	// check password
	if !crypt.VerifyPassword(request.Password, db_user_ck.Password) {
		return pack.NewUserLoginResponse(errno.UserPasswordErr), errno.UserPasswordErr
	}
	// generate token
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

// CheckUserIsExist: check if a user exists by user id
func (s *UserService) CheckUserIsExist(request *user.CheckUserIsExistRequest) (resp *user.CheckUserIsExistResponse, err error) {
	is_exist, err := db.CheckUserById(s.ctx, request.UserId)

	if err != nil {
		return pack.NewCheckUserIsExistResponse(err), err
	}

	resp = pack.NewCheckUserIsExistResponse(errno.Success)
	resp.IsExist = is_exist
	return resp, nil
}
