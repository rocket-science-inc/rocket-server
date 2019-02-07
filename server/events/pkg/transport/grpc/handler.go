package grpc

import (
	grpc "github.com/go-kit/kit/transport/grpc"
	endpoint "rocket-server/server/events/pkg/endpoint"
	pb "rocket-server/server/events/pkg/transport/grpc/pb"
)

// NewGRPCServer makes a set of endpoints available as a gRPC AddServer
type grpcServer struct {
	getEvents grpc.Handler
	addEvent  grpc.Handler
}

func NewGRPCServer(endpoints endpoint.Endpoints, options map[string][]grpc.ServerOption) pb.EventsServer {
	return &grpcServer{
		addEvent:  makeAddEventHandler(endpoints, options["AddEvent"]),
		getEvents: makeGetEventsHandler(endpoints, options["GetEvents"]),
	}
}
