package main

import (
	"github.com/mdvalv/go-tmdb"
	"github.com/mdvalv/go-tmdb/examples"
	"github.com/pkg/errors"
)

type example struct {
	client *tmdb.Client
}

const (
	sessionId = "session_id"
)

func (e example) GetAccountStates() {
	states, _, err := e.client.Movies.GetAccountStates(531428, sessionId)
	if err != nil {
		panic(errors.Wrap(err, "failed to get account states"))
	}
	examples.PrettyPrint(*states)
}

func (e example) Rate() {
	sessionId := tmdb.Auth{
		SessionId: sessionId,
	}
	states, _, err := e.client.Movies.Rate(8.5, 28031, sessionId)
	if err != nil {
		panic(errors.Wrap(err, "failed to rate movie"))
	}
	examples.PrettyPrint(*states)
}

func (e example) DeleteRating() {
	sessionId := tmdb.Auth{
		SessionId: sessionId,
	}
	states, _, err := e.client.Movies.DeleteRating(28031, sessionId)
	if err != nil {
		panic(errors.Wrap(err, "failed to delete rating"))
	}
	examples.PrettyPrint(*states)
}

func (e example) GetMovie() {
	opt := tmdb.MovieDetailsOptions{
		AppendToResponse: "alternative_titles,changes,credits,external_ids,images,keywords,lists,recommendations," +
			"release_dates,reviews,similar,translations,videos",
	}
	movie, _, err := e.client.Movies.GetMovie(430602, &opt)
	if err != nil {
		panic(errors.Wrap(err, "failed to get movie"))
	}
	examples.PrettyPrint(*movie)
}

func (e example) GetAlternativeTitles() {
	opt := tmdb.AlternativeTitlesOptions{
		Country: "BR",
	}
	titles, _, err := e.client.Movies.GetAlternativeTitles(597219, &opt)
	if err != nil {
		panic(errors.Wrap(err, "failed to get alternative titles"))
	}
	examples.PrettyPrint(*titles)
}

func (e example) GetChanges() {
	opt := tmdb.ChangesOptions{
		StartDate: "2021-09-15",
	}
	titles, _, err := e.client.Movies.GetChanges(19316, &opt)
	if err != nil {
		panic(errors.Wrap(err, "failed to get movie changes"))
	}
	examples.PrettyPrint(*titles)
}

func (e example) GetCredits() {
	credits, _, err := e.client.Movies.GetCredits(19316, nil)
	if err != nil {
		panic(errors.Wrap(err, "failed to get movie credits"))
	}
	examples.PrettyPrint(*credits)
}

func (e example) GetExternalIds() {
	ids, _, err := e.client.Movies.GetExternalIds(540)
	if err != nil {
		panic(errors.Wrap(err, "failed to get movie external ids"))
	}
	examples.PrettyPrint(*ids)
}

func (e example) GetImages() {
	opt := tmdb.MovieImagesOptions{
		Language:             "pt-BR",
		IncludeImageLanguage: "null,en",
	}
	images, _, err := e.client.Movies.GetImages(402900, &opt)
	if err != nil {
		panic(errors.Wrap(err, "failed to get movie images"))
	}
	examples.PrettyPrint(*images)
}

func (e example) GetKeywords() {
	keywords, _, err := e.client.Movies.GetKeywords(454889)
	if err != nil {
		panic(errors.Wrap(err, "failed to get movie keywords"))
	}
	examples.PrettyPrint(*keywords)
}

func (e example) GetLists() {
	lists, _, err := e.client.Movies.GetLists(860159, nil)
	if err != nil {
		panic(errors.Wrap(err, "failed to get movie lists"))
	}
	examples.PrettyPrint(*lists)
}

func (e example) GetRecommendations() {
	movies, _, err := e.client.Movies.GetRecommendations(860159, nil)
	if err != nil {
		panic(errors.Wrap(err, "failed to get movie recommendations"))
	}
	examples.PrettyPrint(*movies)
}

