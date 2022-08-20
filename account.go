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

type Gravatar struct {
	Hash string `json:"hash"`
}

type Avatar struct {
	Gravatar Gravatar `json:"gravatar"`
}

type Account struct {
	Avatar       Avatar `json:"avatar"`
	Id           int    `json:"id"`
	ISO6391      string `json:"iso_639_1"`
	ISO31661     string `json:"iso_3166_1"`
	Name         string `json:"name"`
	IncludeAdult bool   `json:"include_adult"`
	Username     string `json:"username"`
}

// Get your account details.
func (ar *AccountResource) GetAccount(sessionId string) (*Account, *http.Response, error) {
	path := "/account"
	var account Account
	resp, err := ar.client.get(path, &account, WithQueryParam("session_id", sessionId))
	return &account, resp, errors.Wrap(err, "failed to get account")
}

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

type AccountListsOptions languagePageOptions

// Get all of the lists created by an account. Will include private lists if you are the owner.
func (ar *AccountResource) GetCreatedLists(accountId int, sessionId string, opt *AccountListsOptions) (*CreatedLists, *http.Response, error) {
	path := fmt.Sprintf("/account/%d/lists", accountId)
	var lists CreatedLists
	resp, err := ar.client.get(path, &lists, WithQueryParams(opt), WithQueryParam("session_id", sessionId))
	return &lists, resp, errors.Wrap(err, "failed to get account lists")
}

type FavoriteMovies paginatedMovies

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

// Get the list of favorite movies.
func (ar *AccountResource) GetFavoriteMovies(accountId int, sessionId string, opt *AccountOptions) (*FavoriteMovies, *http.Response, error) {
	path := fmt.Sprintf("/account/%d/favorite/movies", accountId)
	var movies FavoriteMovies
	resp, err := ar.client.get(path, &movies, WithQueryParams(opt), WithQueryParam("session_id", sessionId))
	return &movies, resp, errors.Wrap(err, "failed to get favorite movies")
}

type FavoriteTVShows paginatedTVShows

// Get the list of favorite tv shows.
func (ar *AccountResource) GetFavoriteTVShows(accountId int, sessionId string, opt *AccountOptions) (*FavoriteTVShows, *http.Response, error) {
	path := fmt.Sprintf("/account/%d/favorite/tv", accountId)
	var tvShows FavoriteTVShows
	resp, err := ar.client.get(path, &tvShows, WithQueryParams(opt), WithQueryParam("session_id", sessionId))
	return &tvShows, resp, errors.Wrap(err, "failed to get favorite tv shows")
}

type RatedMovie struct {
	Adult            bool    `json:"adult"`
	BackdropPath     *string `json:"backdrop_path"`
	GenreIds         []int   `json:"genre_ids"`
	Id               int     `json:"id"`
	OriginalLanguage string  `json:"original_language"`
	OriginalTitle    string  `json:"original_title"`
	Overview         string  `json:"overview"`
	Popularity       float64 `json:"popularity"`
	PosterPath       *string `json:"poster_path"`
	Rating           float64 `json:"rating"`
	ReleaseDate      string  `json:"release_date"`
	Title            string  `json:"title"`
	Video            bool    `json:"video"`
	VoteAverage      float64 `json:"vote_average"`
	VoteCount        int     `json:"vote_count"`
}

type RatedMovies struct {
	pagination
	Movies []RatedMovie `json:"results"`
}

// Get the list of rated movies.
func (ar *AccountResource) GetRatedMovies(accountId int, sessionId string, opt *AccountOptions) (*RatedMovies, *http.Response, error) {
	path := fmt.Sprintf("/account/%d/rated/movies", accountId)
	var movies RatedMovies
	resp, err := ar.client.get(path, &movies, WithQueryParams(opt), WithQueryParam("session_id", sessionId))
	return &movies, resp, errors.Wrap(err, "failed to get rated movies")
}

type RatedTVShow struct {
	Adult            bool     `json:"adult"`
	BackdropPath     *string  `json:"backdrop_path"`
	FirstAirDate     string   `json:"first_air_date"`
	GenreIds         []int    `json:"genre_ids"`
	Id               int      `json:"id"`
	Name             string   `json:"name"`
	OriginalLanguage string   `json:"original_language"`
	OriginalName     string   `json:"original_name"`
	OriginCountry    []string `json:"origin_country"`
	Overview         string   `json:"overview"`
	Popularity       float64  `json:"popularity"`
	PosterPath       *string  `json:"poster_path"`
	Rating           float64  `json:"rating"`
	VoteAverage      float64  `json:"vote_average"`
	VoteCount        int      `json:"vote_count"`
}

