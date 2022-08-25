// Credits examples.
package main

import (
	"github.com/mdvalv/go-tmdb"
	"github.com/mdvalv/go-tmdb/examples"
	"github.com/pkg/errors"
)

type example struct {
	client *tmdb.Client
}

func (e example) GetCreditTV() {
	credit, _, err := e.client.Credits.GetCredit("525331fd19c295794001a5de")
	examples.PanicOnError(err)
	if credit.Media.GetMediaType() != "tv" {
		panic(errors.New("expected tv media type"))
	}
	examples.PrettyPrint(*credit)
}

func (e example) GetCreditMovie() {
	credit, _, err := e.client.Credits.GetCredit("52fe43f9c3a368484e0089e3")
	examples.PanicOnError(err)
	if credit.Media.GetMediaType() != "movie" {
		panic(errors.New("expected movie media type"))
	}
	examples.PrettyPrint(*credit)
}

func main() {
	example := example{
		client: examples.GetClient(),
	}

	examples.RunExamples(
		example.GetCreditTV,    // 1
		example.GetCreditMovie, // 2
	)
}
