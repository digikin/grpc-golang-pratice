package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/grpc-go-course/calculator/calculatorpb"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) Greet(ctx context.Context, req *calcpb.GreetRequest) (*calcpb.GreetResponse, error) {
	fmt.Printf("Greet function invoked %v\n", req)
	first := req.GetGreetingFirst()
	second := req.GetGreetingSecond()
	result := first + second
	res := &calcpb.GreetResponse{
		Result:               result,
		XXX_NoUnkeyedLiteral: struct{}{},
		XXX_unrecognized:     nil,
		XXX_sizecache:        0,
	}
	return res, nil
}

func main() {
	fmt.Println("Server Started!")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	calcpb.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
