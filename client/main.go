package main

import (
	"log"

	pb "grpc-demo/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":8080"
)

func main() {
	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect the server %v", err)
	}
	defer conn.Close()

	client := pb.NewGreetServiceClient(conn)

	names := &pb.NamesList{
		Names: []string{"Arsene", "Alice", "Charlie"},
	}

	// callSayHello(client)
	// callSayHelloServerStreaming(client, names)
	// callSayHelloClientStreaming(client, names)
	callSayHelloBidirectionalStreaming(client, names)
}
