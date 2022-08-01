package main

import (
	"context"
	"log"

	pb "github.com/junkeWu/go-grpc-example/proto"
	"google.golang.org/grpc"
)

// PORT server port
const PORT = "9001"

func main() {
	dial, err := grpc.Dial(":"+PORT, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("grpc.Dial err: %v", err)
	}
	defer dial.Close()

	client := pb.NewSearchServiceClient(dial)
	resp, err := client.Search(context.Background(), &pb.SearchRequest{
		Request: "gRPC",
	})
	if err != nil {
		log.Fatalf("client.Serarch err: %v", err)
	}
	log.Printf("resp: %s", resp.GetResponse())
}
