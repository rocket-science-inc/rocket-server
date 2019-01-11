package grpc

import (
	"context"
	"errors"
	grpc "github.com/go-kit/kit/transport/grpc"
	context1 "golang.org/x/net/context"
	endpoint "rocket-server/server/events/pkg/endpoint"
	pb "rocket-server/server/events/pkg/grpc/pb"
)

// makeGetEventsHandler creates the handler logic
func makeGetEventsHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.GetEventsEndpoint, decodeGetEventsRequest, encodeGetEventsResponse, options...)
}

// decodeGetEventsResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain sum request.
// TODO implement the decoder
func decodeGetEventsRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Events' Decoder is not impelemented")
}

// encodeGetEventsResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeGetEventsResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Events' Encoder is not impelemented")
}
func (g *grpcServer) GetEvents(ctx context1.Context, req *pb.GetEventsRequest) (*pb.GetEventsReply, error) {
	_, rep, err := g.getEvents.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.GetEventsReply), nil
}

// makeAddEventHandler creates the handler logic
func makeAddEventHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.AddEventEndpoint, decodeAddEventRequest, encodeAddEventResponse, options...)
}

// decodeAddEventResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain sum request.
// TODO implement the decoder
func decodeAddEventRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Events' Decoder is not impelemented")
}

// encodeAddEventResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeAddEventResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Events' Encoder is not impelemented")
}
func (g *grpcServer) AddEvent(ctx context1.Context, req *pb.AddEventRequest) (*pb.AddEventReply, error) {
	_, rep, err := g.addEvent.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.AddEventReply), nil
}
