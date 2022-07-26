package tmdb

import (
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

// KeywordsResource handles keyword-related requests of TMDb API.
type KeywordsResource struct {
	client *Client
}

// Keyword represents a keyword in TMDb.
type Keyword struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// GetKeyword retrieves a specific keyword.
func (kr *KeywordsResource) GetKeyword(id int) (*Keyword, *http.Response, error) {
	path := fmt.Sprintf("/keyword/%d", id)
	var keyword Keyword
	resp, err := kr.client.get(path, &keyword)
	return &keyword, resp, errors.Wrap(err, "failed to get keyword")
}

// KeywordMovies represents keyword movies in TMDb.
type KeywordMovies struct {
	ID           int           `json:"id"`
	Page         int           `json:"page"`
	Movies       []MovieResult `json:"results"`
	TotalPages   int           `json:"total_pages"`
	TotalResults int           `json:"total_results"`
}

// KeywordMoviesOptions represents the available options for the request.
type KeywordMoviesOptions struct {
	// Pass a ISO 639-1 value to display translated data for the fields that support it.
	// minLength: 2
	// pattern: ([a-z]{2})-([A-Z]{2})
	// default: en-US
	// If the provided language is wrong, it is ignored.
	Language string `url:"language,omitempty" json:"language,omitempty"`

	// Choose whether to include adult (pornography) content in the results.
	// default: false
	IncludeAdult bool `url:"include_adult,omitempty" json:"include_adult,omitempty"`

	// Specify which page to query.
	Page *int `url:"page,omitempty" json:"page,omitempty"`
}

// GetKeywordMovies retrieves the movies that belong to a keyword.
// It is highly recommend using movie discover instead of this method as it is much more flexible.
func (kr *KeywordsResource) GetKeywordMovies(id int, opt *KeywordMoviesOptions) (*KeywordMovies, *http.Response, error) {
	path := fmt.Sprintf("/keyword/%d/movies", id)
	var keyword KeywordMovies
	resp, err := kr.client.get(path, &keyword, WithQueryParams(opt))
	return &keyword, resp, errors.Wrap(err, "failed to get keyword")
}
