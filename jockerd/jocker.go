package jockerd

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	adapter "github.com/gwatts/gin-adapter"
	"github.com/jaronnie/jocker/jockerd/internal/config"
	"github.com/jaronnie/jocker/jockerd/internal/router"
	"github.com/jaronnie/jocker/jockerd/internal/server"
	"github.com/jaronnie/jocker/jockerd/internal/svc"
	"github.com/jaronnie/jocker/jockerd/jockerdpb"
	"github.com/spf13/viper"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func StartJockerdZrpcServer(configFile string) {
	var c config.Config
	conf.MustLoad(configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		jockerdpb.RegisterJockerdServer(grpcServer, server.NewJockerdServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}

func StartJockerdGatewayServer() error {
	mux := runtime.NewServeMux()

	httpAddress := viper.GetString("ListenOnHTTP")
	grpcAddress := viper.GetString("ListenOn")

	opts := []grpc.DialOption{grpc.WithInsecure()}
	if err := jockerdpb.RegisterJockerdHandlerFromEndpoint(context.Background(), mux, grpcAddress, opts); err != nil {
		return err
	}

	g := gin.New()

	// wrap grpc gateway handler
	handler := adapter.Wrap(func(h http.Handler) http.Handler {
		return mux
	})

	// load gin handler
	g = router.Load(g)

	g.Use(handler)

	s := &http.Server{
		Addr:    httpAddress,
		Handler: g,
	}

	fmt.Printf("Starting http server at %s...\n", httpAddress)
	return s.ListenAndServe()
}
