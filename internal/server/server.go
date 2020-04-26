package server

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/zees-dev/go-grpc/internal/protos"
	"google.golang.org/grpc"
)

// server is used to implement helloworld.GreeterServer.
type server struct{}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("recv: %v", in.GetName())
	return &pb.HelloReply{Message: in.GetName()}, nil
}

func RunServer(port uint16) {
	portStr := fmt.Sprintf(":%v", port)
	lis, err := net.Listen("tcp", portStr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterGreeterServer(grpcServer, &server{})
	// determine whether to use TLS

	log.Printf("Running gRPC server on port %v...", portStr)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
