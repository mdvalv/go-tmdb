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

func (e example) GetTVShow() {
	opt := tmdb.TVShowDetailsOptions{
		AppendToResponse: "aggregate_credits,alternative_titles,changes,content_ratings,credits,episode_groups," +
			"external_ids,images,keywords,recommendations,reviews,screened_theatrically,similar,translations,videos",
	}
	tvShow, _, err := e.client.TV.GetTVShow(107005, &opt)
	examples.PanicOnError(err)
	examples.PrettyPrint(*tvShow)
}

func (e example) GetAccountStates() {
	states, _, err := e.client.TV.GetAccountStates(47801, sessionID)
	examples.PanicOnError(err)
	examples.PrettyPrint(*states)
}

func (e example) GetAggregateCredits() {
	states, _, err := e.client.TV.GetAggregateCredits(76148, nil)
	examples.PanicOnError(err)
	examples.PrettyPrint(*states)
}

func (e example) GetAlternativeTitles() {
	titles, _, err := e.client.TV.GetAlternativeTitles(66732, nil)
	examples.PanicOnError(err)
	examples.PrettyPrint(*titles)
}

func (e example) GetChanges() {
	opt := tmdb.ChangesOptions{
		StartDate: "2018-05-05",
	}
	titles, _, err := e.client.TV.GetChanges(76438, &opt)
	examples.PanicOnError(err)
	examples.PrettyPrint(*titles)
}

func (e example) GetContentRatings() {
	ratings, _, err := e.client.TV.GetContentRatings(61664, nil)
	examples.PanicOnError(err)
	examples.PrettyPrint(*ratings)
}

func (e example) GetCredits() {
	credits, _, err := e.client.TV.GetCredits(90766, nil)
	examples.PanicOnError(err)
	examples.PrettyPrint(*credits)
}

func (e example) GetEpisodeGroups() {
	groups, _, err := e.client.TV.GetEpisodeGroups(30983, nil)
	examples.PanicOnError(err)
	examples.PrettyPrint(*groups)
}

func (e example) GetExternalIDs() {
	ids, _, err := e.client.TV.GetExternalIDs(4616, nil)
	examples.PanicOnError(err)
	examples.PrettyPrint(*ids)
}

func (e example) GetImages() {
	opt := tmdb.ImagesOptions{
		Language:             "pt-BR",
		IncludeImageLanguage: "null,en",
	}
	images, _, err := e.client.TV.GetImages(91977, &opt)
	examples.PanicOnError(err)
	examples.PrettyPrint(*images)
}

func (e example) GetKeywords() {
	keywords, _, err := e.client.TV.GetKeywords(65320)
	examples.PanicOnError(err)
	examples.PrettyPrint(*keywords)
}

func (e example) GetRecommendations() {
	movies, _, err := e.client.TV.GetRecommendations(79084, nil)
	examples.PanicOnError(err)
	examples.PrettyPrint(*movies)
}

func (e example) GetReviews() {
	reviews, _, err := e.client.TV.GetReviews(1399, nil)
	examples.PanicOnError(err)
	examples.PrettyPrint(*reviews)
}

func (e example) GetScreenedTheatrically() {
	screenedTheatrically, _, err := e.client.TV.GetScreenedTheatrically(68716)
	examples.PanicOnError(err)
	examples.PrettyPrint(*screenedTheatrically)
}

func (e example) GetAiringToday() {
	tvShows, _, err := e.client.TV.GetAiringToday(nil)
	examples.PanicOnError(err)
	examples.PrettyPrint(*tvShows)
}

func (e example) GetTVShowsChanges() {
	changes, _, err := e.client.TV.GetTVShowsChanges(nil)
	examples.PanicOnError(err)
	examples.PrettyPrint(*changes)
}

func (e example) GetOnTheAir() {
	changes, _, err := e.client.TV.GetOnTheAir(nil)
	examples.PanicOnError(err)
	examples.PrettyPrint(*changes)
}

func (e example) GetSimilar() {
	similar, _, err := e.client.TV.GetSimilar(33852, nil)
	examples.PanicOnError(err)
	examples.PrettyPrint(*similar)
}

func (e example) GetTranslations() {
	translations, _, err := e.client.TV.GetTranslations(61056)
	examples.PanicOnError(err)
	examples.PrettyPrint(*translations)
}

func (e example) GetVideos() {
	videos, _, err := e.client.TV.GetVideos(56296, nil)
	examples.PanicOnError(err)
	examples.PrettyPrint(*videos)
}

func (e example) GetWatchProviders() {
	providers, _, err := e.client.TV.GetWatchProviders(42009)
	examples.PanicOnError(err)
	examples.PrettyPrint(*providers)
}

func (e example) GetLatest() {
	latest, _, err := e.client.TV.GetLatest(nil)
	examples.PanicOnError(err)
	examples.PrettyPrint(*latest)
}

func (e example) GetPopular() {
	popular, _, err := e.client.TV.GetPopular(nil)
	examples.PanicOnError(err)
	examples.PrettyPrint(*popular)
}

func (e example) GetTopRated() {
	topRated, _, err := e.client.TV.GetTopRated(nil)
	examples.PanicOnError(err)
	examples.PrettyPrint(*topRated)
}

func (e example) GetEpisodeGroup() {
	group, _, err := e.client.TV.GetEpisodeGroup("5acf93e60e0a26346d0000ce", nil)
	examples.PanicOnError(err)
	examples.PrettyPrint(*group)
}

func (e example) Rate() {
	sessionID := tmdb.Auth{
		SessionID: sessionID,
	}
	states, _, err := e.client.TV.Rate(31917, 7.5, sessionID)
	examples.PanicOnError(err)
	examples.PrettyPrint(*states)
}

func (e example) DeleteRating() {
	sessionID := tmdb.Auth{
		SessionID: sessionID,
	}
	states, _, err := e.client.TV.DeleteRating(31917, sessionID)
	examples.PanicOnError(err)
	examples.PrettyPrint(*states)
}

func main() {
	example := example{
		client: examples.GetClient(),
	}

	examples.RunExamples(
		example.GetTVShow,               // 1
		example.GetAccountStates,        // 2
		example.GetAggregateCredits,     // 3
		example.GetAlternativeTitles,    // 4
		example.GetChanges,              // 5
		example.GetContentRatings,       // 6
		example.GetCredits,              // 7
		example.GetEpisodeGroups,        // 8
		example.GetExternalIDs,          // 9
		example.GetImages,               // 10
		example.GetKeywords,             // 11
		example.GetRecommendations,      // 12
		example.GetReviews,              // 13
		example.GetScreenedTheatrically, // 14
		example.GetAiringToday,          // 15
		example.GetTVShowsChanges,       // 16
		example.GetOnTheAir,             // 17
		example.GetSimilar,              // 18
		example.GetTranslations,         // 19
		example.GetVideos,               // 20
		example.GetWatchProviders,       // 21
		example.GetLatest,               // 22
		example.GetPopular,              // 23
		example.GetTopRated,             // 24
		example.GetEpisodeGroup,         // 25
		example.Rate,                    // 26
		example.DeleteRating,            // 27
	)
}
