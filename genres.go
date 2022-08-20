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

type GenresResponse struct {
	Genres []Genre `json:"genres"`
}

type GenresOptions languageOptions

// Get the list of official genres for movies.
func (gr *GenresResource) GetMovieGenres(opt *GenresOptions) (*GenresResponse, *http.Response, error) {
	return gr.getGenres("movie", opt)
}

// Get the list of official genres for TV shows.
func (gr *GenresResource) GetTVGenres(opt *GenresOptions) (*GenresResponse, *http.Response, error) {
	return gr.getGenres("tv", opt)
}

func (gr *GenresResource) getGenres(listType string, opt *GenresOptions) (*GenresResponse, *http.Response, error) {
	path := fmt.Sprintf("/genre/%s/list", listType)
	var response GenresResponse
	resp, err := gr.client.get(path, &response, WithQueryParams(opt))
	return &response, resp, errors.Wrap(err, fmt.Sprintf("failed to get %s genres", listType))
}
