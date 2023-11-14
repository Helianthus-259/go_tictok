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
	rpcRelation "rpc/kitex_gen/relation"
	"rpc/kitex_gen/relation/relationservice"
	rpcUser "rpc/kitex_gen/user"
)

var (
	relationClient relationservice.Client
)

func initRelation() {
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
	rc, err := relationservice.NewClient(
		constants.RelationServiceName,
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
	relationClient = rc
}

// FollowAction Rpc Call Follow Somebody
func FollowAction(ctx context.Context, req *rpcRelation.FollowActionRequest) error {
	resp, err := relationClient.FollowAction(ctx, req)
	if err != nil {
		return err
	}
	if resp.BaseResp.StatusCode != 0 {
		return errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg)
	}
	return nil
}

// FollowList Rpc Call Get Follow List
func FollowList(ctx context.Context, req *rpcRelation.FollowListRequest) ([]*rpcUser.User, error) {
	resp, err := relationClient.FollowList(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg)
	}
	return resp.UserList, nil
}

// FansList Rpc Call Get Fans List
func FansList(ctx context.Context, req *rpcRelation.FansListRequest) ([]*rpcUser.User, error) {
	resp, err := relationClient.FansList(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg)
	}
	return resp.UserList, nil
}

// FriendList Rpc Call Get Friend List
func FriendList(ctx context.Context, req *rpcRelation.FriendListRequest) ([]*rpcRelation.FriendUser, error) {
	resp, err := relationClient.FriendList(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg)
	}
	return resp.UserList, nil
}

// IsFollowTarget Rpc Call Check Is Follow Target User
func IsFollowTarget(ctx context.Context, req *rpcRelation.IsFollowTargetRequest) (bool, error) {
	resp, err := relationClient.IsFollowTarget(ctx, req)
	if err != nil {
		return false, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return false, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg)
	}
	return resp.Exist, nil
}

// IsFollowManyTargets Rpc Call Check Is Follow TargetList User
func IsFollowManyTargets(ctx context.Context, req *rpcRelation.IsFollowManyTargetsRequest) ([]bool, error) {
	resp, err := relationClient.IsFollowManyTargets(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg)
	}
	return resp.ManyExist, nil
}

// IsFriend Rpc Call Check Is Friend
func IsFriend(ctx context.Context, req *rpcRelation.IsFriendRequest) (bool, error) {
	resp, err := relationClient.IsFriend(ctx, req)
	if err != nil {
		return false, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return false, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg)
	}
	return resp.IsFriend, nil
}
