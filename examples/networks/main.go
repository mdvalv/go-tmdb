// Networks examples.
package main

import (
	"github.com/mdvalv/go-tmdb"
	"github.com/mdvalv/go-tmdb/examples"
)

type example struct {
	client *tmdb.Client
}

func (e example) GetNetwork() {
	network, _, err := e.client.Networks.GetNetwork(1)
	examples.PanicOnError(err)
	examples.PrettyPrint(*network)
}

func (e example) GetNetworkAlternativeNames() {
	names, _, err := e.client.Networks.GetAlternativeNames(1)
	examples.PanicOnError(err)
	examples.PrettyPrint(*names)
}

func (e example) GetNetworkImages() {
	images, _, err := e.client.Networks.GetImages(1)
	examples.PanicOnError(err)
	examples.PrettyPrint(*images)
}

func main() {
	example := example{
		client: examples.GetClient(),
	}

	examples.RunExamples(
		example.GetNetwork,                 // 1
		example.GetNetworkAlternativeNames, // 2
		example.GetNetworkImages,           // 3
	)
}
