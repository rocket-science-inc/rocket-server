package service

import (
	"flag"
	"net"

	endpoint1 "github.com/go-kit/kit/endpoint"
	log "github.com/go-kit/kit/log"
	group "github.com/oklog/oklog/pkg/group"
	grpc1 "google.golang.org/grpc"
	opentracinggo "github.com/opentracing/opentracing-go"

	endpoint "rocket-server/server/events/pkg/endpoint"
	grpc "rocket-server/server/events/pkg/grpc/handler"
	pb "rocket-server/server/events/pkg/grpc/pb"
	service "rocket-server/server/events/pkg/service"
)

var tracer opentracinggo.Tracer
var logger log.Logger

var fs = flag.NewFlagSet("events", flag.ExitOnError)
var grpcAddr = fs.String("grpc-addr", ":8082", "gRPC listen address")

func Run() {

}

func initGRPCHandler(endpoints endpoint.Endpoints, g *group.Group) {
	options := defaultGRPCOptions(logger, tracer)

	grpcServer := grpc.NewGRPCServer(endpoints, options)
	grpcListener, err := net.Listen("tcp", *grpcAddr)
	if err != nil {
		logger.Log("transport", "gRPC", "during", "Listen", "err", err)
	}
	g.Add(func() error {
		logger.Log("transport", "gRPC", "addr", *grpcAddr)
		baseServer := grpc1.NewServer()
		pb.RegisterEventsServer(baseServer, grpcServer)
		return baseServer.Serve(grpcListener)
	}, func(error) {
		grpcListener.Close()
	})

}
func getServiceMiddleware(logger log.Logger) (mw []service.Middleware) {
	mw = []service.Middleware{}

	return
}
func getEndpointMiddleware(logger log.Logger) (mw map[string][]endpoint1.Middleware) {
	mw = map[string][]endpoint1.Middleware{}

	return
}
