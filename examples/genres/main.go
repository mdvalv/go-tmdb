package main

import (
	"github.com/mdvalv/go-tmdb"
	"github.com/mdvalv/go-tmdb/examples"
	"github.com/pkg/errors"
)

type example struct {
	client *tmdb.Client
}

func (e example) GetMovieGenres() {
	genres, _, err := e.client.Genres.GetMovieGenres(nil)
	if err != nil {
		panic(errors.Wrap(err, "failed to get movie genres"))
	}
	examples.PrettyPrint(genres)
}

func (e example) GetMovieGenresWithOptions() {
	options := tmdb.GenresOptions{
		Language: "pt-BR",
	}
	genres, _, err := e.client.Genres.GetMovieGenres(&options)
	if err != nil {
		panic(errors.Wrap(err, "failed get movie genres with options"))
	}
	examples.PrettyPrint(genres)
}

func (e example) GetTVGenres() {
	genres, _, err := e.client.Genres.GetTVGenres(nil)
	if err != nil {
		panic(errors.Wrap(err, "failed get TV genres"))
	}
	examples.PrettyPrint(genres)
}

func (e example) GetTVGenresWithOptions() {
	options := tmdb.GenresOptions{
		Language: "pt-BR",
	}
	genres, _, err := e.client.Genres.GetTVGenres(&options)
	if err != nil {
		panic(errors.Wrap(err, "failed get TV genres with options"))
	}
	examples.PrettyPrint(genres)
}

func main() {
	example := example{
		client: examples.GetClient(),
	}

	examples.RunExamples(
		example.GetMovieGenres,
		example.GetMovieGenresWithOptions,
		example.GetTVGenres,
		example.GetTVGenresWithOptions,
	)
}
