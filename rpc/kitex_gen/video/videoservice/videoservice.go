// Code generated by Kitex v0.7.3. DO NOT EDIT.

package videoservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	video "rpc/kitex_gen/video"
)

func serviceInfo() *kitex.ServiceInfo {
	return videoServiceServiceInfo
}

var videoServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "VideoService"
	handlerType := (*video.VideoService)(nil)
	methods := map[string]kitex.MethodInfo{
		"VideoFeed":             kitex.NewMethodInfo(videoFeedHandler, newVideoServiceVideoFeedArgs, newVideoServiceVideoFeedResult, false),
		"VideoPublish":          kitex.NewMethodInfo(videoPublishHandler, newVideoServiceVideoPublishArgs, newVideoServiceVideoPublishResult, false),
		"PublishList":           kitex.NewMethodInfo(publishListHandler, newVideoServicePublishListArgs, newVideoServicePublishListResult, false),
		"GetVideoInfo":          kitex.NewMethodInfo(getVideoInfoHandler, newVideoServiceGetVideoInfoArgs, newVideoServiceGetVideoInfoResult, false),
		"GetManyVideoInfos":     kitex.NewMethodInfo(getManyVideoInfosHandler, newVideoServiceGetManyVideoInfosArgs, newVideoServiceGetManyVideoInfosResult, false),
		"AddVideoFavoriteCount": kitex.NewMethodInfo(addVideoFavoriteCountHandler, newVideoServiceAddVideoFavoriteCountArgs, newVideoServiceAddVideoFavoriteCountResult, false),
		"SubVideoFavoriteCount": kitex.NewMethodInfo(subVideoFavoriteCountHandler, newVideoServiceSubVideoFavoriteCountArgs, newVideoServiceSubVideoFavoriteCountResult, false),
		"AddVideoCommentCount":  kitex.NewMethodInfo(addVideoCommentCountHandler, newVideoServiceAddVideoCommentCountArgs, newVideoServiceAddVideoCommentCountResult, false),
		"SubVideoCommentCount":  kitex.NewMethodInfo(subVideoCommentCountHandler, newVideoServiceSubVideoCommentCountArgs, newVideoServiceSubVideoCommentCountResult, false),
	}
	extra := map[string]interface{}{
		"PackageName":     "video",
		"ServiceFilePath": `idl\video.thrift`,
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

func videoFeedHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*video.VideoServiceVideoFeedArgs)
	realResult := result.(*video.VideoServiceVideoFeedResult)
	success, err := handler.(video.VideoService).VideoFeed(ctx, realArg.Request)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newVideoServiceVideoFeedArgs() interface{} {
	return video.NewVideoServiceVideoFeedArgs()
}

func newVideoServiceVideoFeedResult() interface{} {
	return video.NewVideoServiceVideoFeedResult()
}

func videoPublishHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*video.VideoServiceVideoPublishArgs)
	realResult := result.(*video.VideoServiceVideoPublishResult)
	success, err := handler.(video.VideoService).VideoPublish(ctx, realArg.Request)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newVideoServiceVideoPublishArgs() interface{} {
	return video.NewVideoServiceVideoPublishArgs()
}

func newVideoServiceVideoPublishResult() interface{} {
	return video.NewVideoServiceVideoPublishResult()
}

func publishListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*video.VideoServicePublishListArgs)
	realResult := result.(*video.VideoServicePublishListResult)
	success, err := handler.(video.VideoService).PublishList(ctx, realArg.Request)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newVideoServicePublishListArgs() interface{} {
	return video.NewVideoServicePublishListArgs()
}

func newVideoServicePublishListResult() interface{} {
	return video.NewVideoServicePublishListResult()
}

func getVideoInfoHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*video.VideoServiceGetVideoInfoArgs)
	realResult := result.(*video.VideoServiceGetVideoInfoResult)
	success, err := handler.(video.VideoService).GetVideoInfo(ctx, realArg.Request)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newVideoServiceGetVideoInfoArgs() interface{} {
	return video.NewVideoServiceGetVideoInfoArgs()
}

func newVideoServiceGetVideoInfoResult() interface{} {
	return video.NewVideoServiceGetVideoInfoResult()
}

func getManyVideoInfosHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*video.VideoServiceGetManyVideoInfosArgs)
	realResult := result.(*video.VideoServiceGetManyVideoInfosResult)
	success, err := handler.(video.VideoService).GetManyVideoInfos(ctx, realArg.Request)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newVideoServiceGetManyVideoInfosArgs() interface{} {
	return video.NewVideoServiceGetManyVideoInfosArgs()
}

func newVideoServiceGetManyVideoInfosResult() interface{} {
	return video.NewVideoServiceGetManyVideoInfosResult()
}

func addVideoFavoriteCountHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*video.VideoServiceAddVideoFavoriteCountArgs)
	realResult := result.(*video.VideoServiceAddVideoFavoriteCountResult)
	success, err := handler.(video.VideoService).AddVideoFavoriteCount(ctx, realArg.Request)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newVideoServiceAddVideoFavoriteCountArgs() interface{} {
	return video.NewVideoServiceAddVideoFavoriteCountArgs()
}

