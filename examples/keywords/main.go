package main

import (
	"github.com/mdvalv/go-tmdb"
	"github.com/mdvalv/go-tmdb/examples"
	"github.com/pkg/errors"
)

type example struct {
	client *tmdb.Client
}

func (e example) GetKeyword() {
	keyword, _, err := e.client.Keywords.GetKeyword(4344)
	if err != nil {
		panic(errors.Wrap(err, "failed to get keyword"))
	}
	examples.PrettyPrint(*keyword)
}

func (e example) GetKeywordMovies() {
	keywordMovies, _, err := e.client.Keywords.GetKeywordMovies(4344, nil)
	if err != nil {
		panic(errors.Wrap(err, "failed to get keyword movies"))
	}
	examples.PrettyPrint(*keywordMovies)
}

func (e example) GetKeywordMoviesWithOptions() {
	page := 8
	options := tmdb.KeywordMoviesOptions{
		Page: &page,
	}
	keywordMovies, _, err := e.client.Keywords.GetKeywordMovies(4344, &options)
	if err != nil {
		panic(errors.Wrap(err, "failed get keyword movies with options"))
	}
	examples.PrettyPrint(*keywordMovies)
}

func main() {
	example := example{
		client: examples.GetClient(),
	}

	examples.RunExamples(
		example.GetKeyword,
		example.GetKeywordMovies,
		example.GetKeywordMoviesWithOptions,
	)
}
