// Genres examples.
package main

import (
	"github.com/mdvalv/go-tmdb"
	"github.com/mdvalv/go-tmdb/examples"
)

type example struct {
	client *tmdb.Client
}

func (e example) GetMovieGenres() {
	genres, _, err := e.client.Genres.GetMovieGenres(nil)
	examples.PanicOnError(err)
	examples.PrettyPrint(genres)
}

func (e example) GetMovieGenresWithOptions() {
	options := tmdb.GenresOptions{
		Language: "pt-BR",
	}
	genres, _, err := e.client.Genres.GetMovieGenres(&options)
	examples.PanicOnError(err)
	examples.PrettyPrint(genres)
}

func (e example) GetTVGenres() {
	genres, _, err := e.client.Genres.GetTVGenres(nil)
	examples.PanicOnError(err)
	examples.PrettyPrint(genres)
}

func (e example) GetTVGenresWithOptions() {
	options := tmdb.GenresOptions{
		Language: "pt-BR",
	}
	genres, _, err := e.client.Genres.GetTVGenres(&options)
	examples.PanicOnError(err)
	examples.PrettyPrint(genres)
}

func main() {
	example := example{
		client: examples.GetClient(),
	}

	examples.RunExamples(
		example.GetMovieGenres,            // 1
		example.GetMovieGenresWithOptions, // 2
		example.GetTVGenres,               // 3
		example.GetTVGenresWithOptions,    // 4
	)
}
