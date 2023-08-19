package pack

import (
	user "github.com/HUST-MiniTiktok/mini_tiktok/kitex_gen/user"
	"github.com/HUST-MiniTiktok/mini_tiktok/pkg/utils/conv"
)

func NewUserResponse(err error) *user.UserResponse {
	return conv.ToKitexBaseResponse(err, &user.UserResponse{}).(*user.UserResponse)
}

func NewUserRegisterResponse(err error) *user.UserRegisterResponse {
	return conv.ToKitexBaseResponse(err, &user.UserRegisterResponse{}).(*user.UserRegisterResponse)
}

func NewUserLoginResponse(err error) *user.UserLoginResponse {
	return conv.ToKitexBaseResponse(err, &user.UserLoginResponse{}).(*user.UserLoginResponse)
}

func NewCheckUserIsExistResponse(err error) *user.CheckUserIsExistResponse {
	return conv.ToKitexBaseResponse(err, &user.CheckUserIsExistResponse{}).(*user.CheckUserIsExistResponse)
}
