package handler

import (
	grpc "github.com/go-kit/kit/transport/grpc"
	
	pb "rocket-server/server/api/pkg/grpc/pb"
	endpoint "rocket-server/server/api/pkg/endpoint"
)

// NewGRPCServer makes a set of endpoints available as a gRPC AddServer
type grpcServer struct {

}

func NewGRPCServer(endpoints endpoint.Endpoints, options map[string][]grpc.ServerOption) pb.ApiServer {
	return &grpcServer {

	}
}
