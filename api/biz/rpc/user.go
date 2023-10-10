package rpc

import (
	"api/pkg/constants"
	"common-components/errno"
	mw "common-components/rpc-middleware"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	etcd "github.com/kitex-contrib/registry-etcd"
	"golang.org/x/net/context"
	rpcUser "rpc/kitex_gen/user"
	"rpc/kitex_gen/user/userservice"
)

var (
	userClient userservice.Client
	//InteractionClient interactionservice.Client
	//VideoClient       videoservice.Client
	//ChatClient        chatservice.Client
	//FavorClient       favorservice.Client
	//Comment           commentservice.Client
)

func initUser() {
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
	c, err := userservice.NewClient(
		constants.UserServiceName,
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
	userClient = c
}

// UserRegister Rpc Call Register User
func UserRegister(ctx context.Context, req *rpcUser.RegisterRequest) error {
	resp, err := userClient.Register(ctx, req)
	if err != nil {
		return err
	}
	if resp.BaseResp.StatusCode != 0 {
		return errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg)
	}
	return nil
}

// UserIndex Rpc Call Request User Index
func UserIndex(ctx context.Context, req *rpcUser.UserIndexRequest) (*rpcUser.User, error) {
	resp, err := userClient.UserIndex(ctx, req)
	if err != nil {
		return resp.User, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return resp.User, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg)
	}
	return resp.User, nil
}

// Login Rpc Call Request User Login
func Login(ctx context.Context, req *rpcUser.LoginRequest) (int64, error) {
	resp, err := userClient.Login(ctx, req)
	if err != nil {
		return 0, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return 0, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg)
	}
	return resp.UserId, nil
}