func newVideoServiceAddVideoFavoriteCountResult() interface{} {
	return video.NewVideoServiceAddVideoFavoriteCountResult()
}

func subVideoFavoriteCountHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*video.VideoServiceSubVideoFavoriteCountArgs)
	realResult := result.(*video.VideoServiceSubVideoFavoriteCountResult)
	success, err := handler.(video.VideoService).SubVideoFavoriteCount(ctx, realArg.Request)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newVideoServiceSubVideoFavoriteCountArgs() interface{} {
	return video.NewVideoServiceSubVideoFavoriteCountArgs()
}

func newVideoServiceSubVideoFavoriteCountResult() interface{} {
	return video.NewVideoServiceSubVideoFavoriteCountResult()
}

func addVideoCommentCountHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*video.VideoServiceAddVideoCommentCountArgs)
	realResult := result.(*video.VideoServiceAddVideoCommentCountResult)
	success, err := handler.(video.VideoService).AddVideoCommentCount(ctx, realArg.Request)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newVideoServiceAddVideoCommentCountArgs() interface{} {
	return video.NewVideoServiceAddVideoCommentCountArgs()
}

func newVideoServiceAddVideoCommentCountResult() interface{} {
	return video.NewVideoServiceAddVideoCommentCountResult()
}

func subVideoCommentCountHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*video.VideoServiceSubVideoCommentCountArgs)
	realResult := result.(*video.VideoServiceSubVideoCommentCountResult)
	success, err := handler.(video.VideoService).SubVideoCommentCount(ctx, realArg.Request)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newVideoServiceSubVideoCommentCountArgs() interface{} {
	return video.NewVideoServiceSubVideoCommentCountArgs()
}

func newVideoServiceSubVideoCommentCountResult() interface{} {
	return video.NewVideoServiceSubVideoCommentCountResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) VideoFeed(ctx context.Context, request *video.VideoFeedRequest) (r *video.VideoFeedResponse, err error) {
	var _args video.VideoServiceVideoFeedArgs
	_args.Request = request
	var _result video.VideoServiceVideoFeedResult
	if err = p.c.Call(ctx, "VideoFeed", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) VideoPublish(ctx context.Context, request *video.VideoPublishRequest) (r *video.VideoPublishResponse, err error) {
	var _args video.VideoServiceVideoPublishArgs
	_args.Request = request
	var _result video.VideoServiceVideoPublishResult
	if err = p.c.Call(ctx, "VideoPublish", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) PublishList(ctx context.Context, request *video.PublishListRequest) (r *video.PublishListResponse, err error) {
	var _args video.VideoServicePublishListArgs
	_args.Request = request
	var _result video.VideoServicePublishListResult
	if err = p.c.Call(ctx, "PublishList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetVideoInfo(ctx context.Context, request *video.GetVideoInfoRequest) (r *video.GetVideoInfoResponse, err error) {
	var _args video.VideoServiceGetVideoInfoArgs
	_args.Request = request
	var _result video.VideoServiceGetVideoInfoResult
	if err = p.c.Call(ctx, "GetVideoInfo", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetManyVideoInfos(ctx context.Context, request *video.GetManyVideoInfosRequest) (r *video.GetManyVideoInfosResponse, err error) {
	var _args video.VideoServiceGetManyVideoInfosArgs
	_args.Request = request
	var _result video.VideoServiceGetManyVideoInfosResult
	if err = p.c.Call(ctx, "GetManyVideoInfos", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) AddVideoFavoriteCount(ctx context.Context, request *video.AddVideoFavoriteCountRequest) (r *video.AddVideoFavoriteCountResponse, err error) {
	var _args video.VideoServiceAddVideoFavoriteCountArgs
	_args.Request = request
	var _result video.VideoServiceAddVideoFavoriteCountResult
	if err = p.c.Call(ctx, "AddVideoFavoriteCount", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) SubVideoFavoriteCount(ctx context.Context, request *video.SubVideoFavoriteCountRequest) (r *video.SubVideoFavoriteCountResponse, err error) {
	var _args video.VideoServiceSubVideoFavoriteCountArgs
	_args.Request = request
	var _result video.VideoServiceSubVideoFavoriteCountResult
	if err = p.c.Call(ctx, "SubVideoFavoriteCount", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) AddVideoCommentCount(ctx context.Context, request *video.AddVideoCommentCountRequest) (r *video.AddVideoCommentCountResponse, err error) {
	var _args video.VideoServiceAddVideoCommentCountArgs
	_args.Request = request
	var _result video.VideoServiceAddVideoCommentCountResult
	if err = p.c.Call(ctx, "AddVideoCommentCount", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) SubVideoCommentCount(ctx context.Context, request *video.SubVideoCommentCountRequest) (r *video.SubVideoCommentCountResponse, err error) {
	var _args video.VideoServiceSubVideoCommentCountArgs
	_args.Request = request
	var _result video.VideoServiceSubVideoCommentCountResult
	if err = p.c.Call(ctx, "SubVideoCommentCount", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
