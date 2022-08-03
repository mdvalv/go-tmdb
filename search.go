package tmdb

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

// SearchResource handles search-related requests of TMDb API.
type SearchResource struct {
	client *Client
}

type SearchCompanies struct {
	pagination
	Results []company `json:"results"`
}

type SearchCompaniesOptions struct {
	// Specify which page to query.
	Page *int `url:"page,omitempty" json:"page,omitempty"`
}

// Search for companies.
func (sr *SearchResource) Companies(query string, opt *SearchCompaniesOptions) (*SearchCompanies, *http.Response, error) {
	path := "/search/company"
	var companies SearchCompanies
	resp, err := sr.client.get(path, &companies, WithQueryParam("query", query), WithQueryParams(opt))
	return &companies, resp, errors.Wrap(err, "failed to search for companies")
}

type SearchCollection struct {
	Adult            bool    `json:"adult"`
	BackdropPath     *string `json:"backdrop_path"`
	Id               int     `json:"id"`
	Name             string  `json:"name"`
	OriginalLanguage string  `json:"original_language"`
	OriginalName     string  `json:"original_name"`
	Overview         string  `json:"overview"`
	PosterPath       *string `json:"poster_path"`
}

type SearchCollections struct {
	pagination
	Collections []SearchCollection `json:"results"`
}

type SearchCollectionsOptions struct {
	// Pass a ISO 639-1 value to display translated data for the fields that support it.
	// minLength: 2
	// pattern: ([a-z]{2})-([A-Z]{2})
	// default: en-US
	// If the provided language is wrong, it is ignored.
	Language string `url:"language,omitempty" json:"language,omitempty"`

	// Specify which page to query.
	Page *int `url:"page,omitempty" json:"page,omitempty"`
}

// Search for collections.
func (sr *SearchResource) Collections(query string, opt *SearchCollectionsOptions) (*SearchCollections, *http.Response, error) {
	path := "/search/collection"
	var collections SearchCollections
	resp, err := sr.client.get(path, &collections, WithQueryParam("query", query), WithQueryParams(opt))
	return &collections, resp, errors.Wrap(err, "failed to search for collections")
}

type SearchKeywordsOptions struct {
	// Specify which page to query.
	Page *int `url:"page,omitempty" json:"page,omitempty"`
}

type SearchKeywords struct {
	pagination
	Keywords []Keyword `json:"results"`
}

// Search for keywords.
func (sr *SearchResource) Keywords(query string, opt *SearchKeywordsOptions) (*SearchKeywords, *http.Response, error) {
	path := "/search/keyword"
	var keywords SearchKeywords
	resp, err := sr.client.get(path, &keywords, WithQueryParam("query", query), WithQueryParams(opt))
	return &keywords, resp, errors.Wrap(err, "failed to search for keywords")
}

type SearchMoviesOptions struct {
	// Pass a ISO 639-1 value to display translated data for the fields that support it.
	// minLength: 2
	// pattern: ([a-z]{2})-([A-Z]{2})
	// default: en-US
	// If the provided language is wrong, it is ignored.
	Language string `url:"language,omitempty" json:"language,omitempty"`

	// Specify which page to query.
	Page *int `url:"page,omitempty" json:"page,omitempty"`

	// A filter and include or exclude adult movies.
	// default: false
	IncludeAdult bool `url:"include_adult,omitempty" json:"include_adult,omitempty"`

	// A filter to limit the results to a specific year (looking at all release dates).
	Year *int `url:"year,omitempty" json:"year,omitempty"`

	// A filter to limit the results to a specific primary release year.
	PrimaryReleaseYear *int `url:"primary_release_year,omitempty" json:"primary_release_year,omitempty"`

	// Specify a ISO 3166-1 code to filter release dates. Must be uppercase.
	Region string `url:"region,omitempty" json:"region,omitempty"`
}

type SearchMovies paginatedMovies

// Search for movies.
func (sr *SearchResource) Movies(query string, opt *SearchMoviesOptions) (*SearchMovies, *http.Response, error) {
	path := "/search/movie"
	var movies SearchMovies
	resp, err := sr.client.get(path, &movies, WithQueryParam("query", query), WithQueryParams(opt))
	return &movies, resp, errors.Wrap(err, "failed to search for movies")
}

type SearchPeopleOptions struct {
	// Pass a ISO 639-1 value to display translated data for the fields that support it.
	// minLength: 2
	// pattern: ([a-z]{2})-([A-Z]{2})
	// default: en-US
	// If the provided language is wrong, it is ignored.
	Language string `url:"language,omitempty" json:"language,omitempty"`

	// Specify which page to query.
	Page *int `url:"page,omitempty" json:"page,omitempty"`

	// A filter and include or exclude adult movies.
	// default: false
	IncludeAdult bool `url:"include_adult,omitempty" json:"include_adult,omitempty"`

	// Specify a ISO 3166-1 code to filter release dates. Must be uppercase.
	Region string `url:"region,omitempty" json:"region,omitempty"`
}

