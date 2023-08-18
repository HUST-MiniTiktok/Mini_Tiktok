package main

import (
	"context"

	service "github.com/HUST-MiniTiktok/mini_tiktok/cmd/user/service"
	user "github.com/HUST-MiniTiktok/mini_tiktok/kitex_gen/user"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// User implements the UserServiceImpl interface.
func (s *UserServiceImpl) User(ctx context.Context, request *user.UserRequest) (resp *user.UserResponse, err error) {
	user_service := service.NewUserService(ctx)
	resp, err = user_service.GetUserById(ctx, request)
	return
}

// Register implements the UserServiceImpl interface.
func (s *UserServiceImpl) Register(ctx context.Context, request *user.UserRegisterRequest) (resp *user.UserRegisterResponse, err error) {
	user_service := service.NewUserService(ctx)
	resp, err = user_service.Register(ctx, request)
	return
}

// Login implements the UserServiceImpl interface.
func (s *UserServiceImpl) Login(ctx context.Context, request *user.UserLoginRequest) (resp *user.UserLoginResponse, err error) {
	user_service := service.NewUserService(ctx)
	resp, err = user_service.Login(ctx, request)
	return
}

// CheckUserIsExist implements the UserServiceImpl interface.
func (s *UserServiceImpl) CheckUserIsExist(ctx context.Context, request *user.CheckUserIsExistRequest) (resp *user.CheckUserIsExistResponse, err error) {
	user_service := service.NewUserService(ctx)
	resp, err = user_service.CheckUserIsExist(ctx, request)
	return
}
