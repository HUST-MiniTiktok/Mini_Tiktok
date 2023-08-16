// Code generated by hertz generator.

package publish

import (
	"context"

	publish "github.com/HUST-MiniTiktok/mini_tiktok/cmd/api/biz/model/publish"
	rpc "github.com/HUST-MiniTiktok/mini_tiktok/rpc"
	"github.com/HUST-MiniTiktok/mini_tiktok/utils"
	"github.com/HUST-MiniTiktok/mini_tiktok/utils/conv"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/remote/trans/nphttp2/codes"
)

// PublishAction .
// @router /douyin/publish/action/ [POST]
func PublishAction(ctx context.Context, c *app.RequestContext) {
	klog.Info("1111111111111111111")
	var err error
	var req publish.PublishActionRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.JSON(consts.StatusBadRequest, utils.NewRespMap(int32(codes.InvalidArgument), err.Error()))
		return
	}
	rpc.
	klog.Infof("req: %+v", req)
	kitex_resp, err := rpc.PublishAction(ctx, conv.ToKitexPublishActionRequest(&req))

	if err == nil {
		c.JSON(consts.StatusOK, conv.ToHertzPublishActionResponse(kitex_resp))
	} else {
		c.JSON(consts.StatusOK, utils.NewRespMap(int32(codes.Internal), err.Error()))
	}
}

// PublishList .
// @router /douyin/publish/list/ [GET]
func PublishList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req publish.PublishListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.JSON(consts.StatusBadRequest, utils.NewRespMap(int32(codes.InvalidArgument), err.Error()))
		return
	}

	kitex_resp, err := rpc.PublishList(ctx, conv.ToKitexPublishListRequest(&req))

	if err == nil {
		c.JSON(consts.StatusOK, conv.ToHertzPublishListResponse(kitex_resp))
	} else {
		c.JSON(consts.StatusOK, utils.NewRespMap(int32(codes.Internal), err.Error()))
	}
}
