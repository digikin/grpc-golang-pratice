package main

import (
	"context"
	"fmt"
	"log"

	"github.com/grpc-go-course/calculator/calculatorpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Client starting")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer cc.Close()

	c := calcpb.NewGreetServiceClient(cc)
	doUnary(c)
}

func doUnary(c calcpb.GreetServiceClient) {
	fmt.Println("Completing Unary RPC connection...")
	req := &calcpb.GreetRequest{
		GreetingFirst:  10,
		GreetingSecond: 3,
	}
	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("error while connecting: %v", err)
	}
	log.Printf("Response from Greet: %v", res.Result)
}
