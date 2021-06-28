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
		Type:          util.RandomString(8),
		Summary:       util.RandomSummary(),
		NumOfEpisodes: util.RandomEpisodes(),
		OtherNames:    util.RandomNames(),
		Status:        "COMPLETED",
		Genre:         util.RandomGenre(),
		Released:      util.RandomInt(1990, 2022),
	}

	anime, err := testQueries.CreateAnime(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, anime)

	require.Equal(t, arg.Name, anime.Name)
	require.Equal(t, arg.Type, anime.Type)
	require.Equal(t, arg.Summary, anime.Summary)
	require.Equal(t, arg.NumOfEpisodes, anime.NumOfEpisodes)
	require.Equal(t, arg.OtherNames, anime.OtherNames)
	require.Equal(t, arg.Status, anime.Status)
	require.Equal(t, arg.Genre, anime.Genre)
	require.Equal(t, arg.Released, anime.Released)

	return anime, err
}

func TestCreateAnime(t *testing.T) {
	createRadomAnime(t)
}

func TestListAnimes(t *testing.T) {

	arg := ListAnimesParams{
		Column1: "a",
		Column2: 1,
		// Column3: "FCI",
	}

	animes, err := testQueries.ListAnimes(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, animes)

	for _, anime := range animes {
		require.NotEmpty(t, anime)
	}
}
