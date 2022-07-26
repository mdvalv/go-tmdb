package main

import (
	"github.com/mdvalv/go-tmdb"
	"github.com/mdvalv/go-tmdb/examples"
	"github.com/pkg/errors"
)

type example struct {
	client *tmdb.Client
}

func (e example) GetCredit() {
	credit, _, err := e.client.Credits.GetCredit("525331fd19c295794001a5de")
	if err != nil {
		panic(errors.Wrap(err, "failed to get credits"))
	}
	examples.PrettyPrint(*credit)
}

func main() {
	example := example{
		client: examples.GetClient(),
	}

	examples.RunExamples(
		example.GetCredit,
	)
}
