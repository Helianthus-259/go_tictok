package main

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	kitexlogrus "github.com/kitex-contrib/obs-opentelemetry/logging/logrus"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	etcd "github.com/kitex-contrib/registry-etcd"
	"net"
	"rpc/kitex_gen/user/userservice"
	"user/dal"
	"user/pkg/constants"
	mw "user/pkg/rpc-middleware"
	"user/rpc"
)

func Init() {
	dal.Init()
	rpc.InitRpcClient()
	// klog init
	klog.SetLogger(kitexlogrus.NewLogger())
	klog.SetLevel(klog.LevelInfo)
}

func main() {
	r, err := etcd.NewEtcdRegistry([]string{constants.ETCDAddress})
	if err != nil {
		panic(err)
	}
	addr, err := net.ResolveTCPAddr(constants.TCP, constants.UserServiceAddr)
	if err != nil {
		panic(err)
	}
	Init()
	//p := provider.NewOpenTelemetryProvider(
	//	provider.WithServiceName(constants.UserServiceName),
	//	provider.WithExportEndpoint(constants.ExportEndpoint),
	//	provider.WithInsecure(),
	//)
	//defer func(ctx context.Context, p provider.OtelProvider) {
	//	_ = p.Shutdown(ctx)
	//}(context.Background(), p)
	svr := userservice.NewServer(new(UserServiceImpl),
		server.WithServiceAddr(addr),
		server.WithRegistry(r),
		server.WithLimit(&limit.Option{MaxConnections: 1000, MaxQPS: 100}),
		server.WithMuxTransport(),
		server.WithMiddleware(mw.CommonMiddleware),
		server.WithMiddleware(mw.ServerMiddleware),
		server.WithSuite(tracing.NewServerSuite()),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: constants.UserServiceName}),
	)
	err = svr.Run()
	if err != nil {
		klog.Fatal(err)
	}
}
