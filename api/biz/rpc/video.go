package rpc

import (
	"api/pkg/constants"
	mw "common-components/rpc-middleware"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	etcd "github.com/kitex-contrib/registry-etcd"
	"golang.org/x/net/context"
	"rpc/kitex_gen/video/videoservice"
)

var (
	//InteractionClient interactionservice.Client
	VideoClient videoservice.Client
	//ChatClient        chatservice.Client
	//FavorClient       favorservice.Client
	//Comment           commentservice.Client
)

func initVideo() {
	r, err := etcd.NewEtcdResolver([]string{constants.ETCDAddress})
	if err != nil {
		panic(err)
	}
	p := provider.NewOpenTelemetryProvider(
		provider.WithServiceName(constants.ApiServiceName),
		provider.WithExportEndpoint(constants.ExportEndpoint),
		provider.WithInsecure(),
	)
	defer func(ctx context.Context, p provider.OtelProvider) {
		_ = p.Shutdown(ctx)
	}(context.Background(), p)
	c, err := videoservice.NewClient(
		constants.VideoServiceName,
		client.WithResolver(r),
		client.WithMuxConnection(1),
		client.WithMiddleware(mw.CommonMiddleware),
		client.WithInstanceMW(mw.ClientMiddleware),
		client.WithSuite(tracing.NewClientSuite()),
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: constants.ApiServiceName}),
	)
	if err != nil {
		panic(err)
	}
	VideoClient = c
}
