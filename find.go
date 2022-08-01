package tmdb

import (
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

// FindResource handles find-related requests of TMDb API.
type FindResource struct {
	client *Client
}

type Findings struct {
	MovieResults     []Movie          `json:"movie_results"`
	PersonResults    []personKnownFor `json:"person_results"`
	TvResults        []TVShow         `json:"tv_results"`
	TvEpisodeResults []Episode        `json:"tv_episode_results"`
	TvSeasonResults  []Season         `json:"tv_season_results"`
}

type FindOptions struct {
	// Pass a ISO 639-1 value to display translated data for the fields that support it.
	// minLength: 2
	// pattern: ([a-z]{2})-([A-Z]{2})
	// default: en-US
	// If the provided language is wrong, it is ignored.
	Language string `url:"language,omitempty" json:"language,omitempty"`
}

// The find method makes it easy to search for objects in our database by an external id.
// This method will search all objects (movies, TV shows and people) and return the results in a single response.
// The supported external sources for each object are as follows.
// Allowed values for external source:
//    imdb_id, freebase_mid, freebase_id, tvdb_id, tvrage_id, facebook_id, twitter_id, instagram_id
func (fr *FindResource) Find(externalId, externalSource string, opt *FindOptions) (*Findings, *http.Response, error) {
	path := fmt.Sprintf("/find/%s", externalId)
	var collection Findings
	resp, err := fr.client.get(path, &collection, WithQueryParams(opt), WithQueryParam("external_source", externalSource))
	return &collection, resp, errors.Wrap(err, "failed to find by external id")
}
