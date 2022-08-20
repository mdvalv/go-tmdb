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

type EpisodeFinding struct {
	AirDate        string  `json:"air_date"`
	EpisodeNumber  int     `json:"episode_number"`
	Id             int     `json:"id"`
	MediaType      string  `json:"media_type"`
	Name           string  `json:"name"`
	Overview       string  `json:"overview"`
	ProductionCode string  `json:"production_code"`
	Runtime        int     `json:"runtime"`
	SeasonNumber   int     `json:"season_number"`
	ShowId         int     `json:"show_id"`
	StillPath      *string `json:"still_path"`
	VoteAverage    float64 `json:"vote_average"`
	VoteCount      int     `json:"vote_count"`
}

type SeasonFinding struct {
	AirDate      string  `json:"air_date"`
	EpisodeCount int     `json:"episode_count"`
	Id           int     `json:"id"`
	MediaType    string  `json:"media_type"`
	Name         string  `json:"name"`
	Overview     string  `json:"overview"`
	PosterPath   *string `json:"poster_path"`
	SeasonNumber int     `json:"season_number"`
}

type Findings struct {
	MovieResults     []Movie          `json:"movie_results"`
	PersonResults    []Person         `json:"person_results"`
	TvResults        []TVShow         `json:"tv_results"`
	TvEpisodeResults []EpisodeFinding `json:"tv_episode_results"`
	TvSeasonResults  []SeasonFinding  `json:"tv_season_results"`
}

type FindOptions languageOptions

// The find method makes it easy to search for objects by an external id.
// This method will search all objects (movies, TV shows and people) and return the results in a single response.
// Allowed values for external source:
//    imdb_id, freebase_mid, freebase_id, tvdb_id, tvrage_id, facebook_id, twitter_id, instagram_id
func (fr *FindResource) Find(externalId, externalSource string, opt *FindOptions) (*Findings, *http.Response, error) {
	path := fmt.Sprintf("/find/%s", externalId)
	var collection Findings
	resp, err := fr.client.get(path, &collection, WithQueryParams(opt), WithQueryParam("external_source", externalSource))
	return &collection, resp, errors.Wrap(err, "failed to find by external id")
}
