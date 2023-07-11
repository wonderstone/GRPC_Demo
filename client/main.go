package main

import (
	"context"
	"log"

	pb "github.com/wonderstone/grpc_demo/example"
	"google.golang.org/grpc/credentials/insecure"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewGreeterClient(conn)

	resp, err := c.SayHello(context.Background(), &pb.HelloRequest{Name: "World"})
	if err != nil {
		log.Fatalf("Failed to call SayHello: %v", err)
	}

	log.Printf("Message: %s", resp.GetMessage())

	calculator := pb.NewCalculatorClient(conn)
	addResp, err := calculator.Add(context.Background(), &pb.AddRequest{A: 1, B: 2})
	if err != nil {
		log.Fatalf("Failed to call Add: %v", err)
	}
	log.Printf("Result: %d", addResp.GetResult())

	subResp, err := calculator.Subtract(context.Background(), &pb.SubtractRequest{A: 5, B: 3})
	if err != nil {
		log.Fatalf("Failed to call Subtract: %v", err)
	}
	log.Printf("Result: %d", subResp.GetResult())

}
