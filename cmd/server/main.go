package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	animepb "github.com/AkashGit21/Otaku_API/pb"
	"google.golang.org/grpc"
)

func main() {
	port := flag.Int("port", 0, "Server Port")
	serverType := flag.String("type", "grpc", "Type of Server - gRPC/REST")
	endPoint := flag.String("endpoint", "", "gRPC endpoint")
	flag.Parse()

	fmt.Println("Hello from Server...")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	animepb.RegisterAnimeServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
