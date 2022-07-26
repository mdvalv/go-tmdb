// Guest Session examples.
package main

import (
	"github.com/mdvalv/go-tmdb"
	"github.com/mdvalv/go-tmdb/examples"
)

type example struct {
	client *tmdb.Client
}

const (
	sessionID = "guest_session_id"
)

func (e example) GetRatedMovies() {
	account, _, err := e.client.GuestSession.GetRatedMovies(sessionID, nil)
	examples.PanicOnError(err)
	examples.PrettyPrint(*account)
}

func (e example) GetRatedTVShows() {
	tvShows, _, err := e.client.GuestSession.GetRatedTVShows(sessionID, nil)
	examples.PanicOnError(err)
	examples.PrettyPrint(*tvShows)
}

func (e example) GetRatedTVEpisodes() {
	episodes, _, err := e.client.GuestSession.GetRatedTVEpisodes(sessionID, nil)
	examples.PanicOnError(err)
	examples.PrettyPrint(*episodes)
}

func main() {
	example := example{
		client: examples.GetClient(),
	}

	examples.RunExamples(
		example.GetRatedMovies,     // 1
		example.GetRatedTVShows,    // 2
		example.GetRatedTVEpisodes, // 3
	)
}
