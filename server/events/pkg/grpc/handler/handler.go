package handler

import (
	grpc "github.com/go-kit/kit/transport/grpc"
	context "golang.org/x/net/context"

	pb "rocket-server/server/events/pkg/grpc/pb"
	endpoint "rocket-server/server/events/pkg/endpoint"
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

func (g *grpcServer) AddEvent(ctx context.Context, req *pb.AddEventRequest) (*pb.AddEventReply, error) {
	_, rep, err := g.addEvent.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.AddEventReply), nil
}

func (g *grpcServer) GetEvents(ctx context.Context, req *pb.GetEventsRequest) (*pb.GetEventsReply, error) {
	_, rep, err := g.getEvents.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.GetEventsReply), nil
}
