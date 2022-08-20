package main

import (
	"github.com/mdvalv/go-tmdb"
	"github.com/mdvalv/go-tmdb/examples"
)

type example struct {
	client *tmdb.Client
}

func (e example) GetCompany() {
	company, _, err := e.client.Companies.GetCompany(1)
	examples.PanicOnError(err)
	examples.PrettyPrint(*company)
}

func (e example) GetCompanyAlternativeNames() {
	names, _, err := e.client.Companies.GetAlternativeNames(3)
	examples.PanicOnError(err)
	examples.PrettyPrint(*names)
}

func (e example) GetCompanyImages() {
	images, _, err := e.client.Companies.GetImages(3)
	examples.PanicOnError(err)
	examples.PrettyPrint(*images)
}

func main() {
	example := example{
		client: examples.GetClient(),
	}

	examples.RunExamples(
		example.GetCompany,                 // 1
		example.GetCompanyAlternativeNames, // 2
		example.GetCompanyImages,           // 3
	)
}
