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
	rpcComment "rpc/kitex_gen/comment"
	"rpc/kitex_gen/comment/commentservice"
)

var (
	commentClient commentservice.Client
)

func initComment() {
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
	uc, err := commentservice.NewClient(
		constants.CommentServiceName,
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
	commentClient = uc
}

// CommentAction Rpc Call Comment Somebody's Video
func CommentAction(ctx context.Context, req *rpcComment.CommentActionRequest) (*rpcComment.Comment, error) {
	resp, err := commentClient.CommentAction(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg)
	}
	return resp.Comment, nil
}

// CommentList Rpc Call Get Comment List
func CommentList(ctx context.Context, req *rpcComment.CommentListRequest) ([]*rpcComment.Comment, error) {
	resp, err := commentClient.CommentList(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg)
	}
	return resp.CommentList, nil
}