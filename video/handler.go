package main

import (
	"context"
	rpcVideo "rpc/kitex_gen/video"
)

// VideoServiceImpl implements the last service interface defined in the IDL.
type VideoServiceImpl struct{}

func (v VideoServiceImpl) VideoFeed(ctx context.Context, request *rpcVideo.VideoFeedRequest) (r *rpcVideo.VideoFeedResponse, err error) {
	//TODO implement me
	panic("implement me")
}

func (v VideoServiceImpl) VideoPublish(ctx context.Context, request *rpcVideo.VideoPublishRequest) (r *rpcVideo.VideoPublishResponse, err error) {
	//TODO implement me
	panic("implement me")
}

func (v VideoServiceImpl) PublishList(ctx context.Context, request *rpcVideo.PublishListRequest) (r *rpcVideo.PublishListResponse, err error) {
	//TODO implement me
	panic("implement me")
}

func (v VideoServiceImpl) GetVideoInfo(ctx context.Context, request *rpcVideo.GetVideoInfoRequest) (r *rpcVideo.GetVideoInfoResponse, err error) {
	//TODO implement me
	panic("implement me")
}

func (v VideoServiceImpl) GetManyVideoInfos(ctx context.Context, request *rpcVideo.GetManyVideoInfosRequest) (r *rpcVideo.GetManyVideoInfosResponse, err error) {
	//TODO implement me
	panic("implement me")
}

func (v VideoServiceImpl) AddVideoFavoriteCount(ctx context.Context, request *rpcVideo.AddVideoFavoriteCountRequest) (r *rpcVideo.AddVideoFavoriteCountResponse, err error) {
	//TODO implement me
	panic("implement me")
}

func (v VideoServiceImpl) SubVideoFavoriteCount(ctx context.Context, request *rpcVideo.SubVideoFavoriteCountRequest) (r *rpcVideo.SubVideoFavoriteCountResponse, err error) {
	//TODO implement me
	panic("implement me")
}

func (v VideoServiceImpl) AddVideoCommentCount(ctx context.Context, request *rpcVideo.AddVideoCommentCountRequest) (r *rpcVideo.AddVideoCommentCountResponse, err error) {
	//TODO implement me
	panic("implement me")
}

func (v VideoServiceImpl) SubVideoCommentCount(ctx context.Context, request *rpcVideo.SubVideoCommentCountRequest) (r *rpcVideo.SubVideoCommentCountResponse, err error) {
	//TODO implement me
	panic("implement me")
}
