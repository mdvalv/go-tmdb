// Configuration examples.
package main

import (
	"github.com/mdvalv/go-tmdb"
	"github.com/mdvalv/go-tmdb/examples"
)

type example struct {
	client *tmdb.Client
}

func (e example) GetAPIConfiguration() {
	configuration, _, err := e.client.Configuration.GetAPIConfiguration()
	examples.PanicOnError(err)
	examples.PrettyPrint(*configuration)
}

func (e example) GetCountries() {
	countries, _, err := e.client.Configuration.GetCountries()
	examples.PanicOnError(err)
	examples.PrettyPrint(countries)
}

func (e example) GetJobs() {
	jobs, _, err := e.client.Configuration.GetJobs()
	examples.PanicOnError(err)
	examples.PrettyPrint(jobs)
}

func (e example) GetLanguages() {
	languages, _, err := e.client.Configuration.GetLanguages()
	examples.PanicOnError(err)
	examples.PrettyPrint(languages)
}

func (e example) GetPrimaryTranslations() {
	translations, _, err := e.client.Configuration.GetPrimaryTranslations()
	examples.PanicOnError(err)
	examples.PrettyPrint(translations)
}

func (e example) GetTimezones() {
	timezones, _, err := e.client.Configuration.GetTimezones()
	examples.PanicOnError(err)
	examples.PrettyPrint(timezones)
}

func main() {
	example := example{
		client: examples.GetClient(),
	}

	examples.RunExamples(
		example.GetAPIConfiguration,    // 1
		example.GetCountries,           // 2
		example.GetJobs,                // 3
		example.GetLanguages,           // 4
		example.GetPrimaryTranslations, // 5
		example.GetTimezones,           // 6
	)
}
