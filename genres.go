package tmdb

import (
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

// GenresResource handles genre-related requests of TMDb API.
type GenresResource struct {
	client *Client
}

type Genre struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type genres struct {
	Genres []Genre `json:"genres"`
}

type GenresOptions struct {
	// Pass a ISO 639-1 value to display translated data for the fields that support it.
	// minLength: 2
	// pattern: ([a-z]{2})-([A-Z]{2})
	// default: en-US
	// If the provided language is wrong, it is ignored.
	Language string `url:"language,omitempty" json:"language,omitempty"`
}

// Get the list of official genres for movies.
func (gr *GenresResource) GetMovieGenres(opt *GenresOptions) ([]Genre, *http.Response, error) {
	return gr.getGenres("movie", opt)
}

// Get the list of official genres for TV shows.
func (gr *GenresResource) GetTVGenres(opt *GenresOptions) ([]Genre, *http.Response, error) {
	return gr.getGenres("tv", opt)
}

func (gr *GenresResource) getGenres(listType string, opt *GenresOptions) ([]Genre, *http.Response, error) {
	path := fmt.Sprintf("/genre/%s/list", listType)
	var genres genres
	resp, err := gr.client.get(path, &genres, WithQueryParams(opt))
	return genres.Genres, resp, errors.Wrap(err, fmt.Sprintf("failed to get %s genres", listType))
}
