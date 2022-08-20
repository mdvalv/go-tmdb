package main

import (
	"os"

	"github.com/mdvalv/go-tmdb"
	"github.com/mdvalv/go-tmdb/examples"
)

type example struct {
	client *tmdb.Client
}

const (
	accountId = 123
)

var sessionId = os.Getenv("SESSIONID")

func (e example) GetAccount() {
	account, _, err := e.client.Account.GetAccount(sessionId)
	examples.PanicOnError(err)
	examples.PrettyPrint(*account)
}

func (e example) GetCreatedLists() {
	account, _, err := e.client.Account.GetCreatedLists(accountId, sessionId, nil)
	examples.PanicOnError(err)
	examples.PrettyPrint(*account)
}

func (e example) GetCreatedListsWithOptions() {
	opt := tmdb.AccountListsOptions{
		Language: "pt-BR",
	}
	account, _, err := e.client.Account.GetCreatedLists(accountId, sessionId, &opt)
	examples.PanicOnError(err)
	examples.PrettyPrint(*account)
}

func (e example) GetFavoriteMovies() {
	movies, _, err := e.client.Account.GetFavoriteMovies(accountId, sessionId, nil)
	examples.PanicOnError(err)
	examples.PrettyPrint(*movies)
}

func (e example) GetFavoriteMoviesWithOptions() {
	opt := tmdb.AccountOptions{
		Language: "pt-BR",
		SortBy:   "created_at.desc",
	}
	movies, _, err := e.client.Account.GetFavoriteMovies(accountId, sessionId, &opt)
	examples.PanicOnError(err)
	examples.PrettyPrint(*movies)
}

func (e example) GetFavoriteTVShows() {
	movies, _, err := e.client.Account.GetFavoriteTVShows(accountId, sessionId, nil)
	examples.PanicOnError(err)
	examples.PrettyPrint(*movies)
}

func (e example) GetFavoriteTVShowsWithOptions() {
	opt := tmdb.AccountOptions{
		Language: "pt-BR",
		SortBy:   "created_at.desc",
	}
	movies, _, err := e.client.Account.GetFavoriteTVShows(accountId, sessionId, &opt)
	examples.PanicOnError(err)
	examples.PrettyPrint(*movies)
}

func (e example) GetRatedMovies() {
	movies, _, err := e.client.Account.GetRatedMovies(accountId, sessionId, nil)
	examples.PanicOnError(err)
	examples.PrettyPrint(*movies)
}

func (e example) GetRatedMoviesWithOptions() {
	opt := tmdb.AccountOptions{
		Language: "pt-BR",
		SortBy:   "created_at.desc",
	}
	movies, _, err := e.client.Account.GetRatedMovies(accountId, sessionId, &opt)
	examples.PanicOnError(err)
	examples.PrettyPrint(*movies)
}

func (e example) GetRatedTVShows() {
	tvShows, _, err := e.client.Account.GetRatedTVShows(accountId, sessionId, nil)
	examples.PanicOnError(err)
	examples.PrettyPrint(*tvShows)
}

func (e example) GetRatedTVShowsWithOptions() {
	opt := tmdb.AccountOptions{
		Language: "pt-BR",
		SortBy:   "created_at.desc",
	}
	tvShows, _, err := e.client.Account.GetRatedTVShows(accountId, sessionId, &opt)
	examples.PanicOnError(err)
	examples.PrettyPrint(*tvShows)
}

func (e example) GetRatedTVEpisodes() {
	episodes, _, err := e.client.Account.GetRatedTVEpisodes(accountId, sessionId, nil)
	examples.PanicOnError(err)
	examples.PrettyPrint(*episodes)
}

func (e example) GetRatedTVEpisodesWithOptions() {
	opt := tmdb.AccountOptions{
		Language: "pt-BR",
		SortBy:   "created_at.desc",
	}
	episodes, _, err := e.client.Account.GetRatedTVEpisodes(accountId, sessionId, &opt)
	examples.PanicOnError(err)
	examples.PrettyPrint(*episodes)
}

func (e example) GetWatchlistMovies() {
	movies, _, err := e.client.Account.GetWatchlistMovies(accountId, sessionId, nil)
	examples.PanicOnError(err)
	examples.PrettyPrint(*movies)
}

func (e example) GetWatchlistMoviesWithOptions() {
	opt := tmdb.AccountOptions{
		Language: "pt-BR",
		SortBy:   "created_at.desc",
	}
	movies, _, err := e.client.Account.GetWatchlistMovies(accountId, sessionId, &opt)
	examples.PanicOnError(err)
	examples.PrettyPrint(*movies)
}

func (e example) GetWatchlistTVShows() {
	tvShows, _, err := e.client.Account.GetWatchlistTVShows(accountId, sessionId, nil)
	examples.PanicOnError(err)
	examples.PrettyPrint(*tvShows)
}

func (e example) GetWatchlistTVShowsWithOptions() {
	opt := tmdb.AccountOptions{
		Language: "pt-BR",
		SortBy:   "created_at.desc",
	}
	tvShows, _, err := e.client.Account.GetWatchlistTVShows(accountId, sessionId, &opt)
	examples.PanicOnError(err)
	examples.PrettyPrint(*tvShows)
}

func (e example) Favorite() {
	opt := tmdb.Favorite{
		Favorite:  true,
		MediaId:   458723,
		MediaType: "movie",
	}
	movies, _, err := e.client.Account.Favorite(accountId, sessionId, opt)
	examples.PanicOnError(err)
	examples.PrettyPrint(*movies)
}

func (e example) Watchlist() {
	opt := tmdb.Watchlist{
		Watchlist: true,
		MediaId:   776503,
		MediaType: "movie",
	}
	movies, _, err := e.client.Account.Watchlist(accountId, sessionId, opt)
	examples.PanicOnError(err)
	examples.PrettyPrint(*movies)
}

func main() {
	example := example{
		client: examples.GetClient(),
	}

	examples.RunExamples(
		example.GetAccount,                     // 1
		example.GetCreatedLists,                // 2
		example.GetCreatedListsWithOptions,     // 3
		example.GetFavoriteMovies,              // 4
		example.GetFavoriteMoviesWithOptions,   // 5
		example.GetFavoriteTVShows,             // 6
		example.GetFavoriteTVShowsWithOptions,  // 7
		example.GetRatedMovies,                 // 8
		example.GetRatedMoviesWithOptions,      // 9
		example.GetRatedTVShows,                // 10
		example.GetRatedTVShowsWithOptions,     // 11
		example.GetRatedTVEpisodes,             // 12
		example.GetRatedTVEpisodesWithOptions,  // 13
		example.GetWatchlistMovies,             // 14
		example.GetWatchlistMoviesWithOptions,  // 15
		example.GetWatchlistTVShows,            // 16
		example.GetWatchlistTVShowsWithOptions, // 17
		example.Favorite,                       // 18
		example.Watchlist,                      // 19
	)
}
