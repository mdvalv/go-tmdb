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

func (e example) GetEpisode() {
	opt := tmdb.TVEpisodeDetailsOptions{
		AppendToResponse: "account_states,credits,external_ids,images,rating,translations,videos",
	}
	episode, _, err := e.client.TVEpisodes.GetEpisode(69630, 1, 9, &opt)
	if err != nil {
		panic(errors.Wrap(err, "failed to get episode"))
	}
	examples.PrettyPrint(*episode)
}

func (e example) GetAccountStates() {
	states, _, err := e.client.TVEpisodes.GetAccountStates(31420, 2, 6, sessionId)
	if err != nil {
		panic(errors.Wrap(err, "failed to get account states"))
	}
	examples.PrettyPrint(*states)
}

func (e example) GetChanges() {
	opt := tmdb.ChangesOptions{
		StartDate: "2021-10-07",
	}
	titles, _, err := e.client.TVEpisodes.GetChanges(1997017, &opt)
	if err != nil {
		panic(errors.Wrap(err, "failed to get changes"))
	}
	examples.PrettyPrint(*titles)
}

func (e example) GetCredits() {
	credits, _, err := e.client.TVEpisodes.GetCredits(98888, 1, 4, nil)
	if err != nil {
		panic(errors.Wrap(err, "failed to get credits"))
	}
	examples.PrettyPrint(*credits)
}

func (e example) GetExternalIds() {
	ids, _, err := e.client.TVEpisodes.GetExternalIds(60705, 3, 8, nil)
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
	images, _, err := e.client.TVEpisodes.GetImages(47039, 2, 6, &opt)
	if err != nil {
		panic(errors.Wrap(err, "failed to get tv show images"))
	}
	examples.PrettyPrint(*images)
}

func (e example) GetTranslations() {
	translations, _, err := e.client.TVEpisodes.GetTranslations(68035, 1, 1)
	if err != nil {
		panic(errors.Wrap(err, "failed to get translations"))
	}
	examples.PrettyPrint(*translations)
}

func (e example) GetVideos() {
	videos, _, err := e.client.TVEpisodes.GetVideos(1399, 1, 2, nil)
	if err != nil {
		panic(errors.Wrap(err, "failed to get videos"))
	}
	examples.PrettyPrint(*videos)
}

func (e example) Rate() {
	sessionId := tmdb.Auth{
		SessionId: sessionId,
	}
	states, _, err := e.client.TVEpisodes.Rate(106431, 1, 1, 8.5, sessionId)
	if err != nil {
		panic(errors.Wrap(err, "failed to rate tv show episode"))
	}
	examples.PrettyPrint(*states)
}

func (e example) DeleteRating() {
	sessionId := tmdb.Auth{
		SessionId: sessionId,
	}
	states, _, err := e.client.TVEpisodes.DeleteRating(106431, 1, 1, sessionId)
	if err != nil {
		panic(errors.Wrap(err, "failed to delete rating"))
	}
	examples.PrettyPrint(*states)
}

func main() {
	example := example{
		client: examples.GetClient(),
	}

	examples.RunExamples(
		example.GetEpisode,       // 1
		example.GetAccountStates, // 2
		example.GetChanges,       // 3
		example.GetCredits,       // 4
		example.GetExternalIds,   // 5
		example.GetImages,        // 6
		example.GetTranslations,  // 7
		example.GetVideos,        // 8
		example.Rate,             // 9
		example.DeleteRating,     // 10
	)
}
