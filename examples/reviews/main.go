// Reviews examples.
package main

import (
	"github.com/mdvalv/go-tmdb"
	"github.com/mdvalv/go-tmdb/examples"
)

type example struct {
	client *tmdb.Client
}

func (e example) GetReview() {
	credit, _, err := e.client.Reviews.GetReview("62e08236a44d0907dc16bce4")
	examples.PanicOnError(err)
	examples.PrettyPrint(*credit)
}

func main() {
	example := example{
		client: examples.GetClient(),
	}

	examples.RunExamples(
		example.GetReview, // 1
	)
}