func (e example) GetReleaseDates() {
	dates, _, err := e.client.Movies.GetReleaseDates(65229)
	if err != nil {
		panic(errors.Wrap(err, "failed to get movie release dates"))
	}
	examples.PrettyPrint(*dates)
}

func (e example) GetReviews() {
	reviews, _, err := e.client.Movies.GetReviews(44214, nil)
	if err != nil {
		panic(errors.Wrap(err, "failed to get movie reviews"))
	}
	examples.PrettyPrint(*reviews)
}

func (e example) GetSimilar() {
	movies, _, err := e.client.Movies.GetSimilar(71325, nil)
	if err != nil {
		panic(errors.Wrap(err, "failed to get similar movies"))
	}
	examples.PrettyPrint(*movies)
}

func (e example) GetTranslations() {
	translations, _, err := e.client.Movies.GetTranslations(11634)
	if err != nil {
		panic(errors.Wrap(err, "failed to get movie translations"))
	}
	examples.PrettyPrint(*translations)
}

func (e example) GetVideos() {
	videos, _, err := e.client.Movies.GetVideos(664300, nil)
	if err != nil {
		panic(errors.Wrap(err, "failed to get movie videos"))
	}
	examples.PrettyPrint(*videos)
}

func (e example) GetWatchProviders() {
	providers, _, err := e.client.Movies.GetWatchProviders(18620)
	if err != nil {
		panic(errors.Wrap(err, "failed to get movie watch providers"))
	}
	examples.PrettyPrint(*providers)
}

func (e example) GetMoviesChanges() {
	opt := tmdb.ChangesOptions{
		StartDate: "2020-03-26",
	}
	changes, _, err := e.client.Movies.GetMoviesChanges(&opt)
	if err != nil {
		panic(errors.Wrap(err, "failed to get movies changes"))
	}
	examples.PrettyPrint(*changes)
}

func (e example) GetLatest() {
	latest, _, err := e.client.Movies.GetLatest(nil)
	if err != nil {
		panic(errors.Wrap(err, "failed to get movie latest"))
	}
	examples.PrettyPrint(*latest)
}

func (e example) GetNowPlaying() {
	movies, _, err := e.client.Movies.GetNowPlaying(nil)
	if err != nil {
		panic(errors.Wrap(err, "failed to get movies playing now"))
	}
	examples.PrettyPrint(*movies)
}

func (e example) GetPopular() {
	movies, _, err := e.client.Movies.GetPopular(nil)
	if err != nil {
		panic(errors.Wrap(err, "failed to get popular movies"))
	}
	examples.PrettyPrint(*movies)
}

func (e example) GetTopRated() {
	movies, _, err := e.client.Movies.GetTopRated(nil)
	if err != nil {
		panic(errors.Wrap(err, "failed to get top rated movies"))
	}
	examples.PrettyPrint(*movies)
}

func (e example) GetUpcoming() {
	movies, _, err := e.client.Movies.GetUpcoming(nil)
	if err != nil {
		panic(errors.Wrap(err, "failed to get upcoming movies"))
	}
	examples.PrettyPrint(*movies)
}

func main() {
	example := example{
		client: examples.GetClient(),
	}

	examples.RunExamples(
		example.GetAccountStates,
		example.Rate,
		example.DeleteRating,
		example.GetMovie,
		example.GetAlternativeTitles,
		example.GetChanges,
		example.GetCredits,
		example.GetExternalIds,
		example.GetImages,
		example.GetKeywords,
		example.GetLists,
		example.GetRecommendations,
		example.GetReleaseDates,
		example.GetReviews,
		example.GetSimilar,
		example.GetTranslations,
		example.GetVideos,
		example.GetWatchProviders,
		example.GetMoviesChanges,
		example.GetLatest,
		example.GetNowPlaying,
		example.GetPopular,
		example.GetTopRated,
		example.GetUpcoming,
	)
}
