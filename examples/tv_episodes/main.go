package main

import (
	"os"

	"github.com/mdvalv/go-tmdb"
	"github.com/mdvalv/go-tmdb/examples"
)

type example struct {
	client *tmdb.Client
}

var sessionID = os.Getenv("SESSIONID")

func (e example) GetEpisode() {
	opt := tmdb.TVEpisodeDetailsOptions{
		AppendToResponse: "account_states,credits,external_ids,images,rating,translations,videos",
	}
	episode, _, err := e.client.TVEpisodes.GetEpisode(69630, 1, 9, &opt)
	examples.PanicOnError(err)
	examples.PrettyPrint(*episode)
}

func (e example) GetAccountStates() {
	states, _, err := e.client.TVEpisodes.GetAccountStates(31420, 2, 6, sessionID)
	examples.PanicOnError(err)
	examples.PrettyPrint(*states)
}

func (e example) GetChanges() {
	opt := tmdb.ChangesOptions{
		StartDate: "2021-10-07",
	}
	titles, _, err := e.client.TVEpisodes.GetChanges(1997017, &opt)
	examples.PanicOnError(err)
	examples.PrettyPrint(*titles)
}

func (e example) GetCredits() {
	credits, _, err := e.client.TVEpisodes.GetCredits(98888, 1, 4, nil)
	examples.PanicOnError(err)
	examples.PrettyPrint(*credits)
}

func (e example) GetExternalIDs() {
	ids, _, err := e.client.TVEpisodes.GetExternalIDs(60705, 3, 8, nil)
	examples.PanicOnError(err)
	examples.PrettyPrint(*ids)
}

func (e example) GetImages() {
	opt := tmdb.ImagesOptions{
		Language:             "pt-BR",
		IncludeImageLanguage: "null,en",
	}
	images, _, err := e.client.TVEpisodes.GetImages(47039, 2, 6, &opt)
	examples.PanicOnError(err)
	examples.PrettyPrint(*images)
}

func (e example) GetTranslations() {
	translations, _, err := e.client.TVEpisodes.GetTranslations(68035, 1, 1)
	examples.PanicOnError(err)
	examples.PrettyPrint(*translations)
}

func (e example) GetVideos() {
	videos, _, err := e.client.TVEpisodes.GetVideos(1399, 1, 2, nil)
	examples.PanicOnError(err)
	examples.PrettyPrint(*videos)
}

func (e example) Rate() {
	sessionID := tmdb.Auth{
		SessionID: sessionID,
	}
	states, _, err := e.client.TVEpisodes.Rate(106431, 1, 1, 8.5, sessionID)
	examples.PanicOnError(err)
	examples.PrettyPrint(*states)
}

func (e example) DeleteRating() {
	sessionID := tmdb.Auth{
		SessionID: sessionID,
	}
	states, _, err := e.client.TVEpisodes.DeleteRating(106431, 1, 1, sessionID)
	examples.PanicOnError(err)
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
		example.GetExternalIDs,   // 5
		example.GetImages,        // 6
		example.GetTranslations,  // 7
		example.GetVideos,        // 8
		example.Rate,             // 9
		example.DeleteRating,     // 10
	)
}
