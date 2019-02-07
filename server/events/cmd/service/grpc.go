package service

import (
	"net"

	group "github.com/oklog/oklog/pkg/group"
	grpcgo "google.golang.org/grpc"
	opentracinggo "github.com/opentracing/opentracing-go"

	log "github.com/go-kit/kit/log"
	opentracing "github.com/go-kit/kit/tracing/opentracing"
	grpckit "github.com/go-kit/kit/transport/grpc"

	endpoint "rocket-server/server/events/pkg/endpoint"
	grpc "rocket-server/server/events/pkg/transport/grpc"
	pb "rocket-server/server/events/pkg/transport/grpc/pb"
)

var grpcAddr = fs.String("grpc-addr", ":8082", "gRPC listen address")

func initGRPCHandler(endpoints endpoint.Endpoints, g *group.Group) {
	options := getGRPCOptions(logger, tracer)

	grpcServer := grpc.NewGRPCServer(endpoints, options)
	grpcListener, err := net.Listen("tcp", *grpcAddr)
	if err != nil {
		logger.Log("transport", "gRPC", "during", "Listen", "err", err)
	}
	g.Add(func() error {
		logger.Log("transport", "gRPC", "addr", *grpcAddr)
		baseServer := grpcgo.NewServer()
		pb.RegisterEventsServer(baseServer, grpcServer)
		return baseServer.Serve(grpcListener)
	}, func(error) {
		grpcListener.Close()
	})
}

func getGRPCOptions(logger log.Logger, tracer opentracinggo.Tracer) map[string][]grpckit.ServerOption {
	options := map[string][]grpckit.ServerOption{
		"AddEvent":  {grpckit.ServerErrorLogger(logger), grpckit.ServerBefore(opentracing.GRPCToContext(tracer, "AddEvent", logger))},
		"GetEvents": {grpckit.ServerErrorLogger(logger), grpckit.ServerBefore(opentracing.GRPCToContext(tracer, "GetEvents", logger))},
	}
	return options
}
