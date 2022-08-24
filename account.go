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

// Gravatar represents a gravatar object.
type Gravatar struct {
	Hash string `json:"hash"`
}

// Avatar represents an avatar object.
type Avatar struct {
	Gravatar Gravatar `json:"gravatar"`
}

// Account represents a TMDb account.
type Account struct {
	Avatar       Avatar `json:"avatar"`
	ID           int    `json:"id"`
	ISO6391      string `json:"iso_639_1"`
	ISO31661     string `json:"iso_3166_1"`
	Name         string `json:"name"`
	IncludeAdult bool   `json:"include_adult"`
	Username     string `json:"username"`
}

// GetAccount retrieves account details from TMDb.
func (ar *AccountResource) GetAccount(sessionID string) (*Account, *http.Response, error) {
	path := "/account"
	var account Account
	resp, err := ar.client.get(path, &account, WithSessionID(sessionID))
	return &account, resp, errors.Wrap(err, "failed to get account")
}

// CreatedList represents a created list in TMDb.
type CreatedList struct {
	Description   string  `json:"description"`
	FavoriteCount int     `json:"favorite_count"`
	ID            int     `json:"id"`
	ISO6391       string  `json:"iso_639_1"`
	ItemCount     int     `json:"item_count"`
	ListType      string  `json:"list_type"`
	Name          string  `json:"name"`
	PosterPath    *string `json:"poster_path"`
}

// CreatedLists represents the created lists in TMDb.
type CreatedLists struct {
	pagination
	Lists []CreatedList `json:"results"`
}

// AccountListsOptions represents the available options for the request.
type AccountListsOptions languagePageOptions

// GetCreatedLists retrieves all of the lists created by an account.
// Will include private lists if the requester is the owner.
func (ar *AccountResource) GetCreatedLists(accountID int, sessionID string, opt *AccountListsOptions) (*CreatedLists, *http.Response, error) {
	path := fmt.Sprintf("/account/%d/lists", accountID)
	var lists CreatedLists
	resp, err := ar.client.get(path, &lists, WithQueryParams(opt), WithSessionID(sessionID))
	return &lists, resp, errors.Wrap(err, "failed to get account lists")
}

// FavoriteMovies represents the favorite movies in TMDb.
type FavoriteMovies paginatedMovies

// AccountOptions represents the available options for the request.
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

// GetFavoriteMovies retrieves the list of favorite movies.
func (ar *AccountResource) GetFavoriteMovies(accountID int, sessionID string, opt *AccountOptions) (*FavoriteMovies, *http.Response, error) {
	path := fmt.Sprintf("/account/%d/favorite/movies", accountID)
	var movies FavoriteMovies
	resp, err := ar.client.get(path, &movies, WithQueryParams(opt), WithSessionID(sessionID))
	return &movies, resp, errors.Wrap(err, "failed to get favorite movies")
}

// FavoriteTVShows represents the favorite tv shows in TMDb.
type FavoriteTVShows paginatedTVShows

// GetFavoriteTVShows retrieves the list of favorite tv shows.
func (ar *AccountResource) GetFavoriteTVShows(accountID int, sessionID string, opt *AccountOptions) (*FavoriteTVShows, *http.Response, error) {
	path := fmt.Sprintf("/account/%d/favorite/tv", accountID)
	var tvShows FavoriteTVShows
	resp, err := ar.client.get(path, &tvShows, WithQueryParams(opt), WithSessionID(sessionID))
	return &tvShows, resp, errors.Wrap(err, "failed to get favorite tv shows")
}

