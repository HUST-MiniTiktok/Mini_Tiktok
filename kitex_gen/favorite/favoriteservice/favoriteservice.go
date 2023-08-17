// Code generated by Kitex v0.7.0. DO NOT EDIT.

package favoriteservice

import (
	"context"
	favorite "github.com/HUST-MiniTiktok/mini_tiktok/kitex_gen/favorite"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

func serviceInfo() *kitex.ServiceInfo {
	return favoriteServiceServiceInfo
}

var favoriteServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "FavoriteService"
	handlerType := (*favorite.FavoriteService)(nil)
	methods := map[string]kitex.MethodInfo{
		"FavoriteAction":        kitex.NewMethodInfo(favoriteActionHandler, newFavoriteServiceFavoriteActionArgs, newFavoriteServiceFavoriteActionResult, false),
		"FavoriteList":          kitex.NewMethodInfo(favoriteListHandler, newFavoriteServiceFavoriteListArgs, newFavoriteServiceFavoriteListResult, false),
		"GetVideoFavoriteCount": kitex.NewMethodInfo(getVideoFavoriteCountHandler, newFavoriteServiceGetVideoFavoriteCountArgs, newFavoriteServiceGetVideoFavoriteCountResult, false),
		"CheckIsFavorite":       kitex.NewMethodInfo(checkIsFavoriteHandler, newFavoriteServiceCheckIsFavoriteArgs, newFavoriteServiceCheckIsFavoriteResult, false),
	}
	extra := map[string]interface{}{
		"PackageName":     "favorite",
		"ServiceFilePath": "../../idl/favorite.thrift",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.7.0",
		Extra:           extra,
	}
	return svcInfo
}

func favoriteActionHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*favorite.FavoriteServiceFavoriteActionArgs)
	realResult := result.(*favorite.FavoriteServiceFavoriteActionResult)
	success, err := handler.(favorite.FavoriteService).FavoriteAction(ctx, realArg.Request)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newFavoriteServiceFavoriteActionArgs() interface{} {
	return favorite.NewFavoriteServiceFavoriteActionArgs()
}

func newFavoriteServiceFavoriteActionResult() interface{} {
	return favorite.NewFavoriteServiceFavoriteActionResult()
}

func favoriteListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*favorite.FavoriteServiceFavoriteListArgs)
	realResult := result.(*favorite.FavoriteServiceFavoriteListResult)
	success, err := handler.(favorite.FavoriteService).FavoriteList(ctx, realArg.Request)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newFavoriteServiceFavoriteListArgs() interface{} {
	return favorite.NewFavoriteServiceFavoriteListArgs()
}

func newFavoriteServiceFavoriteListResult() interface{} {
	return favorite.NewFavoriteServiceFavoriteListResult()
}

func getVideoFavoriteCountHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*favorite.FavoriteServiceGetVideoFavoriteCountArgs)
	realResult := result.(*favorite.FavoriteServiceGetVideoFavoriteCountResult)
	success, err := handler.(favorite.FavoriteService).GetVideoFavoriteCount(ctx, realArg.Request)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newFavoriteServiceGetVideoFavoriteCountArgs() interface{} {
	return favorite.NewFavoriteServiceGetVideoFavoriteCountArgs()
}

func newFavoriteServiceGetVideoFavoriteCountResult() interface{} {
	return favorite.NewFavoriteServiceGetVideoFavoriteCountResult()
}

func checkIsFavoriteHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*favorite.FavoriteServiceCheckIsFavoriteArgs)
	realResult := result.(*favorite.FavoriteServiceCheckIsFavoriteResult)
	success, err := handler.(favorite.FavoriteService).CheckIsFavorite(ctx, realArg.Request)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newFavoriteServiceCheckIsFavoriteArgs() interface{} {
	return favorite.NewFavoriteServiceCheckIsFavoriteArgs()
}

func newFavoriteServiceCheckIsFavoriteResult() interface{} {
	return favorite.NewFavoriteServiceCheckIsFavoriteResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) FavoriteAction(ctx context.Context, request *favorite.FavoriteActionRequest) (r *favorite.FavoriteActionResponse, err error) {
	var _args favorite.FavoriteServiceFavoriteActionArgs
	_args.Request = request
	var _result favorite.FavoriteServiceFavoriteActionResult
	if err = p.c.Call(ctx, "FavoriteAction", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) FavoriteList(ctx context.Context, request *favorite.FavoriteListRequest) (r *favorite.FavoriteListResponse, err error) {
	var _args favorite.FavoriteServiceFavoriteListArgs
	_args.Request = request
	var _result favorite.FavoriteServiceFavoriteListResult
	if err = p.c.Call(ctx, "FavoriteList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetVideoFavoriteCount(ctx context.Context, request *favorite.GetVideoFavoriteCountRequest) (r *favorite.GetVideoFavoriteCountResponse, err error) {
	var _args favorite.FavoriteServiceGetVideoFavoriteCountArgs
	_args.Request = request
	var _result favorite.FavoriteServiceGetVideoFavoriteCountResult
	if err = p.c.Call(ctx, "GetVideoFavoriteCount", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) CheckIsFavorite(ctx context.Context, request *favorite.CheckIsFavoriteRequest) (r *favorite.CheckIsFavoriteResponse, err error) {
	var _args favorite.FavoriteServiceCheckIsFavoriteArgs
	_args.Request = request
	var _result favorite.FavoriteServiceCheckIsFavoriteResult
	if err = p.c.Call(ctx, "CheckIsFavorite", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}