package main

import (
	"context"
	rpcFavorite "rpc/kitex_gen/favorite"
)

// FavoriteServiceImpl implements the last service interface defined in the IDL.
type FavoriteServiceImpl struct{}

func (f FavoriteServiceImpl) FavoriteList(ctx context.Context, request *rpcFavorite.FavoriteListRequest) (r *rpcFavorite.FavoriteListResponse, err error) {
	//TODO implement me
	panic("implement me")
}

func (f FavoriteServiceImpl) FavoriteAction(ctx context.Context, request *rpcFavorite.FavoriteActionRequest) (r *rpcFavorite.FavoriteActionResponse, err error) {
	//TODO implement me
	panic("implement me")
}

func (f FavoriteServiceImpl) IsFavoriteVideos(ctx context.Context, request *rpcFavorite.IsFavoriteVideosRequest) (r *rpcFavorite.IsFavoriteVideosResponse, err error) {
	//TODO implement me
	panic("implement me")
}
