// Code generated by hertz generator.

package user

import (
	"context"

	user "github.com/HUST-MiniTiktok/mini_tiktok/cmd/api/biz/model/user"
	rpc "github.com/HUST-MiniTiktok/mini_tiktok/cmd/api/biz/rpc"
	"github.com/HUST-MiniTiktok/mini_tiktok/util"
	"github.com/HUST-MiniTiktok/mini_tiktok/util/conv"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/cloudwego/kitex/pkg/remote/trans/nphttp2/codes"
)

// User .
// @router /douyin/user/ [GET]
func User(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.UserRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.JSON(consts.StatusBadRequest, util.NewRespMap(int32(codes.InvalidArgument), err.Error()))
		return
	}

	kitex_resp, err := rpc.UserRPC.User(ctx, conv.ToKitexUserRequest(&req))

	if err == nil {
		c.JSON(consts.StatusOK, conv.ToHertzUserResponse(kitex_resp))
	} else {
		c.JSON(consts.StatusOK, util.NewRespMap(int32(codes.Internal), err.Error()))
	}
}

// Register .
// @router /douyin/user/register/ [POST]
func Register(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.UserRegisterRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.JSON(consts.StatusBadRequest, util.NewRespMap(int32(codes.InvalidArgument), err.Error()))
		return
	}

	kitex_resp, err := rpc.UserRPC.Register(ctx, conv.ToKitexUserRegisterRequest(&req))

	if err == nil {
		c.JSON(consts.StatusOK, conv.ToHertzUserRegisterResponse(kitex_resp))
	} else {
		c.JSON(consts.StatusOK, util.NewRespMap(int32(codes.Internal), err.Error()))
	}
}

// Login .
// @router /douyin/user/login/ [POST]
func Login(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.UserLoginRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.JSON(consts.StatusBadRequest, util.NewRespMap(int32(codes.InvalidArgument), err.Error()))
		return
	}

	kitex_resp, err := rpc.UserRPC.Login(ctx, conv.ToKitexUserLoginRequest(&req))

	if err == nil {
		c.JSON(consts.StatusOK, conv.ToHertzUserLoginResponse(kitex_resp))
	} else {
		c.JSON(consts.StatusOK, util.NewRespMap(int32(codes.Internal), err.Error()))
	}
}
