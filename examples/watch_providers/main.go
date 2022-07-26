package main

import (
	"github.com/mdvalv/go-tmdb"
	"github.com/mdvalv/go-tmdb/examples"
	"github.com/pkg/errors"
)

type example struct {
	client *tmdb.Client
}

func (e example) GetMovieProviders() {
	providers, _, err := e.client.WatchProviders.GetMovieProviders(nil)
	if err != nil {
		panic(errors.Wrap(err, "failed to get movie providers"))
	}
	examples.PrettyPrint(providers)
}

func (e example) GetMovieProvidersWithOptions() {
	options := tmdb.ProvidersOptions{
		Language:    "pt-BR",
		WatchRegion: "BR",
	}
	providers, _, err := e.client.WatchProviders.GetMovieProviders(&options)
	if err != nil {
		panic(errors.Wrap(err, "failed to get movie providers with options"))
	}
	examples.PrettyPrint(providers)
}

func (e example) GetTVProviders() {
	providers, _, err := e.client.WatchProviders.GetTVProviders(nil)
	if err != nil {
		panic(errors.Wrap(err, "failed to get TV providers"))
	}
	examples.PrettyPrint(providers)
}

func (e example) GetTVProvidersWithOptions() {
	options := tmdb.ProvidersOptions{
		Language:    "pt-BR",
		WatchRegion: "BR",
	}
	providers, _, err := e.client.WatchProviders.GetTVProviders(&options)
	if err != nil {
		panic(errors.Wrap(err, "failed to get TV providers with options"))
	}
	examples.PrettyPrint(providers)
}

func (e example) GetProviderRegions() {
	providerRegions, _, err := e.client.WatchProviders.GetProviderRegions(nil)
	if err != nil {
		panic(errors.Wrap(err, "failed to get provider regions"))
	}
	examples.PrettyPrint(providerRegions)
}

func (e example) GetProviderRegionsWithOptions() {
	options := tmdb.ProviderRegionsOptions{
		Language: "pt-BR",
	}
	providerRegions, _, err := e.client.WatchProviders.GetProviderRegions(&options)
	if err != nil {
		panic(errors.Wrap(err, "failed to get provider regions with options"))
	}
	examples.PrettyPrint(providerRegions)
}

func main() {
	example := example{
		client: examples.GetClient(),
	}

	examples.RunExamples(
		example.GetMovieProviders,
		example.GetMovieProvidersWithOptions,
		example.GetTVProviders,
		example.GetTVProvidersWithOptions,
		example.GetProviderRegions,
		example.GetProviderRegionsWithOptions,
	)
}
