package tmdb

import (
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

// AccountResource handles account-related requests of TMDb API.
type AccountResource struct {
	client *Client
}

type Account struct {
	Avatar struct {
		Gravatar struct {
			Hash string `json:"hash"`
		} `json:"gravatar"`
	} `json:"avatar"`
	Id           int    `json:"id"`
	ISO6391      string `json:"iso_639_1"`
	ISO31661     string `json:"iso_3166_1"`
	Name         string `json:"name"`
	IncludeAdult bool   `json:"include_adult"`
	Username     string `json:"username"`
}

type FavoriteMovies paginatedMovies
type WatchlistMovies paginatedMovies

type FavoriteTVShows paginatedTVShows
type WatchlistTVShows paginatedTVShows

type CreatedList struct {
	Description   string  `json:"description"`
	FavoriteCount int     `json:"favorite_count"`
	Id            int     `json:"id"`
	ISO6391       string  `json:"iso_639_1"`
	ItemCount     int     `json:"item_count"`
	ListType      string  `json:"list_type"`
	Name          string  `json:"name"`
	PosterPath    *string `json:"poster_path"`
}

type CreatedLists struct {
	pagination
	Lists []CreatedList `json:"results"`
}

type ratedMovie struct {
	movie
	Rating float64 `json:"rating"`
}

type RatedMovies struct {
	pagination
	Movies []ratedMovie `json:"results"`
}

type ratedTVShow struct {
	tv
	Rating float64 `json:"rating"`
}

type RatedTVShows struct {
	pagination
	TVShows []ratedTVShow `json:"results"`
}

type ratedTVEpisode struct {
	episode
	Rating float64 `json:"rating"`
}

type RatedTVEpisodes struct {
	pagination
	TVShows []ratedTVEpisode `json:"results"`
}

type AccountOptions struct {
	// Pass a ISO 639-1 value to display translated data for the fields that support it.
	// minLength: 2
	// pattern: ([a-z]{2})-([A-Z]{2})
	// default: en-US
	// If the provided language is wrong, it is ignored.
	Language string `url:"language,omitempty" json:"language,omitempty"`

	// Sort the results.
	// Allowed Values: created_at.asc, created_at.desc
	SortBy string `url:"sort_by,omitempty" json:"sort_by,omitempty"`

	// Specify which page to query.
	Page *int `url:"page,omitempty" json:"page,omitempty"`
}

type AccountListsOptions struct {
	// Pass a ISO 639-1 value to display translated data for the fields that support it.
	// minLength: 2
	// pattern: ([a-z]{2})-([A-Z]{2})
	// default: en-US
	// If the provided language is wrong, it is ignored.
	Language string `url:"language,omitempty" json:"language,omitempty"`

	// Specify which page to query.
	Page *int `url:"page,omitempty" json:"page,omitempty"`
}

// Get your account details.
func (ar *AccountResource) GetAccount(sessionId string) (*Account, *http.Response, error) {
	path := "/account"
	var account Account
	resp, err := ar.client.get(path, &account, WithQueryParam("session_id", sessionId))
	return &account, resp, errors.Wrap(err, "failed to get account")
}

// Get all of the lists created by an account. Will include private lists if you are the owner.
func (ar *AccountResource) GetCreatedLists(accountId int, sessionId string, opt *AccountListsOptions) (*CreatedLists, *http.Response, error) {
	path := fmt.Sprintf("/account/%d/lists", accountId)
	var lists CreatedLists
	resp, err := ar.client.get(path, &lists, WithQueryParams(opt), WithQueryParam("session_id", sessionId))
	return &lists, resp, errors.Wrap(err, "failed to get account lists")
}

// Get the list of favorite movies.
func (ar *AccountResource) GetFavoriteMovies(accountId int, sessionId string, opt *AccountOptions) (*FavoriteMovies, *http.Response, error) {
	path := fmt.Sprintf("/account/%d/favorite/movies", accountId)
	var movies FavoriteMovies
	resp, err := ar.client.get(path, &movies, WithQueryParams(opt), WithQueryParam("session_id", sessionId))
	return &movies, resp, errors.Wrap(err, "failed to get favorite movies")
}

// Get the list of favorite tv shows.
func (ar *AccountResource) GetFavoriteTVShows(accountId int, sessionId string, opt *AccountOptions) (*FavoriteTVShows, *http.Response, error) {
	path := fmt.Sprintf("/account/%d/favorite/tv", accountId)
	var tvShows FavoriteTVShows
	resp, err := ar.client.get(path, &tvShows, WithQueryParams(opt), WithQueryParam("session_id", sessionId))
	return &tvShows, resp, errors.Wrap(err, "failed to get favorite tv shows")
}

// Get the list of rated movies.
func (ar *AccountResource) GetRatedMovies(accountId int, sessionId string, opt *AccountOptions) (*RatedMovies, *http.Response, error) {
	path := fmt.Sprintf("/account/%d/rated/movies", accountId)
	var movies RatedMovies
	resp, err := ar.client.get(path, &movies, WithQueryParams(opt), WithQueryParam("session_id", sessionId))
	return &movies, resp, errors.Wrap(err, "failed to get rated movies")
}

// Get the list of rated tv shows.
func (ar *AccountResource) GetRatedTVShows(accountId int, sessionId string, opt *AccountOptions) (*RatedTVShows, *http.Response, error) {
	path := fmt.Sprintf("/account/%d/rated/tv", accountId)
	var tvShows RatedTVShows
	resp, err := ar.client.get(path, &tvShows, WithQueryParams(opt), WithQueryParam("session_id", sessionId))
	return &tvShows, resp, errors.Wrap(err, "failed to get rated tv shows")
}

// Get the list of rated tv episodes.
func (ar *AccountResource) GetRatedTVEpisodes(accountId int, sessionId string, opt *AccountOptions) (*RatedTVEpisodes, *http.Response, error) {
	path := fmt.Sprintf("/account/%d/rated/tv/episodes", accountId)
	var episodes RatedTVEpisodes
	resp, err := ar.client.get(path, &episodes, WithQueryParams(opt), WithQueryParam("session_id", sessionId))
	return &episodes, resp, errors.Wrap(err, "failed to get rated tv episodes")
}

// Get the list of rated movies.
func (ar *AccountResource) GetWatchlistMovies(accountId int, sessionId string, opt *AccountOptions) (*WatchlistMovies, *http.Response, error) {
	path := fmt.Sprintf("/account/%d/watchlist/movies", accountId)
	var movies WatchlistMovies
	resp, err := ar.client.get(path, &movies, WithQueryParams(opt), WithQueryParam("session_id", sessionId))
	return &movies, resp, errors.Wrap(err, "failed to get movies in watchlist")
}

// Get the list of rated tv shows.
func (ar *AccountResource) GetWatchlistTVShows(accountId int, sessionId string, opt *AccountOptions) (*WatchlistTVShows, *http.Response, error) {
	path := fmt.Sprintf("/account/%d/watchlist/tv", accountId)
	var tvShows WatchlistTVShows
	resp, err := ar.client.get(path, &tvShows, WithQueryParams(opt), WithQueryParam("session_id", sessionId))
	return &tvShows, resp, errors.Wrap(err, "failed to get tv shows in watchlist")
}

type statusResponse struct {
	StatusCode    int    `json:"status_code"`
	StatusMessage string `json:"status_message"`
}

type Favorite struct {
	MediaId   int    `json:"media_id"`
	MediaType string `json:"media_type"`
	Favorite  bool   `json:"favorite"`
}

type Watchlist struct {
	MediaId   int    `json:"media_id"`
	MediaType string `json:"media_type"`
	Watchlist bool   `json:"watchlist"`
}

type FavoriteResponse statusResponse
type WatchlistResponse statusResponse

// Add/remove some media to favorites.
func (ar *AccountResource) Favorite(accountId int, sessionId string, favorite Favorite) (*FavoriteResponse, *http.Response, error) {
	path := fmt.Sprintf("/account/%d/favorite", accountId)
	var favoriteResp FavoriteResponse
	resp, err := ar.client.post(path, &favoriteResp, WithBody(favorite), WithQueryParam("session_id", sessionId))
	return &favoriteResp, resp, errors.Wrap(err, "failed to mark as favorite")
}

// Add/remove some media to watchlist.
func (ar *AccountResource) Watchlist(accountId int, sessionId string, watchlist Watchlist) (*WatchlistResponse, *http.Response, error) {
	path := fmt.Sprintf("/account/%d/watchlist", accountId)
	var watchlistResp WatchlistResponse
	resp, err := ar.client.post(path, &watchlistResp, WithBody(watchlist), WithQueryParam("session_id", sessionId))
	return &watchlistResp, resp, errors.Wrap(err, "failed to mark as favorite")
}