type RatedTVShows struct {
	pagination
	TVShows []RatedTVShow `json:"results"`
}

// Get the list of rated tv shows.
func (ar *AccountResource) GetRatedTVShows(accountId int, sessionId string, opt *AccountOptions) (*RatedTVShows, *http.Response, error) {
	path := fmt.Sprintf("/account/%d/rated/tv", accountId)
	var tvShows RatedTVShows
	resp, err := ar.client.get(path, &tvShows, WithQueryParams(opt), WithQueryParam("session_id", sessionId))
	return &tvShows, resp, errors.Wrap(err, "failed to get rated tv shows")
}

type RatedTVEpisode struct {
	AirDate        string  `json:"air_date"`
	EpisodeNumber  int     `json:"episode_number"`
	Id             int     `json:"id"`
	Name           string  `json:"name"`
	Overview       string  `json:"overview"`
	ProductionCode string  `json:"production_code"`
	Rating         float64 `json:"rating"`
	Runtime        int     `json:"runtime"`
	SeasonNumber   int     `json:"season_number"`
	ShowId         int     `json:"show_id"`
	StillPath      *string `json:"still_path"`
	VoteAverage    float64 `json:"vote_average"`
	VoteCount      int     `json:"vote_count"`
}

type RatedTVEpisodes struct {
	pagination
	TVShows []RatedTVEpisode `json:"results"`
}

// Get the list of rated tv episodes.
func (ar *AccountResource) GetRatedTVEpisodes(accountId int, sessionId string, opt *AccountOptions) (*RatedTVEpisodes, *http.Response, error) {
	path := fmt.Sprintf("/account/%d/rated/tv/episodes", accountId)
	var episodes RatedTVEpisodes
	resp, err := ar.client.get(path, &episodes, WithQueryParams(opt), WithQueryParam("session_id", sessionId))
	return &episodes, resp, errors.Wrap(err, "failed to get rated tv episodes")
}

type WatchlistMovies paginatedMovies

// Get the list of rated movies.
func (ar *AccountResource) GetWatchlistMovies(accountId int, sessionId string, opt *AccountOptions) (*WatchlistMovies, *http.Response, error) {
	path := fmt.Sprintf("/account/%d/watchlist/movies", accountId)
	var movies WatchlistMovies
	resp, err := ar.client.get(path, &movies, WithQueryParams(opt), WithQueryParam("session_id", sessionId))
	return &movies, resp, errors.Wrap(err, "failed to get movies in watchlist")
}

type WatchlistTVShows paginatedTVShows

// Get the list of rated tv shows.
func (ar *AccountResource) GetWatchlistTVShows(accountId int, sessionId string, opt *AccountOptions) (*WatchlistTVShows, *http.Response, error) {
	path := fmt.Sprintf("/account/%d/watchlist/tv", accountId)
	var tvShows WatchlistTVShows
	resp, err := ar.client.get(path, &tvShows, WithQueryParams(opt), WithQueryParam("session_id", sessionId))
	return &tvShows, resp, errors.Wrap(err, "failed to get tv shows in watchlist")
}

type Favorite struct {
	MediaId   int    `json:"media_id"`
	MediaType string `json:"media_type"`
	Favorite  bool   `json:"favorite"`
}

type FavoriteResponse statusResponse

// Add/remove some media to favorites.
func (ar *AccountResource) Favorite(accountId int, sessionId string, favorite Favorite) (*FavoriteResponse, *http.Response, error) {
	path := fmt.Sprintf("/account/%d/favorite", accountId)
	var favoriteResp FavoriteResponse
	resp, err := ar.client.post(path, &favoriteResp, WithBody(favorite), WithQueryParam("session_id", sessionId))
	return &favoriteResp, resp, errors.Wrap(err, "failed to mark as favorite")
}

type Watchlist struct {
	MediaId   int    `json:"media_id"`
	MediaType string `json:"media_type"`
	Watchlist bool   `json:"watchlist"`
}

type WatchlistResponse statusResponse

// Add/remove some media to watchlist.
func (ar *AccountResource) Watchlist(accountId int, sessionId string, watchlist Watchlist) (*WatchlistResponse, *http.Response, error) {
	path := fmt.Sprintf("/account/%d/watchlist", accountId)
	var watchlistResp WatchlistResponse
	resp, err := ar.client.post(path, &watchlistResp, WithBody(watchlist), WithQueryParam("session_id", sessionId))
	return &watchlistResp, resp, errors.Wrap(err, "failed to mark as favorite")
}
