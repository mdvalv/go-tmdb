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
		example.GetFavoriteMovies,
		example.GetFavoriteMoviesWithOptions,
		example.GetFavoriteTVShows,
		example.GetFavoriteTVShowsWithOptions,
		example.Favorite,
		example.Watchlist,
	)
}
