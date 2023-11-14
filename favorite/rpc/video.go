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
	rpcVideo "rpc/kitex_gen/video"
	"rpc/kitex_gen/video/videoservice"
)

var (
	videoClient videoservice.Client
)

func initVideo() {
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
	vc, err := videoservice.NewClient(
		constants.VideoServiceName,
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
	videoClient = vc
}

// VideoFeed Rpc Call Video Feed
func VideoFeed(ctx context.Context, req *rpcVideo.VideoFeedRequest) ([]*rpcVideo.Video, *int64, error) {
	resp, err := videoClient.VideoFeed(ctx, req)
	if err != nil {
		return nil, nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, nil, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg)
	}

	return resp.VideoList, resp.NextTime, nil
}

// VideoPublish Rpc Call Video Publish
func VideoPublish(ctx context.Context, req *rpcVideo.VideoPublishRequest) error {
	resp, err := videoClient.VideoPublish(ctx, req)
	if err != nil {
		return err
	}
	if resp.BaseResp.StatusCode != 0 {
		return errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg)
	}

	return nil
}

// PublishList Rpc Call Get Video Publish List
func PublishList(ctx context.Context, req *rpcVideo.PublishListRequest) ([]*rpcVideo.Video, error) {
	resp, err := videoClient.PublishList(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg)
	}

	return resp.VideoList, nil
}

// GetVideoInfo Rpc Call Get VideoInfo
func GetVideoInfo(ctx context.Context, req *rpcVideo.GetVideoInfoRequest) (*rpcVideo.Video, error) {
	resp, err := videoClient.GetVideoInfo(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg)
	}

	return resp.VideoInfo, nil
}

// GetManyVideoInfos Rpc Call Get VideoList Info
func GetManyVideoInfos(ctx context.Context, req *rpcVideo.GetManyVideoInfosRequest) ([]*rpcVideo.Video, error) {
	resp, err := videoClient.GetManyVideoInfos(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg)
	}

	return resp.VideoInfos, nil
}

// AddVideoFavoriteCount Rpc Call  Add Video Favorite Count
func AddVideoFavoriteCount(ctx context.Context, req *rpcVideo.AddVideoFavoriteCountRequest) error {
	resp, err := videoClient.AddVideoFavoriteCount(ctx, req)
	if err != nil {
		return err
	}
	if resp.BaseResp.StatusCode != 0 {
		return errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg)
	}

	return nil
}

// SubVideoFavoriteCount Rpc Call  Sub Video Favorite Count
func SubVideoFavoriteCount(ctx context.Context, req *rpcVideo.SubVideoFavoriteCountRequest) error {
	resp, err := videoClient.SubVideoFavoriteCount(ctx, req)
	if err != nil {
		return err
	}
	if resp.BaseResp.StatusCode != 0 {
		return errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg)
	}

	return nil
}

// AddVideoCommentCount Rpc Call  Add Video Comment Count
func AddVideoCommentCount(ctx context.Context, req *rpcVideo.AddVideoCommentCountRequest) error {
	resp, err := videoClient.AddVideoCommentCount(ctx, req)
	if err != nil {
		return err
	}
	if resp.BaseResp.StatusCode != 0 {
		return errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg)
	}

	return nil
}

// SubVideoCommentCount Rpc Call  Sub Video Comment Count
func SubVideoCommentCount(ctx context.Context, req *rpcVideo.SubVideoCommentCountRequest) error {
	resp, err := videoClient.SubVideoCommentCount(ctx, req)
	if err != nil {
		return err
	}
	if resp.BaseResp.StatusCode != 0 {
		return errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg)
	}

	return nil
}
