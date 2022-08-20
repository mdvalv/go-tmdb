package main

import (
	"github.com/mdvalv/go-tmdb"
	"github.com/mdvalv/go-tmdb/examples"
)

type example struct {
	client *tmdb.Client
}

func (e example) GetKeyword() {
	keyword, _, err := e.client.Keywords.GetKeyword(4344)
	examples.PanicOnError(err)
	examples.PrettyPrint(*keyword)
}

func (e example) GetKeywordMovies() {
	keywordMovies, _, err := e.client.Keywords.GetKeywordMovies(4344, nil)
	examples.PanicOnError(err)
	examples.PrettyPrint(*keywordMovies)
}

func (e example) GetKeywordMoviesWithOptions() {
	page := 8
	options := tmdb.KeywordMoviesOptions{
		Page: &page,
	}
	keywordMovies, _, err := e.client.Keywords.GetKeywordMovies(4344, &options)
	examples.PanicOnError(err)
	examples.PrettyPrint(*keywordMovies)
}

func main() {
	example := example{
		client: examples.GetClient(),
	}

	examples.RunExamples(
		example.GetKeyword,                  // 1
		example.GetKeywordMovies,            // 2
		example.GetKeywordMoviesWithOptions, // 3
	)
}
