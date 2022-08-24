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

func (e example) GetAccountStates() {
	states, _, err := e.client.Movies.GetAccountStates(531428, sessionID)
	examples.PanicOnError(err)
	examples.PrettyPrint(*states)
}

func (e example) Rate() {
	sessionID := tmdb.Auth{
		SessionID: sessionID,
	}
	states, _, err := e.client.Movies.Rate(28031, 8.5, sessionID)
	examples.PanicOnError(err)
	examples.PrettyPrint(*states)
}

func (e example) DeleteRating() {
	sessionID := tmdb.Auth{
		SessionID: sessionID,
	}
	states, _, err := e.client.Movies.DeleteRating(28031, sessionID)
	examples.PanicOnError(err)
	examples.PrettyPrint(*states)
}

func (e example) GetMovie() {
	opt := tmdb.MovieDetailsOptions{
		AppendToResponse: "alternative_titles,changes,credits,external_ids,images,keywords,lists,recommendations," +
			"release_dates,reviews,similar,translations,videos",
	}
	movie, _, err := e.client.Movies.GetMovie(430602, &opt)
	examples.PanicOnError(err)
	examples.PrettyPrint(*movie)
}

func (e example) GetAlternativeTitles() {
	opt := tmdb.MovieAlternativeTitlesOptions{
		Country: "BR",
	}
	titles, _, err := e.client.Movies.GetAlternativeTitles(597219, &opt)
	examples.PanicOnError(err)
	examples.PrettyPrint(*titles)
}

func (e example) GetChanges() {
	opt := tmdb.ChangesOptions{
		StartDate: "2021-09-15",
	}
	titles, _, err := e.client.Movies.GetChanges(19316, &opt)
	examples.PanicOnError(err)
	examples.PrettyPrint(*titles)
}

func (e example) GetCredits() {
	credits, _, err := e.client.Movies.GetCredits(19316, nil)
	examples.PanicOnError(err)
	examples.PrettyPrint(*credits)
}

func (e example) GetExternalIDs() {
	ids, _, err := e.client.Movies.GetExternalIDs(540)
	examples.PanicOnError(err)
	examples.PrettyPrint(*ids)
}

func (e example) GetImages() {
	opt := tmdb.ImagesOptions{
		Language:             "pt-BR",
		IncludeImageLanguage: "null,en",
	}
	images, _, err := e.client.Movies.GetImages(402900, &opt)
	examples.PanicOnError(err)
	examples.PrettyPrint(*images)
}

func (e example) GetKeywords() {
	keywords, _, err := e.client.Movies.GetKeywords(454889)
	examples.PanicOnError(err)
	examples.PrettyPrint(*keywords)
}

func (e example) GetLists() {
	lists, _, err := e.client.Movies.GetLists(860159, nil)
	examples.PanicOnError(err)
	examples.PrettyPrint(*lists)
}

func (e example) GetRecommendations() {
	movies, _, err := e.client.Movies.GetRecommendations(860159, nil)
	examples.PanicOnError(err)
	examples.PrettyPrint(*movies)
}

func (e example) GetReleaseDates() {
	dates, _, err := e.client.Movies.GetReleaseDates(65229)
	examples.PanicOnError(err)
	examples.PrettyPrint(*dates)
}

func (e example) GetReviews() {
	reviews, _, err := e.client.Movies.GetReviews(44214, nil)
	examples.PanicOnError(err)
	examples.PrettyPrint(*reviews)
}

func (e example) GetSimilar() {
	movies, _, err := e.client.Movies.GetSimilar(71325, nil)
	examples.PanicOnError(err)
	examples.PrettyPrint(*movies)
}

func (e example) GetTranslations() {
	translations, _, err := e.client.Movies.GetTranslations(11634)
	examples.PanicOnError(err)
	examples.PrettyPrint(*translations)
}

func (e example) GetVideos() {
	videos, _, err := e.client.Movies.GetVideos(664300, nil)
	examples.PanicOnError(err)
	examples.PrettyPrint(*videos)
}

func (e example) GetWatchProviders() {
	providers, _, err := e.client.Movies.GetWatchProviders(18620)
	examples.PanicOnError(err)
	examples.PrettyPrint(*providers)
}

func (e example) GetMoviesChanges() {
	opt := tmdb.ChangesOptions{
		StartDate: "2020-03-26",
	}
	changes, _, err := e.client.Movies.GetMoviesChanges(&opt)
	examples.PanicOnError(err)
	examples.PrettyPrint(*changes)
}

func (e example) GetLatest() {
	latest, _, err := e.client.Movies.GetLatest(nil)
	examples.PanicOnError(err)
	examples.PrettyPrint(*latest)
}

func (e example) GetNowPlaying() {
	movies, _, err := e.client.Movies.GetNowPlaying(nil)
	examples.PanicOnError(err)
	examples.PrettyPrint(*movies)
}

func (e example) GetPopular() {
	movies, _, err := e.client.Movies.GetPopular(nil)
	examples.PanicOnError(err)
	examples.PrettyPrint(*movies)
}

func (e example) GetTopRated() {
	movies, _, err := e.client.Movies.GetTopRated(nil)
	examples.PanicOnError(err)
	examples.PrettyPrint(*movies)
}

func (e example) GetUpcoming() {
	movies, _, err := e.client.Movies.GetUpcoming(nil)
	examples.PanicOnError(err)
	examples.PrettyPrint(*movies)
}

func main() {
	example := example{
		client: examples.GetClient(),
	}

	examples.RunExamples(
		example.GetAccountStates,     // 1
		example.Rate,                 // 2
		example.DeleteRating,         // 3
		example.GetMovie,             // 4
		example.GetAlternativeTitles, // 5
		example.GetChanges,           // 6
		example.GetCredits,           // 7
		example.GetExternalIDs,       // 8
		example.GetImages,            // 9
		example.GetKeywords,          // 10
		example.GetLists,             // 11
		example.GetRecommendations,   // 12
		example.GetReleaseDates,      // 13
		example.GetReviews,           // 14
		example.GetSimilar,           // 15
		example.GetTranslations,      // 16
		example.GetVideos,            // 17
		example.GetWatchProviders,    // 18
		example.GetMoviesChanges,     // 19
		example.GetLatest,            // 20
		example.GetNowPlaying,        // 21
		example.GetPopular,           // 22
		example.GetTopRated,          // 23
		example.GetUpcoming,          // 24
	)
}
