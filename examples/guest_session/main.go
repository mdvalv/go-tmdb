package main

import (
	"github.com/mdvalv/go-tmdb"
	"github.com/mdvalv/go-tmdb/examples"
	"github.com/pkg/errors"
)

type example struct {
	client *tmdb.Client
}

const (
	sessionId = "gest_session_id"
)

func (e example) GetRatedMovies() {
	account, _, err := e.client.GuestSession.GetRatedMovies(sessionId, nil)
	if err != nil {
		panic(errors.Wrap(err, "failed to get rated movies"))
	}
	examples.PrettyPrint(*account)
}

func (e example) GetRatedTVShows() {
	tvShows, _, err := e.client.GuestSession.GetRatedTVShows(sessionId, nil)
	if err != nil {
		panic(errors.Wrap(err, "failed to get rated tv shows"))
	}
	examples.PrettyPrint(*tvShows)
}

func (e example) GetRatedTVEpisodes() {
	episodes, _, err := e.client.GuestSession.GetRatedTVEpisodes(sessionId, nil)
	if err != nil {
		panic(errors.Wrap(err, "failed to get rated tv episodes"))
	}
	examples.PrettyPrint(*episodes)
}

func main() {
	example := example{
		client: examples.GetClient(),
	}

	examples.RunExamples(
		example.GetRatedMovies,
		example.GetRatedTVShows,
		example.GetRatedTVEpisodes,
	)
}
