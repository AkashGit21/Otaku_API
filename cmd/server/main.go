package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"

	animepb "github.com/AkashGit21/Otaku_API/pb/github.com/AkashGit21/Otaku_API/proto/anime"
	"github.com/AkashGit21/Otaku_API/service"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

func runGRPCServer(animeServer animepb.AnimeServiceServer, listener net.Listener) error {

	grpcServer := grpc.NewServer()
	animepb.RegisterAnimeServiceServer(grpcServer, animeServer)

	err := grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("Failed to Serve\n ")
		return err
	}

	return nil
}

func runRESTServer(animeServer animepb.AnimeServiceServer, listener net.Listener, grpcEndpoint string) error {
	mux := runtime.NewServeMux()
	// dialOptions := []grpc.DialOption{grpc.WithInsecure()}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := animepb.RegisterAnimeServiceHandlerServer(ctx, mux, animeServer)
	if err != nil {
		return err
	}

	// err := animepb.RegisterAnimeServiceHandlerFromEndpoint(ctx, mux, grpcEndpoint, dialOptions)
	// if err != nil {
	// 	log.Fatalf("Error while Registering Server from Endpoint")
	// 	return err
	// }

	log.Printf("Start REST server at %s", listener.Addr().String())

	return http.Serve(listener, mux)
}

func main() {
	port := flag.Int("port", 8080, "Server Port")
	serverType := flag.String("type", "grpc", "Type of Server - grpc/rest")
	endPoint := flag.String("endpoint", "", "gRPC endpoint")
	flag.Parse()
	log.Printf("Server starting with Port:%d and Endpoint:%s", *port, *endPoint)

	animeServer := service.NewAnimeServer()

	address := fmt.Sprintf("localhost:%d", *port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	if *serverType == "rest" {
		runRESTServer(animeServer, listener, *endPoint)
	} else if *serverType == "grpc" {
		runGRPCServer(animeServer, listener)
	}
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}
}
