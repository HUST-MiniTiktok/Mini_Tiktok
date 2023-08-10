// Code generated by Kitex v0.6.2. DO NOT EDIT.

package userservice

import (
	"context"
	user "github.com/HUST-MiniTiktok/mini_tiktok/biz/user/kitex_gen/user"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

func serviceInfo() *kitex.ServiceInfo {
	return userServiceServiceInfo
}

var userServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "UserService"
	handlerType := (*user.UserService)(nil)
	methods := map[string]kitex.MethodInfo{
		"Register":    kitex.NewMethodInfo(registerHandler, newUserServiceRegisterArgs, newUserServiceRegisterResult, false),
		"Login":       kitex.NewMethodInfo(loginHandler, newUserServiceLoginArgs, newUserServiceLoginResult, false),
		"GetUserById": kitex.NewMethodInfo(getUserByIdHandler, newUserServiceGetUserByIdArgs, newUserServiceGetUserByIdResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "user",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.6.2",
		Extra:           extra,
	}
	return svcInfo
}

func registerHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceRegisterArgs)
	realResult := result.(*user.UserServiceRegisterResult)
	success, err := handler.(user.UserService).Register(ctx, realArg.Request)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceRegisterArgs() interface{} {
	return user.NewUserServiceRegisterArgs()
}

func newUserServiceRegisterResult() interface{} {
	return user.NewUserServiceRegisterResult()
}

func loginHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceLoginArgs)
	realResult := result.(*user.UserServiceLoginResult)
	success, err := handler.(user.UserService).Login(ctx, realArg.Request)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceLoginArgs() interface{} {
	return user.NewUserServiceLoginArgs()
}

func newUserServiceLoginResult() interface{} {
	return user.NewUserServiceLoginResult()
}

func getUserByIdHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceGetUserByIdArgs)
	realResult := result.(*user.UserServiceGetUserByIdResult)
	success, err := handler.(user.UserService).GetUserById(ctx, realArg.Request)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceGetUserByIdArgs() interface{} {
	return user.NewUserServiceGetUserByIdArgs()
}

func newUserServiceGetUserByIdResult() interface{} {
	return user.NewUserServiceGetUserByIdResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) Register(ctx context.Context, request *user.UserRegisterRequest) (r *user.UserRegisterResponse, err error) {
	var _args user.UserServiceRegisterArgs
	_args.Request = request
	var _result user.UserServiceRegisterResult
	if err = p.c.Call(ctx, "Register", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) Login(ctx context.Context, request *user.UserLoginRequest) (r *user.UserLoginResponse, err error) {
	var _args user.UserServiceLoginArgs
	_args.Request = request
	var _result user.UserServiceLoginResult
	if err = p.c.Call(ctx, "Login", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetUserById(ctx context.Context, request *user.UserRequest) (r *user.UserResponse, err error) {
	var _args user.UserServiceGetUserByIdArgs
	_args.Request = request
	var _result user.UserServiceGetUserByIdResult
	if err = p.c.Call(ctx, "GetUserById", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
