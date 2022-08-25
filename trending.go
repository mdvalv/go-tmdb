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

// TrendingMovies represents trending movies in TMDb.
type TrendingMovies struct {
	pagination
	Movies []Movie `json:"results"`
}

// GetTrendingMovies retrieves the daily or weekly trending movies.
// The daily trending list tracks items over the period of a day while items have a 24 hour half life.
// The weekly list tracks items over a 7 day period, with a 7 day half life.
// Allowed timeWindow: day, week
func (tr *TrendingResource) GetTrendingMovies(timeWindow string) (*TrendingMovies, *http.Response, error) {
	path := fmt.Sprintf("/trending/movie/%s", timeWindow)
	var trending TrendingMovies
	resp, err := tr.client.get(path, &trending)
	return &trending, resp, errors.Wrap(err, "failed to get trending movies")
}

// TrendingTVShows represents trending tv shows in TMDb.
type TrendingTVShows struct {
	pagination
	TVShows []TVShow `json:"results"`
}

// GetTrendingTVShows retrieves the daily or weekly trending tv shows.
// The daily trending list tracks items over the period of a day while items have a 24 hour half life.
// The weekly list tracks items over a 7 day period, with a 7 day half life.
// Allowed timeWindow: day, week
func (tr *TrendingResource) GetTrendingTVShows(timeWindow string) (*TrendingTVShows, *http.Response, error) {
	path := fmt.Sprintf("/trending/tv/%s", timeWindow)
	var trending TrendingTVShows
	resp, err := tr.client.get(path, &trending)
	return &trending, resp, errors.Wrap(err, "failed to get trending tv")
}

// TrendingPerson represents a trending person in TMDb.
type TrendingPerson Person

// TrendingPeople represents trending people in TMDb.
type TrendingPeople struct {
	pagination
	People []TrendingPerson `json:"results"`
}

// GetTrendingPeople retrieves the daily or weekly trending people.
// The daily trending list tracks items over the period of a day while items have a 24 hour half life.
// The weekly list tracks items over a 7 day period, with a 7 day half life.
// Allowed timeWindow: day, week
func (tr *TrendingResource) GetTrendingPeople(timeWindow string) (*TrendingPeople, *http.Response, error) {
	path := fmt.Sprintf("/trending/person/%s", timeWindow)
	var trending TrendingPeople
	resp, err := tr.client.get(path, &trending)
	return &trending, resp, errors.Wrap(err, "failed to get trending people")
}

// TrendingResult represents a trending result in TMDb.
type TrendingResult map[string]interface{}

// GetMediaType retrieves the media type from a trending result.
func (tr TrendingResult) GetMediaType() string {
	return tr["media_type"].(string)
}

// ToMovie converts the data to a movie.
func (tr TrendingResult) ToMovie() (*Movie, error) {
	return convertToMovie(tr)
}

// ToTVShow converts the data to a tv show.
func (tr TrendingResult) ToTVShow() (*TVShow, error) {
	return convertToTVShow(tr)
}

// ToPerson converts the data to a person.
func (tr TrendingResult) ToPerson() (*TrendingPerson, error) {
	var person TrendingPerson
	err := convert("person", tr, &person)
	return &person, errors.Wrap(err, "failed to convert object to person")
}

// Trending represents trending information in TMDb.
type Trending struct {
	pagination
	Results []TrendingResult `json:"results"`
}

// GetTrending retrieves the daily or weekly trending items.
// The daily trending list tracks items over the period of a day while items have a 24 hour half life.
// The weekly list tracks items over a 7 day period, with a 7 day half life.
// Allowed timeWindow: day, week
func (tr *TrendingResource) GetTrending(timeWindow string) (*Trending, *http.Response, error) {
	path := fmt.Sprintf("/trending/all/%s", timeWindow)
	var trending Trending
	resp, err := tr.client.get(path, &trending)
	return &trending, resp, errors.Wrap(err, "failed to get trending information")
}