type searchPerson struct {
	Adult              bool        `json:"adult"`
	Gender             int         `json:"gender"`
	Id                 int         `json:"id"`
	KnownFor           []MovieOrTV `json:"known_for"`
	KnownForDepartment string      `json:"known_for_department"`
	Name               string      `json:"name"`
	Popularity         float64     `json:"popularity"`
	ProfilePath        *string     `json:"profile_path"`
}

type SearchPeople struct {
	pagination
	People []searchPerson `json:"results"`
}

// Search for people.
func (sr *SearchResource) People(query string, opt *SearchPeopleOptions) (*SearchPeople, *http.Response, error) {
	path := "/search/person"
	var people SearchPeople
	resp, err := sr.client.get(path, &people, WithQueryParam("query", query), WithQueryParams(opt))
	return &people, resp, errors.Wrap(err, "failed to search for people")
}

type SearchTVShowsOptions struct {
	// Pass a ISO 639-1 value to display translated data for the fields that support it.
	// minLength: 2
	// pattern: ([a-z]{2})-([A-Z]{2})
	// default: en-US
	// If the provided language is wrong, it is ignored.
	Language string `url:"language,omitempty" json:"language,omitempty"`

	// Specify which page to query.
	Page *int `url:"page,omitempty" json:"page,omitempty"`

	// A filter and include or exclude adult movies.
	// default: false
	IncludeAdult bool `url:"include_adult,omitempty" json:"include_adult,omitempty"`

	// Filter and only include TV shows that have a original air date year that equal to the specified value.
	// Can be used in conjunction with the "include_null_first_air_dates" filter if you want to include items with no air date.
	FirstAirDateYear *int `url:"first_air_date_year,omitempty" json:"first_air_date_year,omitempty"`
}

type SearchTVShows paginatedTVShows

// Search for TV shows.
func (sr *SearchResource) TVShows(query string, opt *SearchTVShowsOptions) (*SearchTVShows, *http.Response, error) {
	path := "/search/tv"
	var tvShows SearchTVShows
	resp, err := sr.client.get(path, &tvShows, WithQueryParam("query", query), WithQueryParams(opt))
	return &tvShows, resp, errors.Wrap(err, "failed to search for tv shows")
}

type SearchMultiOptions struct {
	// Pass a ISO 639-1 value to display translated data for the fields that support it.
	// minLength: 2
	// pattern: ([a-z]{2})-([A-Z]{2})
	// default: en-US
	// If the provided language is wrong, it is ignored.
	Language string `url:"language,omitempty" json:"language,omitempty"`

	// Specify which page to query.
	Page *int `url:"page,omitempty" json:"page,omitempty"`

	// A filter and include or exclude adult movies.
	// default: false
	IncludeAdult bool `url:"include_adult,omitempty" json:"include_adult,omitempty"`

	// Specify a ISO 3166-1 code to filter release dates. Must be uppercase.
	Region string `url:"region,omitempty" json:"region,omitempty"`
}

type SearchPerson struct {
	searchPerson
	MediaType string `json:"media_type"`
}

type SearchMultiResults map[string]interface{}

type SearchMulti struct {
	pagination
	Results []SearchMultiResults `json:"results"`
}

// Search multiple models in a single request.
// Multi search currently supports searching for movies, tv shows and people in a single request.
func (sr *SearchResource) Multi(query string, opt *SearchTVShowsOptions) (*SearchMulti, *http.Response, error) {
	path := "/search/multi"
	var multi SearchMulti
	resp, err := sr.client.get(path, &multi, WithQueryParam("query", query), WithQueryParams(opt))
	return &multi, resp, errors.Wrap(err, "failed to search multi media")
}

func (sr SearchMultiResults) GetMediaType() string {
	return sr["media_type"].(string)
}

func (sr SearchMultiResults) ToMovie() (*Movie, error) {
	if sr.GetMediaType() != "movie" {
		return nil, errors.New(fmt.Sprintf("invalid conversion from %s to movie", sr.GetMediaType()))
	}
	return convertToMovie(sr)
}

func (sr SearchMultiResults) ToTVShow() (*TVShow, error) {
	if sr.GetMediaType() != "tv" {
		return nil, errors.New(fmt.Sprintf("invalid conversion from %s to tv", sr.GetMediaType()))
	}
	return convertToTVShow(sr)
}

func (sr SearchMultiResults) ToPerson() (*SearchPerson, error) {
	if sr.GetMediaType() != "person" {
		return nil, errors.New(fmt.Sprintf("invalid conversion from %s to person", sr.GetMediaType()))
	}
	result, err := json.Marshal(sr)
	if err != nil {
		return nil, err
	}
	var person SearchPerson
	err = json.Unmarshal(result, &person)
	return &person, err
}
