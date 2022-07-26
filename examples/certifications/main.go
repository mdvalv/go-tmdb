package main

import (
	"github.com/mdvalv/go-tmdb"
	"github.com/mdvalv/go-tmdb/examples"
	"github.com/pkg/errors"
)

type example struct {
	client *tmdb.Client
}

func (e example) GetMovieCertifications() {
	certs, _, err := e.client.Certifications.GetMovieCertifications()
	if err != nil {
		panic(errors.Wrap(err, "failed to get movie certifications"))
	}
	examples.PrettyPrint(*certs)
}

func (e example) GetTVCertifications() {
	certs, _, err := e.client.Certifications.GetTVCertifications()
	if err != nil {
		panic(errors.Wrap(err, "failed to get tv certifications"))
	}
	examples.PrettyPrint(*certs)
}

func main() {
	example := example{
		client: examples.GetClient(),
	}

	examples.RunExamples(
		example.GetMovieCertifications,
		example.GetTVCertifications,
	)
}
