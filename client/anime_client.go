package client

import (
	"context"
	"log"
	"time"

	animepb "github.com/AkashGit21/Otaku_API/pb/github.com/AkashGit21/Otaku_API/proto/anime"
	"google.golang.org/grpc"
)

// AnimeClient is a client to call Anime service RPCs
type AnimeClient struct {
	service animepb.AnimeServiceClient
}

// NewAnimeClient returns a new Anime Client
func NewAnimeClient(cc grpc.ClientConnInterface) *AnimeClient {
	service := animepb.NewAnimeServiceClient(cc)
	return &AnimeClient{service}
}

func (animeClient *AnimeClient) ListAnimes(ctx context.Context, req *animepb.ListAnimesRequest, opts ...grpc.CallOption) (*animepb.ListAnimesResponse, error) {
	log.Printf("Received ListAnimes Request from Client: %v", req)

	// req1 := &animepb.ListAnimesRequest{
	// 	Page:   1,
	// 	Sort:   animepb.ListAnimesRequest_NAME,
	// 	Order:  animepb.ListAnimesRequest_ASCENDING,
	// 	Filter: []string{"genre"},
	// 	Search: "naruto",
	// }

	// set timeout
	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()

	res, err := animeClient.service.ListAnimes(ctx, req)
	if err != nil {
		log.Fatalf("Error while running Service\n")
		return nil, err
	}

	return res, nil
}
