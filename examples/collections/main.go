package main

import (
	"github.com/mdvalv/go-tmdb"
	"github.com/mdvalv/go-tmdb/examples"
	"github.com/pkg/errors"
)

type example struct {
	client *tmdb.Client
}

func (e example) GetCollection() {
	collection, _, err := e.client.Collections.GetCollection(131635, nil)
	if err != nil {
		panic(errors.Wrap(err, "failed to get collection"))
	}
	examples.PrettyPrint(*collection)
}

func (e example) GetCollectionWithOptions() {
	options := tmdb.CollectionsOptions{
		Language: "pt-BR",
	}
	collection, _, err := e.client.Collections.GetCollection(131635, &options)
	if err != nil {
		panic(errors.Wrap(err, "failed to get collection with options"))
	}
	examples.PrettyPrint(*collection)
}

func (e example) GetCollectionImages() {
	images, _, err := e.client.Collections.GetImages(131635, nil)
	if err != nil {
		panic(errors.Wrap(err, "failed to get collection images"))
	}
	examples.PrettyPrint(*images)
}

func (e example) GetCollectionTranslations() {
	images, _, err := e.client.Collections.GetTranslations(131635, nil)
	if err != nil {
		panic(errors.Wrap(err, "failed to get collection translations"))
	}
	examples.PrettyPrint(*images)
}

func main() {
	example := example{
		client: examples.GetClient(),
	}

	examples.RunExamples(
		example.GetCollection,
		example.GetCollectionWithOptions,
		example.GetCollectionImages,
		example.GetCollectionTranslations,
	)
}
