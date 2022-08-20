package main

import (
	"github.com/mdvalv/go-tmdb"
	"github.com/mdvalv/go-tmdb/examples"
)

type example struct {
	client *tmdb.Client
}

func (e example) GetMovieCertifications() {
	certs, _, err := e.client.Certifications.GetMovieCertifications()
	examples.PanicOnError(err)
	examples.PrettyPrint(*certs)
}

func (e example) GetTVCertifications() {
	certs, _, err := e.client.Certifications.GetTVCertifications()
	examples.PanicOnError(err)
	examples.PrettyPrint(*certs)
}

func main() {
	example := example{
		client: examples.GetClient(),
	}

	examples.RunExamples(
		example.GetMovieCertifications, // 1
		example.GetTVCertifications,    // 2
	)
}
