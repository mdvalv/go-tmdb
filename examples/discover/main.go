// Discover examples.
package main

import (
	"github.com/mdvalv/go-tmdb"
	"github.com/mdvalv/go-tmdb/examples"
)

type example struct {
	client *tmdb.Client
}

func (e example) DiscoverMovies() {
	discover, _, err := e.client.Discover.DiscoverMovies(nil)
	examples.PanicOnError(err)
	examples.PrettyPrint(*discover)
}

func (e example) DiscoverMoviesWithOptions() {
	voteCountGte := 800
	voteCountLte := 900
	opt := tmdb.DiscoverMoviesOptions{
		VoteCountGte:         &voteCountGte,
		VoteCountLte:         &voteCountLte,
		WithOriginalLanguage: "pt",
	}
	discover, _, err := e.client.Discover.DiscoverMovies(&opt)
	examples.PanicOnError(err)
	examples.PrettyPrint(*discover)
}

func (e example) DiscoverTvShows() {
	discover, _, err := e.client.Discover.DiscoverTVShows(nil)
	examples.PanicOnError(err)
	examples.PrettyPrint(*discover)
}

func (e example) DiscoverTvShowsWithOptions() {
	year := 2016
	voteCount := 600
	opt := tmdb.DiscoverTVShowsOptions{
		FirstAirDateYear:     &year,
		WithGenres:           "10759,10765,37",
		WithOriginalLanguage: "en",
		VoteCountGte:         &voteCount,
		WithStatus:           "3",
	}
	discover, _, err := e.client.Discover.DiscoverTVShows(&opt)
	examples.PanicOnError(err)
	examples.PrettyPrint(*discover)
}

func main() {
	example := example{
		client: examples.GetClient(),
	}

	examples.RunExamples(
		example.DiscoverMovies,             // 1
		example.DiscoverMoviesWithOptions,  // 2
		example.DiscoverTvShows,            // 3
		example.DiscoverTvShowsWithOptions, // 4
	)
}
