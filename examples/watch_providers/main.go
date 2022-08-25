// Watch Providers examples.
package main

import (
	"github.com/mdvalv/go-tmdb"
	"github.com/mdvalv/go-tmdb/examples"
)

type example struct {
	client *tmdb.Client
}

func (e example) GetMovieProviders() {
	providers, _, err := e.client.WatchProviders.GetMovieProviders(nil)
	examples.PanicOnError(err)
	examples.PrettyPrint(providers)
}

func (e example) GetMovieProvidersWithOptions() {
	options := tmdb.ProvidersOptions{
		Language:    "pt-BR",
		WatchRegion: "BR",
	}
	providers, _, err := e.client.WatchProviders.GetMovieProviders(&options)
	examples.PanicOnError(err)
	examples.PrettyPrint(providers)
}

func (e example) GetTVProviders() {
	providers, _, err := e.client.WatchProviders.GetTVProviders(nil)
	examples.PanicOnError(err)
	examples.PrettyPrint(providers)
}

func (e example) GetTVProvidersWithOptions() {
	options := tmdb.ProvidersOptions{
		Language:    "pt-BR",
		WatchRegion: "BR",
	}
	providers, _, err := e.client.WatchProviders.GetTVProviders(&options)
	examples.PanicOnError(err)
	examples.PrettyPrint(providers)
}

func (e example) GetProviderRegions() {
	providerRegions, _, err := e.client.WatchProviders.GetProviderRegions(nil)
	examples.PanicOnError(err)
	examples.PrettyPrint(providerRegions)
}

func (e example) GetProviderRegionsWithOptions() {
	options := tmdb.ProviderRegionsOptions{
		Language: "pt-BR",
	}
	providerRegions, _, err := e.client.WatchProviders.GetProviderRegions(&options)
	examples.PanicOnError(err)
	examples.PrettyPrint(providerRegions)
}

func main() {
	example := example{
		client: examples.GetClient(),
	}

	examples.RunExamples(
		example.GetMovieProviders,             // 1
		example.GetMovieProvidersWithOptions,  // 2
		example.GetTVProviders,                // 3
		example.GetTVProvidersWithOptions,     // 4
		example.GetProviderRegions,            // 5
		example.GetProviderRegionsWithOptions, // 6
	)
}
