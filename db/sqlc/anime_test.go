package db

import (
	"context"
	"testing"

	animepb "github.com/AkashGit21/Otaku_API/pb/github.com/AkashGit21/Otaku_API/proto/anime"
	"github.com/AkashGit21/Otaku_API/util"
	"github.com/stretchr/testify/require"
)

func createRadomAnime(t *testing.T) (*animepb.Anime, error) {
	arg := &animepb.Anime{
		Name:          util.RandomName(),
		Description:   util.RandomDescriptiion(),
		Status:        "COMPLETED",
		NumOfEpisodes: util.RandomEpisodes(),
		Cast:          util.RandomCast(),
		Genre:         util.RandomGenre(),
	}

	anime, err := testQueries.CreateAnime(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, anime)

	require.Equal(t, arg.Name, anime.Name)
	require.Equal(t, arg.Description, anime.Description)
	require.Equal(t, arg.Status, anime.Status)
	require.Equal(t, arg.NumOfEpisodes, anime.NumOfEpisodes)
	require.Equal(t, arg.Cast, anime.Cast)
	require.Equal(t, arg.Genre, anime.Genre)

	return anime, err
}

func TestCreateAnime(t *testing.T) {
	createRadomAnime(t)
}

func TestListAnimes(t *testing.T) {

	animes, err := testQueries.ListAnimes(context.Background(), 1)

	require.NoError(t, err)
	require.NotEmpty(t, animes)

	for _, anime := range animes {
		require.NotEmpty(t, anime)
	}
}
