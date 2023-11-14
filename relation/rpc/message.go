package rpc

import (
	"context"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	etcd "github.com/kitex-contrib/registry-etcd"
	"relation/pkg/constants"
	"relation/pkg/errno"
	"relation/pkg/rpc-middleware"
	rpcMessage "rpc/kitex_gen/message"
	"rpc/kitex_gen/message/messageservice"
)

var (
	messageClient messageservice.Client
)

func initMessage() {
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
	mc, err := messageservice.NewClient(
		constants.MessageServiceName,
		client.WithResolver(r),
		client.WithMuxConnection(1),
		client.WithMiddleware(rpc_middleware.CommonMiddleware),
		client.WithInstanceMW(rpc_middleware.ClientMiddleware),
		client.WithSuite(tracing.NewClientSuite()),
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: constants.RelationServiceName}),
	)
	if err != nil {
		panic(err)
	}
	messageClient = mc
}

// MessageList Rpc Call Get Message List
func MessageList(ctx context.Context, req *rpcMessage.MessageListRequest) ([]*rpcMessage.Message, error) {
	resp, err := messageClient.MessageList(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg)
	}
	return resp.MessageList, nil
}

// ChatAction Rpc Call To Send Message
func ChatAction(ctx context.Context, req *rpcMessage.ChatActionRequest) error {
	resp, err := messageClient.ChatAction(ctx, req)
	if err != nil {
		return err
	}
	if resp.BaseResp.StatusCode != 0 {
		return errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg)
	}
	return nil
}

// GetFriendLatestMessage Rpc Call Get Friend Latest Message
func GetFriendLatestMessage(ctx context.Context, req *rpcMessage.GetFriendLatestMessageRequest) ([]string, []int32, error) {
	resp, err := messageClient.GetFriendLatestMessage(ctx, req)
	if err != nil {
		return nil, nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, nil, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg)
	}
	return resp.MessageList, resp.MsgTypeList, nil
}
