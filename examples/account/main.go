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
	accountId = 123
	sessionId = "session_id"
)

func (e example) GetAccount() {
	account, _, err := e.client.Account.GetAccount(sessionId)
	if err != nil {
		panic(errors.Wrap(err, "failed to get account"))
	}
	examples.PrettyPrint(*account)
}

func (e example) GetCreatedLists() {
	account, _, err := e.client.Account.GetCreatedLists(accountId, sessionId, nil)
	if err != nil {
		panic(errors.Wrap(err, "failed to get created lists"))
	}
	examples.PrettyPrint(*account)
}

func (e example) GetCreatedListsWithOptions() {
	opt := tmdb.AccountListsOptions{
		Language: "pt-BR",
	}
	account, _, err := e.client.Account.GetCreatedLists(accountId, sessionId, &opt)
	if err != nil {
		panic(errors.Wrap(err, "failed to get created lists"))
	}
	examples.PrettyPrint(*account)
}

func (e example) GetFavoriteMovies() {
	movies, _, err := e.client.Account.GetFavoriteMovies(accountId, sessionId, nil)
	if err != nil {
		panic(errors.Wrap(err, "failed to get favorite movies"))
	}
	examples.PrettyPrint(*movies)
}

func (e example) GetFavoriteMoviesWithOptions() {
	opt := tmdb.AccountOptions{
		Language: "pt-BR",
		SortBy:   "created_at.desc",
	}
	movies, _, err := e.client.Account.GetFavoriteMovies(accountId, sessionId, &opt)
	if err != nil {
		panic(errors.Wrap(err, "failed to get favorite movies"))
	}
	examples.PrettyPrint(*movies)
}

func (e example) GetFavoriteTVShows() {
	movies, _, err := e.client.Account.GetFavoriteTVShows(accountId, sessionId, nil)
	if err != nil {
		panic(errors.Wrap(err, "failed to get favorite movies"))
	}
	examples.PrettyPrint(*movies)
}

func (e example) GetFavoriteTVShowsWithOptions() {
	opt := tmdb.AccountOptions{
		Language: "pt-BR",
		SortBy:   "created_at.desc",
	}
	movies, _, err := e.client.Account.GetFavoriteTVShows(accountId, sessionId, &opt)
	if err != nil {
		panic(errors.Wrap(err, "failed to get favorite movies"))
	}
	examples.PrettyPrint(*movies)
}

func (e example) GetRatedMovies() {
	movies, _, err := e.client.Account.GetRatedMovies(accountId, sessionId, nil)
	if err != nil {
		panic(errors.Wrap(err, "failed to get rated movies"))
	}
	examples.PrettyPrint(*movies)
}

func (e example) GetRatedMoviesWithOptions() {
	opt := tmdb.AccountOptions{
		Language: "pt-BR",
		SortBy:   "created_at.desc",
	}
	movies, _, err := e.client.Account.GetRatedMovies(accountId, sessionId, &opt)
	if err != nil {
		panic(errors.Wrap(err, "failed to get rated movies"))
	}
	examples.PrettyPrint(*movies)
}

func (e example) GetRatedTVShows() {
	tvShows, _, err := e.client.Account.GetRatedTVShows(accountId, sessionId, nil)
	if err != nil {
		panic(errors.Wrap(err, "failed to get rated tv shows"))
	}
	examples.PrettyPrint(*tvShows)
}

func (e example) GetRatedTVShowsWithOptions() {
	opt := tmdb.AccountOptions{
		Language: "pt-BR",
		SortBy:   "created_at.desc",
	}
	tvShows, _, err := e.client.Account.GetRatedTVShows(accountId, sessionId, &opt)
	if err != nil {
		panic(errors.Wrap(err, "failed to get rated tv shows"))
	}
	examples.PrettyPrint(*tvShows)
}

func (e example) GetRatedTVEpisodes() {
	episodes, _, err := e.client.Account.GetRatedTVEpisodes(accountId, sessionId, nil)
	if err != nil {
		panic(errors.Wrap(err, "failed to get rated tv episode"))
	}
	examples.PrettyPrint(*episodes)
}

func (e example) GetRatedTVEpisodesWithOptions() {
	opt := tmdb.AccountOptions{
		Language: "pt-BR",
		SortBy:   "created_at.desc",
	}
	episodes, _, err := e.client.Account.GetRatedTVEpisodes(accountId, sessionId, &opt)
	if err != nil {
		panic(errors.Wrap(err, "failed to get rated tv episode"))
	}
	examples.PrettyPrint(*episodes)
}

func (e example) GetWatchlistMovies() {
	movies, _, err := e.client.Account.GetWatchlistMovies(accountId, sessionId, nil)
	if err != nil {
		panic(errors.Wrap(err, "failed to get watchlist movies"))
	}
	examples.PrettyPrint(*movies)
}

func (e example) GetWatchlistMoviesWithOptions() {
	opt := tmdb.AccountOptions{
		Language: "pt-BR",
		SortBy:   "created_at.desc",
	}
	movies, _, err := e.client.Account.GetWatchlistMovies(accountId, sessionId, &opt)
	if err != nil {
		panic(errors.Wrap(err, "failed to get watchlist movies"))
	}
	examples.PrettyPrint(*movies)
}

func (e example) GetWatchlistTVShows() {
	tvShows, _, err := e.client.Account.GetWatchlistTVShows(accountId, sessionId, nil)
	if err != nil {
		panic(errors.Wrap(err, "failed to get watchlist tv shows"))
	}
	examples.PrettyPrint(*tvShows)
}

func (e example) GetWatchlistTVShowsWithOptions() {
	opt := tmdb.AccountOptions{
		Language: "pt-BR",
		SortBy:   "created_at.desc",
	}
	tvShows, _, err := e.client.Account.GetWatchlistTVShows(accountId, sessionId, &opt)
	if err != nil {
		panic(errors.Wrap(err, "failed to get watchlist tv shows"))
	}
	examples.PrettyPrint(*tvShows)
}

func (e example) Favorite() {
	opt := tmdb.Favorite{
		Favorite:  true,
		MediaId:   458723,
		MediaType: "movie",
	}
	movies, _, err := e.client.Account.Favorite(accountId, sessionId, opt)
	if err != nil {
		panic(errors.Wrap(err, "failed to mark as favorite"))
	}
	examples.PrettyPrint(*movies)
}

func (e example) Watchlist() {
	opt := tmdb.Watchlist{
		Watchlist: true,
		MediaId:   776503,
		MediaType: "movie",
	}
	movies, _, err := e.client.Account.Watchlist(accountId, sessionId, opt)
	if err != nil {
		panic(errors.Wrap(err, "failed to mark as favorite"))
	}
	examples.PrettyPrint(*movies)
}

func main() {
	example := example{
		client: examples.GetClient(),
	}

	examples.RunExamples(
		example.GetAccount,
		example.GetCreatedLists,
		example.GetCreatedListsWithOptions,
		example.GetFavoriteMovies,
		example.GetFavoriteMoviesWithOptions,
		example.GetFavoriteTVShows,
		example.GetFavoriteTVShowsWithOptions,
		example.GetRatedMovies,
		example.GetRatedMoviesWithOptions,
		example.GetRatedTVShows,
		example.GetRatedTVShowsWithOptions,
		example.GetRatedTVEpisodes,
		example.GetRatedTVEpisodesWithOptions,
		example.GetWatchlistMovies,
		example.GetWatchlistMoviesWithOptions,
		example.GetWatchlistTVShows,
		example.GetWatchlistTVShowsWithOptions,
		example.Favorite,
		example.Watchlist,
	)
}
