package main

import (
	"github.com/mdvalv/go-tmdb"
	"github.com/mdvalv/go-tmdb/examples"
	"github.com/pkg/errors"
)

type example struct {
	client *tmdb.Client
}

func (e example) GetAPIConfiguration() {
	configuration, _, err := e.client.Configuration.GetAPIConfiguration()
	if err != nil {
		panic(errors.Wrap(err, "failed to get configuration"))
	}
	examples.PrettyPrint(*configuration)
}

func (e example) GetCountries() {
	countries, _, err := e.client.Configuration.GetCountries()
	if err != nil {
		panic(errors.Wrap(err, "failed to get countries"))
	}
	examples.PrettyPrint(countries)
}

func (e example) GetJobs() {
	jobs, _, err := e.client.Configuration.GetJobs()
	if err != nil {
		panic(errors.Wrap(err, "failed to get jobs"))
	}
	examples.PrettyPrint(jobs)
}

func (e example) GetLanguages() {
	languages, _, err := e.client.Configuration.GetLanguages()
	if err != nil {
		panic(errors.Wrap(err, "failed to get languages"))
	}
	examples.PrettyPrint(languages)
}

func (e example) GetPrimaryTranslations() {
	translations, _, err := e.client.Configuration.GetPrimaryTranslations()
	if err != nil {
		panic(errors.Wrap(err, "failed to get translations"))
	}
	examples.PrettyPrint(translations)
}

func (e example) GetTimezones() {
	timezones, _, err := e.client.Configuration.GetTimezones()
	if err != nil {
		panic(errors.Wrap(err, "failed to get timezones"))
	}
	examples.PrettyPrint(timezones)
}

func main() {
	example := example{
		client: examples.GetClient(),
	}

	examples.RunExamples(
		example.GetAPIConfiguration,
		example.GetCountries,
		example.GetJobs,
		example.GetLanguages,
		example.GetPrimaryTranslations,
		example.GetTimezones,
	)
}
