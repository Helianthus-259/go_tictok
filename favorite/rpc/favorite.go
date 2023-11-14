package rpc

import (
	"context"
	"favorite/pkg/constants"
	"favorite/pkg/errno"
	"favorite/pkg/rpc-middleware"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	etcd "github.com/kitex-contrib/registry-etcd"
	rpcFavorite "rpc/kitex_gen/favorite"
	"rpc/kitex_gen/favorite/favoriteservice"
	rpcVideo "rpc/kitex_gen/video"
)

var (
	favoriteClient favoriteservice.Client
)

func initFavorite() {
	r, err := etcd.NewEtcdResolver([]string{constants.ETCDAddress})
	if err != nil {
		panic(err)
	}
	//p := provider.NewOpenTelemetryProvider(
	//	provider.WithServiceName(constants.ApiServiceName),
	//	provider.WithExportEndpoint(constants.ExportEndpoint),
	//	provider.WithInsecure(),
	//)
	//defer func(ctx context.Context, p provider.OtelProvider) {
	//	_ = p.Shutdown(ctx)
	//}(context.Background(), p)
	fc, err := favoriteservice.NewClient(
		constants.FavoriteServiceName,
		client.WithResolver(r),
		client.WithMuxConnection(1),
		client.WithMiddleware(rpc_middleware.CommonMiddleware),
		client.WithInstanceMW(rpc_middleware.ClientMiddleware),
		client.WithSuite(tracing.NewClientSuite()),
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: constants.FavoriteServiceName}),
	)
	if err != nil {
		panic(err)
	}
	favoriteClient = fc
}

// FavoriteList Rpc Call Get Favorite List
func FavoriteList(ctx context.Context, req *rpcFavorite.FavoriteListRequest) ([]*rpcVideo.Video, error) {
	resp, err := favoriteClient.FavoriteList(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg)
	}
	return resp.VideoList, nil
}

// FavoriteAction Rpc Call Favorite Somebody's Video
func FavoriteAction(ctx context.Context, req *rpcFavorite.FavoriteActionRequest) error {
	resp, err := favoriteClient.FavoriteAction(ctx, req)
	if err != nil {
		return err
	}
	if resp.BaseResp.StatusCode != 0 {
		return errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg)
	}
	return nil
}

// IsFavoriteVideos Rpc Call If Favorite VideoList
func IsFavoriteVideos(ctx context.Context, req *rpcFavorite.IsFavoriteVideosRequest) ([]bool, error) {
	resp, err := favoriteClient.IsFavoriteVideos(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg)
	}
	return resp.ManyIsFavorite, nil
}
