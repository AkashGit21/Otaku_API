package service

import (
	"context"
	"log"
	"strconv"

	db "github.com/AkashGit21/Otaku_API/db/sqlc"
	animepb "github.com/AkashGit21/Otaku_API/pb/github.com/AkashGit21/Otaku_API/proto/anime"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

var dbQueries *db.Queries

//
type AnimeServer struct {
	animepb.UnimplementedAnimeServiceServer
}

//
func NewAnimeServer() *AnimeServer {
	return &AnimeServer{}
}

// ListAnimes is a unary RPC to list all the animes
func (server *AnimeServer) ListAnimes(ctx context.Context, req *animepb.ListAnimesRequest) (*animepb.ListAnimesResponse, error) {
	log.Printf("Received ListAnimes Request from Server: %v\n", req)

	page := req.GetPage()
	sort := req.GetSort()
	order := req.GetOrder()
	filters := req.GetFilter()
	searchQuery := req.GetSearch()

	log.Printf("Query Parameters for ListAnimes are:\n \tPage: %d\n\tSort: %v\n\tOrder: %v\n\tFilters: %v\n\tSearch: %v\n",
		page, sort, order, filters, searchQuery)

	dbQueries, err := db.MakeConnection(dbQueries)
	if err != nil {
		log.Fatalf("Error while making DB connection\n")
		return nil, err
	}

	if ctx.Err() == context.DeadlineExceeded {
		log.Printf("Deadline is Exceeded")
		return nil, status.Error(codes.DeadlineExceeded, "Request deadline Exceeded")
	}

	if ctx.Err() == context.Canceled {
		log.Printf("Request is Canceled")
		return nil, status.Error(codes.Canceled, "Request is Canceled")
	}

	res, err := dbQueries.ListAnimes(ctx, 1)
	if err != nil {
		log.Fatalf("Error while asking response from DB for ListAnimes: %v\n", err)
		return nil, err
	}

	out := &animepb.ListAnimesResponse{
		Animes: res,
	}

	return out, nil
}

// CreateAnime is a unary RPC to create an Anime
func (server *AnimeServer) CreateAnime(ctx context.Context, req *animepb.CreateAnimeRequest) (*animepb.CreateAnimeResponse, error) {

	log.Printf("Received CreateAnime Request from Server: %v", req)

	dbQueries, err := db.MakeConnection(dbQueries)
	if err != nil {
		log.Fatalf("Error while making DB connection\n")
		return nil, err
	}

	if ctx.Err() == context.DeadlineExceeded {
		log.Printf("Deadline is Exceeded")
		return nil, status.Error(codes.DeadlineExceeded, "Request deadline Exceeded")
	}

	if ctx.Err() == context.Canceled {
		log.Printf("Request is Canceled")
		return nil, status.Error(codes.Canceled, "Request is Canceled")
	}

	data := req.Anime
	res, err := dbQueries.CreateAnime(ctx, data)

	out := &animepb.CreateAnimeResponse{
		Id: res.Id,
	}

	return out, nil
}

// GetAnime returns the Anime with given ID
func (server *AnimeServer) GetAnime(ctx context.Context, req *animepb.GetAnimeRequest) (*animepb.GetAnimeResponse, error) {

	log.Printf("Received GetAnime Request from Server: %v\n", req)

	dbQueries, err := db.MakeConnection(dbQueries)
	if err != nil {
		log.Fatalf("Error while making DB connection\n")
		return nil, err
	}

	if ctx.Err() == context.DeadlineExceeded {
		log.Printf("Deadline is Exceeded")
		return nil, status.Error(codes.DeadlineExceeded, "Request deadline Exceeded")
	}

	if ctx.Err() == context.Canceled {
		log.Printf("Request is Canceled")
		return nil, status.Error(codes.Canceled, "Request is Canceled")
	}

	id, err := strconv.ParseInt(req.GetId(), 10, 64)
	if err != nil {
		log.Fatalf("Id is not in proper Format\n")
		return nil, err
	}

	res, err := dbQueries.GetAnime(ctx, id)
	if err != nil {
		log.Fatalf("Error while fetching Anime with given ID:%v", id)
		return nil, err
	}

	out := &animepb.GetAnimeResponse{
		Anime: res,
	}
	return out, nil
}

// UpdateAnime returns the Anime with given ID
func (server *AnimeServer) UpdateAnime(ctx context.Context, req *animepb.UpdateAnimeRequest) (*animepb.Anime, error) {

	log.Printf("Received UpdateAnime Request from Server: %v\n", req)

	dbQueries, err := db.MakeConnection(dbQueries)
	if err != nil {
		log.Fatalf("Error while making DB connection\n")
		return nil, err
	}

	if ctx.Err() == context.DeadlineExceeded {
		log.Printf("Deadline is Exceeded")
		return nil, status.Error(codes.DeadlineExceeded, "Request deadline Exceeded")
	}

	if ctx.Err() == context.Canceled {
		log.Printf("Request is Canceled")
		return nil, status.Error(codes.Canceled, "Request is Canceled")
	}

	err = dbQueries.UpdateAnime(ctx, req.Anime)
	if err != nil {
		log.Fatalf("Error while updating the Anime")
		return nil, err
	}

	id, err := strconv.ParseInt(req.GetId(), 10, 64)
	if err != nil {
		log.Fatalf("Id is not in proper Format\n")
		return nil, err
	}

	res, err := dbQueries.GetAnime(ctx, id)
	if err != nil {
		log.Fatalf("Cannot fetch updated Anime\n")
		return nil, err
	}

	return res, nil
}

// DeleteAnime deletes the Anime with given ID
func (server *AnimeServer) DeleteAnime(ctx context.Context, req *animepb.DeleteAnimeRequest) (*empty.Empty, error) {

	log.Printf("Received DeleteAnime Request from Server: %v\n", req)

	dbQueries, err := db.MakeConnection(dbQueries)
	if err != nil {
		log.Fatalf("Error while making DB connection\n")
		return nil, err
	}

	if ctx.Err() == context.DeadlineExceeded {
		log.Printf("Deadline is Exceeded")
		return nil, status.Error(codes.DeadlineExceeded, "Request deadline Exceeded")
	}

	if ctx.Err() == context.Canceled {
		log.Printf("Request is Canceled")
		return nil, status.Error(codes.Canceled, "Request is Canceled")
	}

	id, err := strconv.ParseInt(req.GetId(), 10, 64)
	if err != nil {
		log.Fatalf("Id is not in proper Format\n")
		return nil, err
	}

	err = dbQueries.DeleteAnime(ctx, id)
	if err != nil {
		log.Fatalf("Error while deleting the Anime\n")
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
