package main

import (
	"github.com/mdvalv/go-tmdb"
	"github.com/mdvalv/go-tmdb/examples"
	"github.com/pkg/errors"
)

type example struct {
	client *tmdb.Client
}

func (e example) GetCompany() {
	company, _, err := e.client.Companies.GetCompany(1)
	if err != nil {
		panic(errors.Wrap(err, "failed to get company"))
	}
	examples.PrettyPrint(*company)
}

func (e example) GetCompanyAlternativeNames() {
	names, _, err := e.client.Companies.GetAlternativeNames(3)
	if err != nil {
		panic(errors.Wrap(err, "failed to get company alternative names"))
	}
	examples.PrettyPrint(*names)
}

func (e example) GetCompanyImages() {
	images, _, err := e.client.Companies.GetImages(3)
	if err != nil {
		panic(errors.Wrap(err, "failed to get company images"))
	}
	examples.PrettyPrint(*images)
}

func main() {
	example := example{
		client: examples.GetClient(),
	}

	examples.RunExamples(
		example.GetCompany,
		example.GetCompanyAlternativeNames,
		example.GetCompanyImages,
	)
}
