// Code generated by Kitex v0.7.3. DO NOT EDIT.

package favoriteservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	favorite "rpc/kitex_gen/favorite"
)

func serviceInfo() *kitex.ServiceInfo {
	return favoriteServiceServiceInfo
}

var favoriteServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "FavoriteService"
	handlerType := (*favorite.FavoriteService)(nil)
	methods := map[string]kitex.MethodInfo{
		"FavoriteList":     kitex.NewMethodInfo(favoriteListHandler, newFavoriteServiceFavoriteListArgs, newFavoriteServiceFavoriteListResult, false),
		"FavoriteAction":   kitex.NewMethodInfo(favoriteActionHandler, newFavoriteServiceFavoriteActionArgs, newFavoriteServiceFavoriteActionResult, false),
		"IsFavoriteVideos": kitex.NewMethodInfo(isFavoriteVideosHandler, newFavoriteServiceIsFavoriteVideosArgs, newFavoriteServiceIsFavoriteVideosResult, false),
	}
	extra := map[string]interface{}{
		"PackageName":     "favorite",
		"ServiceFilePath": `idl\favorite.thrift`,
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.7.3",
		Extra:           extra,
	}
	return svcInfo
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

func isFavoriteVideosHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*favorite.FavoriteServiceIsFavoriteVideosArgs)
	realResult := result.(*favorite.FavoriteServiceIsFavoriteVideosResult)
	success, err := handler.(favorite.FavoriteService).IsFavoriteVideos(ctx, realArg.Request)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newFavoriteServiceIsFavoriteVideosArgs() interface{} {
	return favorite.NewFavoriteServiceIsFavoriteVideosArgs()
}

func newFavoriteServiceIsFavoriteVideosResult() interface{} {
	return favorite.NewFavoriteServiceIsFavoriteVideosResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
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

func (p *kClient) FavoriteAction(ctx context.Context, request *favorite.FavoriteActionRequest) (r *favorite.FavoriteActionResponse, err error) {
	var _args favorite.FavoriteServiceFavoriteActionArgs
	_args.Request = request
	var _result favorite.FavoriteServiceFavoriteActionResult
	if err = p.c.Call(ctx, "FavoriteAction", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) IsFavoriteVideos(ctx context.Context, request *favorite.IsFavoriteVideosRequest) (r *favorite.IsFavoriteVideosResponse, err error) {
	var _args favorite.FavoriteServiceIsFavoriteVideosArgs
	_args.Request = request
	var _result favorite.FavoriteServiceIsFavoriteVideosResult
	if err = p.c.Call(ctx, "IsFavoriteVideos", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}