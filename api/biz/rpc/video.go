package rpc

import (
	"api/pkg/constants"
	"api/pkg/errno"
	"api/pkg/mw/kitex-mw"
	"context"
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
		client.WithMiddleware(kitex_mw.CommonMiddleware),
		client.WithInstanceMW(kitex_mw.ClientMiddleware),
		client.WithSuite(tracing.NewClientSuite()),
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: constants.ApiServiceName}),
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
