package main

import (
	"context"
	"log"
	"net"

	pb "github.com/wonderstone/grpc_demo/example"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedGreeterServer
	pb.UnimplementedCalculatorServer
}

// SayHello(context.Context, *HelloRequest) (*HelloResponse, error)

func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	log.Printf("Received: %v", req.GetName())
	return &pb.HelloResponse{Message: "Hello " + req.GetName()}, nil
}
func (s *server) Add(ctx context.Context, req *pb.AddRequest) (*pb.AddResponse, error) {
	log.Printf("Received: %v + %v", req.GetA(), req.GetB())
	return &pb.AddResponse{Result: req.GetA() + req.GetB()}, nil
}

func (s *server) Subtract(ctx context.Context, req *pb.SubtractRequest) (*pb.SubtractResponse, error) {
	log.Printf("Received: %v - %v", req.GetA(), req.GetB())
	return &pb.SubtractResponse{Result: req.GetA() - req.GetB()}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()

	pb.RegisterGreeterServer(s, &server{})
	pb.RegisterCalculatorServer(s, &server{})
	log.Printf("Server started on port 8080")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
