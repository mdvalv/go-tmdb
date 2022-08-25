package tmdb

import (
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

// GuestSessionResource handles guest session-related requests of TMDb API.
type GuestSessionResource struct {
	client *Client
}

// GuestSessionOptions represents the available options for the request.
type GuestSessionOptions struct {
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

// GetRatedMovies retrieves the list of rated movies.
func (ar *GuestSessionResource) GetRatedMovies(sessionID string, opt *GuestSessionOptions) (*RatedMovies, *http.Response, error) {
	path := fmt.Sprintf("/guest_session/%s/rated/movies", sessionID)
	var movies RatedMovies
	resp, err := ar.client.get(path, &movies, WithQueryParams(opt))
	return &movies, resp, errors.Wrap(err, "failed to get rated movies")
}

// GetRatedTVShows retrieves the list of rated tv shows.
func (ar *GuestSessionResource) GetRatedTVShows(sessionID string, opt *GuestSessionOptions) (*RatedTVShows, *http.Response, error) {
	path := fmt.Sprintf("/guest_session/%s/rated/tv", sessionID)
	var tvShows RatedTVShows
	resp, err := ar.client.get(path, &tvShows, WithQueryParams(opt))
	return &tvShows, resp, errors.Wrap(err, "failed to get rated tv shows")
}

// GetRatedTVEpisodes retrieves the list of rated tv episodes.
func (ar *GuestSessionResource) GetRatedTVEpisodes(sessionID string, opt *GuestSessionOptions) (*RatedTVEpisodes, *http.Response, error) {
	path := fmt.Sprintf("/guest_session/%s/rated/tv/episodes", sessionID)
	var episodes RatedTVEpisodes
	resp, err := ar.client.get(path, &episodes, WithQueryParams(opt))
	return &episodes, resp, errors.Wrap(err, "failed to get rated tv episodes")
}
