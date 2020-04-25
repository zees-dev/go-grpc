package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"

	pb "github.com/zees-dev/go-grpc/protos"
	"google.golang.org/grpc"
)

func main() {
	typePtr := flag.String("type", "", "Denotes what type of grpc service to start: server, client")
	serverPortPtr := flag.Int("port", 50051, "Denotes gRPC server port")
	serverAddrPtr := flag.String("address", "localhost:50051", "Denotes gRPC server address in format: host:port")
	flag.Parse()

	switch *typePtr {
	case "server":
		runServer(uint16(*serverPortPtr))
	case "client":
		runClient(*serverAddrPtr, os.Args[3:]...)
	default:
		log.Printf("Unkown grpc service type: %q", *typePtr)
	}
}

// CLIENT CODE
func runClient(address string, messages ...string) {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	// Contact the server and print out its response.
	msg := strings.Join(messages, " ")
	if len(messages) == 0 {
		msg = "hello world"
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: msg})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("sent: %s", r.GetMessage())
}

// SERVER CODE
func runServer(port uint16) {
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

// server is used to implement helloworld.GreeterServer.
type server struct{}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("recv: %v", in.GetName())
	return &pb.HelloReply{Message: in.GetName()}, nil
}
