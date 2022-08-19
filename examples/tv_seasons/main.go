package main

import (
	"os"

	"github.com/mdvalv/go-tmdb"
	"github.com/mdvalv/go-tmdb/examples"
	"github.com/pkg/errors"
)

type example struct {
	client *tmdb.Client
}

var sessionId = os.Getenv("SESSIONID")

func (e example) GetSeason() {
	opt := tmdb.TVSeasonDetailsOptions{
		AppendToResponse: "aggregate_credits,credits,external_ids,images,translations,videos",
	}
	season, _, err := e.client.TVSeasons.GetSeason(61849, 1, &opt)
	if err != nil {
		panic(errors.Wrap(err, "failed to get season"))
	}
	examples.PrettyPrint(*season)
}

func (e example) GetAccountStates() {
	states, _, err := e.client.TVSeasons.GetAccountStates(1424, 1, sessionId)
	if err != nil {
		panic(errors.Wrap(err, "failed to get account states for season"))
	}
	examples.PrettyPrint(*states)
}

func (e example) GetAggregateCredits() {
	states, _, err := e.client.TVSeasons.GetAggregateCredits(86163, 1, nil)
	if err != nil {
		panic(errors.Wrap(err, "failed to get aggregate credits"))
	}
	examples.PrettyPrint(*states)
}

func (e example) GetChanges() {
	opt := tmdb.ChangesOptions{
		StartDate: "2019-06-19",
	}
	titles, _, err := e.client.TVSeasons.GetChanges(119450, &opt)
	if err != nil {
		panic(errors.Wrap(err, "failed to get changes"))
	}
	examples.PrettyPrint(*titles)
}

func (e example) GetCredits() {
	credits, _, err := e.client.TVSeasons.GetCredits(13812, 1, nil)
	if err != nil {
		panic(errors.Wrap(err, "failed to get credits"))
	}
	examples.PrettyPrint(*credits)
}

func (e example) GetExternalIds() {
	ids, _, err := e.client.TVSeasons.GetExternalIds(67466, 1, nil)
	if err != nil {
		panic(errors.Wrap(err, "failed to get external ids"))
	}
	examples.PrettyPrint(*ids)
}

func (e example) GetImages() {
	opt := tmdb.ImagesOptions{
		Language:             "pt-BR",
		IncludeImageLanguage: "null,en",
	}
	images, _, err := e.client.TVSeasons.GetImages(61406, 1, &opt)
	if err != nil {
		panic(errors.Wrap(err, "failed to get tv show images"))
	}
	examples.PrettyPrint(*images)
}

func (e example) GetTranslations() {
	translations, _, err := e.client.TVSeasons.GetTranslations(407, 1)
	if err != nil {
		panic(errors.Wrap(err, "failed to get translations"))
	}
	examples.PrettyPrint(*translations)
}

func (e example) GetVideos() {
	videos, _, err := e.client.TVSeasons.GetVideos(76922, 6, nil)
	if err != nil {
		panic(errors.Wrap(err, "failed to get videos"))
	}
	examples.PrettyPrint(*videos)
}

func main() {
	example := example{
		client: examples.GetClient(),
	}

	examples.RunExamples(
		example.GetSeason,           // 1
		example.GetAccountStates,    // 2
		example.GetAggregateCredits, // 3
		example.GetChanges,          // 4
		example.GetCredits,          // 5
		example.GetExternalIds,      // 6
		example.GetImages,           // 7
		example.GetTranslations,     // 8
		example.GetVideos,           // 9
	)
}
