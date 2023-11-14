// Code generated by hertz generator.

package main

import (
	"api/biz/rpc"
	"api/pkg/mw/jwt"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	hertzlogrus "github.com/hertz-contrib/obs-opentelemetry/logging/logrus"
	"github.com/hertz-contrib/obs-opentelemetry/tracing"
	"github.com/hertz-contrib/pprof"
)

func Init() {
	rpc.Init()
	jwt.InitJWT()
	// hlog init
	hlog.SetLogger(hertzlogrus.NewLogger())
	hlog.SetLevel(hlog.LevelInfo)
}

func main() {
	Init()
	tracer, cfg := tracing.NewServerTracer()
	h := server.New(
		server.WithHostPorts(":8080"),
		server.WithHandleMethodNotAllowed(true), // coordinate with NoMethod
		tracer,
	)
	// use pprof mw
	pprof.Register(h)
	// use otel mw
	h.Use(tracing.ServerMiddleware(cfg))
	register(h)
	h.Spin()
}
