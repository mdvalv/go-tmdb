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

func (e example) GetTVShow() {
	opt := tmdb.TVShowDetailsOptions{
		AppendToResponse: "aggregate_credits,alternative_titles,changes,content_ratings,credits,episode_groups," +
			"external_ids,images,keywords,recommendations,reviews,screened_theatrically,similar,translations,videos",
	}
	tvShow, _, err := e.client.TV.GetTVShow(107005, &opt)
	if err != nil {
		panic(errors.Wrap(err, "failed to get tv show"))
	}
	examples.PrettyPrint(*tvShow)
}

func (e example) GetAccountStates() {
	states, _, err := e.client.TV.GetAccountStates(47801, sessionId)
	if err != nil {
		panic(errors.Wrap(err, "failed to get account states"))
	}
	examples.PrettyPrint(*states)
}

func (e example) GetAggregateCredits() {
	states, _, err := e.client.TV.GetAggregateCredits(76148, nil)
	if err != nil {
		panic(errors.Wrap(err, "failed to get aggregate credits"))
	}
	examples.PrettyPrint(*states)
}

func (e example) GetAlternativeTitles() {
	titles, _, err := e.client.TV.GetAlternativeTitles(66732, nil)
	if err != nil {
		panic(errors.Wrap(err, "failed to get alternative titles"))
	}
	examples.PrettyPrint(*titles)
}

func (e example) GetChanges() {
	opt := tmdb.ChangesOptions{
		StartDate: "2018-05-05",
	}
	titles, _, err := e.client.TV.GetChanges(76438, &opt)
	if err != nil {
		panic(errors.Wrap(err, "failed to get tv show changes"))
	}
	examples.PrettyPrint(*titles)
}

func (e example) GetContentRatings() {
	ratings, _, err := e.client.TV.GetContentRatings(61664, nil)
	if err != nil {
		panic(errors.Wrap(err, "failed to get content ratings"))
	}
	examples.PrettyPrint(*ratings)
}

func (e example) GetCredits() {
	credits, _, err := e.client.TV.GetCredits(90766, nil)
	if err != nil {
		panic(errors.Wrap(err, "failed to get tv show credits"))
	}
	examples.PrettyPrint(*credits)
}

func (e example) GetEpisodeGroups() {
	groups, _, err := e.client.TV.GetEpisodeGroups(30983, nil)
	if err != nil {
		panic(errors.Wrap(err, "failed to get tv show episode groups"))
	}
	examples.PrettyPrint(*groups)
}

func (e example) GetExternalIds() {
	ids, _, err := e.client.TV.GetExternalIds(4616, nil)
	if err != nil {
		panic(errors.Wrap(err, "failed to get tv show external ids"))
	}
	examples.PrettyPrint(*ids)
}

func (e example) GetImages() {
	opt := tmdb.ImagesOptions{
		Language:             "pt-BR",
		IncludeImageLanguage: "null,en",
	}
	images, _, err := e.client.TV.GetImages(91977, &opt)
	if err != nil {
		panic(errors.Wrap(err, "failed to get tv show images"))
	}
	examples.PrettyPrint(*images)
}

func (e example) GetKeywords() {
	keywords, _, err := e.client.TV.GetKeywords(65320)
	if err != nil {
		panic(errors.Wrap(err, "failed to get tv show keywords"))
	}
	examples.PrettyPrint(*keywords)
}

func (e example) GetRecommendations() {
	movies, _, err := e.client.TV.GetRecommendations(79084, nil)
	if err != nil {
		panic(errors.Wrap(err, "failed to get tv show recommendations"))
	}
	examples.PrettyPrint(*movies)
}

func (e example) GetReviews() {
	reviews, _, err := e.client.TV.GetReviews(1399, nil)
	if err != nil {
		panic(errors.Wrap(err, "failed to get tv show reviews"))
	}
	examples.PrettyPrint(*reviews)
}

