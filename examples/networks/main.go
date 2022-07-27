package main

import (
	"github.com/mdvalv/go-tmdb"
	"github.com/mdvalv/go-tmdb/examples"
	"github.com/pkg/errors"
)

type example struct {
	client *tmdb.Client
}

func (e example) GetNetwork() {
	network, _, err := e.client.Networks.GetNetwork(1)
	if err != nil {
		panic(errors.Wrap(err, "failed to get network"))
	}
	examples.PrettyPrint(*network)
}

func (e example) GetNetworkAlternativeNames() {
	names, _, err := e.client.Networks.GetAlternativeNames(1)
	if err != nil {
		panic(errors.Wrap(err, "failed to get network alternative names"))
	}
	examples.PrettyPrint(*names)
}

func (e example) GetNetworkImages() {
	images, _, err := e.client.Networks.GetImages(1)
	if err != nil {
		panic(errors.Wrap(err, "failed to get network images"))
	}
	examples.PrettyPrint(*images)
}

func main() {
	example := example{
		client: examples.GetClient(),
	}

	examples.RunExamples(
		example.GetNetwork,
		example.GetNetworkAlternativeNames,
		example.GetNetworkImages,
	)
}
