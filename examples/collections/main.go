package main

import (
	"github.com/mdvalv/go-tmdb"
	"github.com/mdvalv/go-tmdb/examples"
)

type example struct {
	client *tmdb.Client
}

func (e example) GetCollection() {
	collection, _, err := e.client.Collections.GetCollection(131635, nil)
	examples.PanicOnError(err)
	examples.PrettyPrint(*collection)
}

func (e example) GetCollectionWithOptions() {
	options := tmdb.CollectionsOptions{
		Language: "pt-BR",
	}
	collection, _, err := e.client.Collections.GetCollection(131635, &options)
	examples.PanicOnError(err)
	examples.PrettyPrint(*collection)
}

func (e example) GetCollectionImages() {
	images, _, err := e.client.Collections.GetImages(131635, nil)
	examples.PanicOnError(err)
	examples.PrettyPrint(*images)
}

func (e example) GetCollectionTranslations() {
	images, _, err := e.client.Collections.GetTranslations(131635, nil)
	examples.PanicOnError(err)
	examples.PrettyPrint(*images)
}

func main() {
	example := example{
		client: examples.GetClient(),
	}

	examples.RunExamples(
		example.GetCollection,             // 1
		example.GetCollectionWithOptions,  // 2
		example.GetCollectionImages,       // 3
		example.GetCollectionTranslations, // 4
	)
}
