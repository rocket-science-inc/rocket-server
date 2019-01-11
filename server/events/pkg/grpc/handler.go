package grpc

import (
	"context"
	"errors"
	grpc "github.com/go-kit/kit/transport/grpc"
	context1 "golang.org/x/net/context"
	endpoint "rocket-server/server/events/pkg/endpoint"
	pb "rocket-server/server/events/pkg/grpc/pb"
)

// makeGetHandler creates the handler logic
func makeGetHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.GetEndpoint, decodeGetRequest, encodeGetResponse, options...)
}

// decodeGetResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain sum request.
// TODO implement the decoder
func decodeGetRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Events' Decoder is not impelemented")
}

// encodeGetResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeGetResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Events' Encoder is not impelemented")
}
func (g *grpcServer) Get(ctx context1.Context, req *pb.GetRequest) (*pb.GetReply, error) {
	_, rep, err := g.get.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.GetReply), nil
}

// makeAddHandler creates the handler logic
func makeAddHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.AddEndpoint, decodeAddRequest, encodeAddResponse, options...)
}

// decodeAddResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain sum request.
// TODO implement the decoder
func decodeAddRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Events' Decoder is not impelemented")
}

// encodeAddResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeAddResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Events' Encoder is not impelemented")
}
func (g *grpcServer) Add(ctx context1.Context, req *pb.AddRequest) (*pb.AddReply, error) {
	_, rep, err := g.add.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.AddReply), nil
}
