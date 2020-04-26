package client

import (
	"context"
	"log"
	"strings"
	"time"

	pb "github.com/zees-dev/go-grpc/internal/protos"
	"google.golang.org/grpc"
)

// RunClient will run a simple gRPC command as a gRPC client
func RunClient(address string, messages ...string) {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	msg := strings.Join(messages, " ")
	if len(messages) == 0 {
		msg = "hello world"
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: msg})
	if err != nil {
		log.Fatalf("could not send message: %v", err)
	}

	log.Printf("sent: %s", r.GetMessage())
}