// RatedMovie represents a rated movie in TMDb.
type RatedMovie struct {
	Adult            bool    `json:"adult"`
	BackdropPath     *string `json:"backdrop_path"`
	GenreIDs         []int   `json:"genre_ids"`
	ID               int     `json:"id"`
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

// RatedMovies represents rated movies in TMDb.
type RatedMovies struct {
	pagination
	Movies []RatedMovie `json:"results"`
}

// GetRatedMovies retrieves the list of rated movies.
func (ar *AccountResource) GetRatedMovies(accountID int, sessionID string, opt *AccountOptions) (*RatedMovies, *http.Response, error) {
	path := fmt.Sprintf("/account/%d/rated/movies", accountID)
	var movies RatedMovies
	resp, err := ar.client.get(path, &movies, WithQueryParams(opt), WithSessionID(sessionID))
	return &movies, resp, errors.Wrap(err, "failed to get rated movies")
}

// RatedTVShow represents a rated tv show in TMDb.
type RatedTVShow struct {
	Adult            bool     `json:"adult"`
	BackdropPath     *string  `json:"backdrop_path"`
	FirstAirDate     string   `json:"first_air_date"`
	GenreIDs         []int    `json:"genre_ids"`
	ID               int      `json:"id"`
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

// RatedTVShows represents rated tv shows in TMDb.
type RatedTVShows struct {
	pagination
	TVShows []RatedTVShow `json:"results"`
}

// GetRatedTVShows retrieves the list of rated tv shows.
func (ar *AccountResource) GetRatedTVShows(accountID int, sessionID string, opt *AccountOptions) (*RatedTVShows, *http.Response, error) {
	path := fmt.Sprintf("/account/%d/rated/tv", accountID)
	var tvShows RatedTVShows
	resp, err := ar.client.get(path, &tvShows, WithQueryParams(opt), WithSessionID(sessionID))
	return &tvShows, resp, errors.Wrap(err, "failed to get rated tv shows")
}

// RatedTVEpisode represents a rated tv episode in TMDb.
type RatedTVEpisode struct {
	AirDate        string  `json:"air_date"`
	EpisodeNumber  int     `json:"episode_number"`
	ID             int     `json:"id"`
	Name           string  `json:"name"`
	Overview       string  `json:"overview"`
	ProductionCode string  `json:"production_code"`
	Rating         float64 `json:"rating"`
	Runtime        int     `json:"runtime"`
	SeasonNumber   int     `json:"season_number"`
	ShowID         int     `json:"show_id"`
	StillPath      *string `json:"still_path"`
	VoteAverage    float64 `json:"vote_average"`
	VoteCount      int     `json:"vote_count"`
}

// RatedTVEpisodes represents rated tv episodes in TMDb.
type RatedTVEpisodes struct {
	pagination
	TVShows []RatedTVEpisode `json:"results"`
}

// GetRatedTVEpisodes retrieves the list of rated tv episodes.
func (ar *AccountResource) GetRatedTVEpisodes(accountID int, sessionID string, opt *AccountOptions) (*RatedTVEpisodes, *http.Response, error) {
	path := fmt.Sprintf("/account/%d/rated/tv/episodes", accountID)
	var episodes RatedTVEpisodes
	resp, err := ar.client.get(path, &episodes, WithQueryParams(opt), WithSessionID(sessionID))
	return &episodes, resp, errors.Wrap(err, "failed to get rated tv episodes")
}

// WatchlistMovies represents movies added to the watchlist in TMDb.
type WatchlistMovies paginatedMovies

// GetWatchlistMovies retrieves the list of rated movies.
func (ar *AccountResource) GetWatchlistMovies(accountID int, sessionID string, opt *AccountOptions) (*WatchlistMovies, *http.Response, error) {
	path := fmt.Sprintf("/account/%d/watchlist/movies", accountID)
	var movies WatchlistMovies
	resp, err := ar.client.get(path, &movies, WithQueryParams(opt), WithSessionID(sessionID))
	return &movies, resp, errors.Wrap(err, "failed to get movies in watchlist")
}

// WatchlistTVShows represents tv shows added to the watchlist in TMDb.
type WatchlistTVShows paginatedTVShows

// GetWatchlistTVShows retrieves the list of rated tv shows.
func (ar *AccountResource) GetWatchlistTVShows(accountID int, sessionID string, opt *AccountOptions) (*WatchlistTVShows, *http.Response, error) {
	path := fmt.Sprintf("/account/%d/watchlist/tv", accountID)
	var tvShows WatchlistTVShows
	resp, err := ar.client.get(path, &tvShows, WithQueryParams(opt), WithSessionID(sessionID))
	return &tvShows, resp, errors.Wrap(err, "failed to get tv shows in watchlist")
}

// Favorite represents a favorite object in TMDb.
type Favorite struct {
	MediaID   int    `json:"media_id"`
	MediaType string `json:"media_type"`
	Favorite  bool   `json:"favorite"`
}

// FavoriteResponse represents the favorite response.
type FavoriteResponse statusResponse

// Favorite adds/removes some media to/from favorites.
func (ar *AccountResource) Favorite(accountID int, sessionID string, favorite Favorite) (*FavoriteResponse, *http.Response, error) {
	path := fmt.Sprintf("/account/%d/favorite", accountID)
	var favoriteResp FavoriteResponse
	resp, err := ar.client.post(path, &favoriteResp, WithBody(favorite), WithSessionID(sessionID))
	return &favoriteResp, resp, errors.Wrap(err, "failed to mark as favorite")
}

// Watchlist represents a watchlist object in TMDb.
type Watchlist struct {
	MediaID   int    `json:"media_id"`
	MediaType string `json:"media_type"`
	Watchlist bool   `json:"watchlist"`
}

// WatchlistResponse represents the watchlist response.
type WatchlistResponse statusResponse

// Watchlist adds/removes some media to/from watchlist.
func (ar *AccountResource) Watchlist(accountID int, sessionID string, watchlist Watchlist) (*WatchlistResponse, *http.Response, error) {
	path := fmt.Sprintf("/account/%d/watchlist", accountID)
	var watchlistResp WatchlistResponse
	resp, err := ar.client.post(path, &watchlistResp, WithBody(watchlist), WithSessionID(sessionID))
	return &watchlistResp, resp, errors.Wrap(err, "failed to mark as favorite")
}
