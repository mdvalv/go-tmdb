package tmdb

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

// TrendingResource handles trending-related requests of TMDb API.
type TrendingResource struct {
	client *Client
}

type TrendingOptions struct {
	// Specify which page to query.
	Page *int `url:"page,omitempty" json:"page,omitempty"`
}

type TrendingPerson personKnownFor

type TrendingPeople struct {
	pagination
	People []TrendingPerson `json:"results"`
}

type TrendingResult map[string]interface{}

type Trending struct {
	pagination
	Results []TrendingResult `json:"results"`
}

type TrendingMovies struct {
	pagination
	Movies []Movie `json:"results"`
}

type TrendingTVShows struct {
	pagination
	TVShows []TVShow `json:"results"`
}

// Get the daily or weekly trending items.
// The daily trending list tracks items over the period of a day while items have a 24 hour half life.
// The weekly list tracks items over a 7 day period, with a 7 day half life.
// Allowed timeWindow: day, week
func (tr *TrendingResource) GetTrending(timeWindow string) (*Trending, *http.Response, error) {
	path := fmt.Sprintf("/trending/all/%s", timeWindow)
	var trending Trending
	resp, err := tr.client.get(path, &trending)
	return &trending, resp, errors.Wrap(err, "failed to get trending information")
}

// Get the daily or weekly trending movies.
// The daily trending list tracks items over the period of a day while items have a 24 hour half life.
// The weekly list tracks items over a 7 day period, with a 7 day half life.
// Allowed timeWindow: day, week
func (tr *TrendingResource) GetTrendingMovies(timeWindow string) (*TrendingMovies, *http.Response, error) {
	path := fmt.Sprintf("/trending/movie/%s", timeWindow)
	var trending TrendingMovies
	resp, err := tr.client.get(path, &trending)
	return &trending, resp, errors.Wrap(err, "failed to get trending movies")
}

// Get the daily or weekly trending tv shows.
// The daily trending list tracks items over the period of a day while items have a 24 hour half life.
// The weekly list tracks items over a 7 day period, with a 7 day half life.
// Allowed timeWindow: day, week
func (tr *TrendingResource) GetTrendingTVShows(timeWindow string) (*TrendingTVShows, *http.Response, error) {
	path := fmt.Sprintf("/trending/tv/%s", timeWindow)
	var trending TrendingTVShows
	resp, err := tr.client.get(path, &trending)
	return &trending, resp, errors.Wrap(err, "failed to get trending tv")
}

// Get the daily or weekly trending people.
// The daily trending list tracks items over the period of a day while items have a 24 hour half life.
// The weekly list tracks items over a 7 day period, with a 7 day half life.
// Allowed timeWindow: day, week
func (tr *TrendingResource) GetTrendingPeople(timeWindow string) (*TrendingPeople, *http.Response, error) {
	path := fmt.Sprintf("/trending/person/%s", timeWindow)
	var trending TrendingPeople
	resp, err := tr.client.get(path, &trending)
	return &trending, resp, errors.Wrap(err, "failed to get trending people")
}

func (tr TrendingResult) GetMediaType() string {
	return tr["media_type"].(string)
}

func (tr TrendingResult) ToMovie() (*Movie, error) {
	if tr.GetMediaType() != "movie" {
		return nil, errors.New(fmt.Sprintf("invalid conversion from %s to movie", tr.GetMediaType()))
	}
	return convertToMovie(tr)
}

func (tr TrendingResult) ToTVShow() (*TVShow, error) {
	if tr.GetMediaType() != "tv" {
		return nil, errors.New(fmt.Sprintf("invalid conversion from %s to tv", tr.GetMediaType()))
	}
	return convertToTVShow(tr)
}

func (tr TrendingResult) ToPerson() (*TrendingPerson, error) {
	if tr.GetMediaType() != "person" {
		return nil, errors.New(fmt.Sprintf("invalid conversion from %s to person", tr.GetMediaType()))
	}
	result, err := json.Marshal(tr)
	if err != nil {
		return nil, err
	}
	var movie TrendingPerson
	err = json.Unmarshal(result, &movie)
	return &movie, err
}
