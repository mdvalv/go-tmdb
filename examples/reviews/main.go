package main

import (
	"github.com/mdvalv/go-tmdb"
	"github.com/mdvalv/go-tmdb/examples"
	"github.com/pkg/errors"
)

type example struct {
	client *tmdb.Client
}

func (e example) GetReview() {
	credit, _, err := e.client.Reviews.GetReview("5e7d573b6c849200126ac199")
	if err != nil {
		panic(errors.Wrap(err, "failed to get review"))
	}
	examples.PrettyPrint(*credit)
}

func main() {
	example := example{
		client: examples.GetClient(),
	}

	examples.RunExamples(
		example.GetReview,
	)
}
