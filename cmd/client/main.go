package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/AkashGit21/Otaku_API/client"
	animepb "github.com/AkashGit21/Otaku_API/pb/github.com/AkashGit21/Otaku_API/proto/anime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func main() {
	serverAddress := flag.String("address", "127.0.0.1:8080", "the server address")

	flag.Parse()
	fmt.Printf("Dial server:%s", *serverAddress)

	conn, err := grpc.Dial(*serverAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Cannot dial Server: %v", err)
		return
	}
	defer conn.Close()

	animeClient := client.NewAnimeClient(conn)
	// animeClient := animepb.NewAnimeServiceClient(conn)
	testListAnimes(animeClient)
}

func testListAnimes(client animepb.AnimeServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	req := &animepb.ListAnimesRequest{
		Page:   1,
		Sort:   animepb.ListAnimesRequest_NAME,
		Order:  animepb.ListAnimesRequest_ASCENDING,
		Filter: []string{"genre"},
		Search: "naruto",
	}
	res, err := client.ListAnimes(ctx, req)
	if err != nil {
		st, ok := status.FromError(err)
		if ok && st.Code() == codes.DeadlineExceeded {
			log.Printf("Deadline Exceeded")
		} else {
			log.Fatalf("Cannot fetch list of Animes: %v", err)
		}
		return
	}

	log.Println(res.GetAnimes())
}
