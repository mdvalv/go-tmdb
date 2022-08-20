package tmdb

import (
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

// TrendingResource handles trending-related requests of TMDb API.
type TrendingResource struct {
	client *Client
}

type TrendingMovies struct {
	pagination
	Movies []Movie `json:"results"`
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

type TrendingTVShows struct {
	pagination
	TVShows []TVShow `json:"results"`
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

type TrendingPerson Person

type TrendingPeople struct {
	pagination
	People []TrendingPerson `json:"results"`
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

type TrendingResult map[string]interface{}

func (tr TrendingResult) GetMediaType() string {
	return tr["media_type"].(string)
}

func (tr TrendingResult) ToMovie() (*Movie, error) {
	return convertToMovie(tr)
}

func (tr TrendingResult) ToTVShow() (*TVShow, error) {
	return convertToTVShow(tr)
}

func (tr TrendingResult) ToPerson() (*TrendingPerson, error) {
	var person TrendingPerson
	err := convert("person", tr, &person)
	return &person, errors.Wrap(err, "failed to convert object to person")
}

type Trending struct {
	pagination
	Results []TrendingResult `json:"results"`
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