func (e example) GetScreenedTheatrically() {
	screenedTheatrically, _, err := e.client.TV.GetScreenedTheatrically(68716)
	if err != nil {
		panic(errors.Wrap(err, "failed to get screened theatrically info"))
	}
	examples.PrettyPrint(*screenedTheatrically)
}

func (e example) GetAiringToday() {
	tvShows, _, err := e.client.TV.GetAiringToday(nil)
	if err != nil {
		panic(errors.Wrap(err, "failed to get airing today"))
	}
	examples.PrettyPrint(*tvShows)
}

func (e example) GetTVShowsChanges() {
	changes, _, err := e.client.TV.GetTVShowsChanges(nil)
	if err != nil {
		panic(errors.Wrap(err, "failed to get tv shows changes"))
	}
	examples.PrettyPrint(*changes)
}

func (e example) GetOnTheAir() {
	changes, _, err := e.client.TV.GetOnTheAir(nil)
	if err != nil {
		panic(errors.Wrap(err, "failed to get on the air"))
	}
	examples.PrettyPrint(*changes)
}

func (e example) GetSimilar() {
	similar, _, err := e.client.TV.GetSimilar(33852, nil)
	if err != nil {
		panic(errors.Wrap(err, "failed to get similar"))
	}
	examples.PrettyPrint(*similar)
}

func (e example) GetTranslations() {
	translations, _, err := e.client.TV.GetTranslations(61056)
	if err != nil {
		panic(errors.Wrap(err, "failed to get translations"))
	}
	examples.PrettyPrint(*translations)
}

func (e example) GetVideos() {
	videos, _, err := e.client.TV.GetVideos(56296, nil)
	if err != nil {
		panic(errors.Wrap(err, "failed to get tv show videos"))
	}
	examples.PrettyPrint(*videos)
}

func (e example) GetWatchProviders() {
	providers, _, err := e.client.TV.GetWatchProviders(42009)
	if err != nil {
		panic(errors.Wrap(err, "failed to get tv show watch providers"))
	}
	examples.PrettyPrint(*providers)
}

func (e example) GetLatest() {
	latest, _, err := e.client.TV.GetLatest(nil)
	if err != nil {
		panic(errors.Wrap(err, "failed to get latest tv show"))
	}
	examples.PrettyPrint(*latest)
}

func (e example) GetPopular() {
	popular, _, err := e.client.TV.GetPopular(nil)
	if err != nil {
		panic(errors.Wrap(err, "failed to get popular tv shows"))
	}
	examples.PrettyPrint(*popular)
}

func (e example) GetTopRated() {
	topRated, _, err := e.client.TV.GetTopRated(nil)
	if err != nil {
		panic(errors.Wrap(err, "failed to get top rated tv shows"))
	}
	examples.PrettyPrint(*topRated)
}

func (e example) GetEpisodeGroup() {
	group, _, err := e.client.TV.GetEpisodeGroup("5acf93e60e0a26346d0000ce", nil)
	if err != nil {
		panic(errors.Wrap(err, "failed to get episode group"))
	}
	examples.PrettyPrint(*group)
}

func (e example) Rate() {
	sessionId := tmdb.Auth{
		SessionId: sessionId,
	}
	states, _, err := e.client.TV.Rate(31917, 7.5, sessionId)
	if err != nil {
		panic(errors.Wrap(err, "failed to rate tv show"))
	}
	examples.PrettyPrint(*states)
}

func (e example) DeleteRating() {
	sessionId := tmdb.Auth{
		SessionId: sessionId,
	}
	states, _, err := e.client.TV.DeleteRating(31917, sessionId)
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
		example.GetTVShow,               // 1
		example.GetAccountStates,        // 2
		example.GetAggregateCredits,     // 3
		example.GetAlternativeTitles,    // 4
		example.GetChanges,              // 5
		example.GetContentRatings,       // 6
		example.GetCredits,              // 7
		example.GetEpisodeGroups,        // 8
		example.GetExternalIds,          // 9
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
