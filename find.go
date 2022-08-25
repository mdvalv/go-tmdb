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

// EpisodeFinding represents an episode finding in TMDb.
type EpisodeFinding struct {
	AirDate        string  `json:"air_date"`
	EpisodeNumber  int     `json:"episode_number"`
	ID             int     `json:"id"`
	MediaType      string  `json:"media_type"`
	Name           string  `json:"name"`
	Overview       string  `json:"overview"`
	ProductionCode string  `json:"production_code"`
	Runtime        int     `json:"runtime"`
	SeasonNumber   int     `json:"season_number"`
	ShowID         int     `json:"show_id"`
	StillPath      *string `json:"still_path"`
	VoteAverage    float64 `json:"vote_average"`
	VoteCount      int     `json:"vote_count"`
}

// SeasonFinding represents a season finding in TMDb.
type SeasonFinding struct {
	AirDate      string  `json:"air_date"`
	EpisodeCount int     `json:"episode_count"`
	ID           int     `json:"id"`
	MediaType    string  `json:"media_type"`
	Name         string  `json:"name"`
	Overview     string  `json:"overview"`
	PosterPath   *string `json:"poster_path"`
	SeasonNumber int     `json:"season_number"`
}

// Findings represents findings in TMDb.
type Findings struct {
	Movies     []Movie          `json:"movie_results"`
	People     []Person         `json:"person_results"`
	TVShows    []TVShow         `json:"tv_results"`
	TVEpisodes []EpisodeFinding `json:"tv_episode_results"`
	TVSeasons  []SeasonFinding  `json:"tv_season_results"`
}

// FindOptions represents the available options for the request.
type FindOptions languageOptions

// Find makes it easy to search for objects by an external id.
// This method will search all objects (movies, TV shows and people) and return the results in a single response.
// Allowed values for external source:
//    imdb_id, freebase_mid, freebase_id, tvdb_id, tvrage_id, facebook_id, twitter_id, instagram_id
func (fr *FindResource) Find(externalID, externalSource string, opt *FindOptions) (*Findings, *http.Response, error) {
	path := fmt.Sprintf("/find/%s", externalID)
	var collection Findings
	resp, err := fr.client.get(path, &collection, WithQueryParams(opt), WithQueryParam("external_source", externalSource))
	return &collection, resp, errors.Wrap(err, "failed to find by external id")
}
