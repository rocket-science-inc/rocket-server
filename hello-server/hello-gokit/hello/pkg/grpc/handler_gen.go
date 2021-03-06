// THIS FILE IS AUTO GENERATED BY GK-CLI DO NOT EDIT!!
package grpc

import (
	grpc "github.com/go-kit/kit/transport/grpc"
	endpoint "rocket-server/hello-server/hello-gokit/hello/pkg/endpoint"
	pb "rocket-server/hello-server/hello-gokit/hello/pkg/grpc/pb"
)

// NewGRPCServer makes a set of endpoints available as a gRPC AddServer
type grpcServer struct {
	sayHello grpc.Handler
}

func NewGRPCServer(endpoints endpoint.Endpoints, options map[string][]grpc.ServerOption) pb.HelloServer {
	return &grpcServer{sayHello: makeSayHelloHandler(endpoints, options["SayHello"])}
}
