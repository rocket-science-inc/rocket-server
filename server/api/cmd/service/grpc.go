package service

import (
	"net"

	group "github.com/oklog/oklog/pkg/group"
	google_grpc "google.golang.org/grpc"
	opentracinggo "github.com/opentracing/opentracing-go"

	kit_grpc "github.com/go-kit/kit/transport/grpc"

	endpoint "rocket-server/server/api/pkg/endpoint"
	grpc "rocket-server/server/api/pkg/grpc/handler"
	pb "rocket-server/server/api/pkg/grpc/pb"
	log "github.com/go-kit/kit/log"
	level "github.com/go-kit/kit/log/level"
)

var grpcAddr = fs.String("grpc-addr", ":8082", "gRPC listen address")

func initGRPCHandler(endpoints endpoint.Endpoints, g *group.Group) {
	options := defaultGRPCOptions(logger, tracer)

	grpcServer := grpc.NewGRPCServer(endpoints, options)
	grpcListener, err := net.Listen("tcp", *grpcAddr)
	if err != nil {
		level.Error(logger).Log("transport", "gRPC", "during", "Listen", "err", err)
	}
	g.Add(func() error {
		level.Info(logger).Log("transport", "gRPC", "addr", *grpcAddr)
		baseServer := google_grpc.NewServer()
		pb.RegisterApiServer(baseServer, grpcServer)
		return baseServer.Serve(grpcListener)
	}, func(error) {
		grpcListener.Close()
	})

}

func defaultGRPCOptions(logger log.Logger, tracer opentracinggo.Tracer) map[string][]kit_grpc.ServerOption {
	options := map[string][]kit_grpc.ServerOption {

	}
	return options
}