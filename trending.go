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

type KnownFor map[string]interface{}

type TrendingPerson struct {
	person
	KnownFor []KnownFor `json:"known_for"`
}

type TrendingPeople struct {
	pagination
	People []TrendingPerson `json:"results"`
}

type TrendingResult map[string]interface{}

type Trending struct {
	pagination
	Results []TrendingResult `json:"results"`
}

type TrendingMovie struct {
	movie
	MediaType string `json:"media_type"`
}

type TrendingMovies struct {
	pagination
	Movies []TrendingMovie `json:"results"`
}

type TrendingTVShow struct {
	tv
	MediaType string `json:"media_type"`
}

type TrendingTVShows struct {
	pagination
	TVShows []TrendingTVShow `json:"results"`
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

func (tr TrendingResult) ToMovie() (*TrendingMovie, error) {
	if tr.GetMediaType() != "movie" {
		return nil, errors.New(fmt.Sprintf("invalid conversion from %s to movie", tr.GetMediaType()))
	}
	return convertToMovie(tr)
}

func (tr TrendingResult) ToTVShow() (*TrendingTVShow, error) {
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

func (kf KnownFor) GetMediaType() string {
	return kf["media_type"].(string)
}

func (kf KnownFor) ToMovie() (*TrendingMovie, error) {
	if kf.GetMediaType() != "movie" {
		return nil, errors.New(fmt.Sprintf("invalid conversion from %s to movie", kf.GetMediaType()))
	}
	return convertToMovie(kf)
}

func (kf KnownFor) ToTVShow() (*TrendingTVShow, error) {
	if kf.GetMediaType() != "tv" {
		return nil, errors.New(fmt.Sprintf("invalid conversion from %s to tv", kf.GetMediaType()))
	}
	return convertToTVShow(kf)
}

func convertToMovie(obj interface{}) (*TrendingMovie, error) {
	result, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}
	var movie TrendingMovie
	err = json.Unmarshal(result, &movie)
	return &movie, err
}

func convertToTVShow(obj interface{}) (*TrendingTVShow, error) {
	result, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}
	var tvShow TrendingTVShow
	err = json.Unmarshal(result, &tvShow)
	return &tvShow, err
}
