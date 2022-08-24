package tmdb

import (
	"net/http"

	"github.com/pkg/errors"
)

// SearchResource handles search-related requests of TMDb API.
type SearchResource struct {
	client *Client
}

// SearchCompanies represents companies in TMDb.
type SearchCompanies struct {
	pagination
	Results []Company `json:"results"`
}

// SearchCompaniesOptions represents the available options for the request.
type SearchCompaniesOptions struct {
	// Specify which page to query.
	Page *int `url:"page,omitempty" json:"page,omitempty"`
}

// Companies searches for companies.
func (sr *SearchResource) Companies(query string, opt *SearchCompaniesOptions) (*SearchCompanies, *http.Response, error) {
	path := "/search/company"
	var companies SearchCompanies
	resp, err := sr.client.get(path, &companies, WithQueryParam("query", query), WithQueryParams(opt))
	return &companies, resp, errors.Wrap(err, "failed to search for companies")
}

// SearchCollection represents a collection in TMDb.
type SearchCollection struct {
	Adult            bool    `json:"adult"`
	BackdropPath     *string `json:"backdrop_path"`
	ID               int     `json:"id"`
	Name             string  `json:"name"`
	OriginalLanguage string  `json:"original_language"`
	OriginalName     string  `json:"original_name"`
	Overview         string  `json:"overview"`
	PosterPath       *string `json:"poster_path"`
}

// SearchCollections represents collections in TMDb.
type SearchCollections struct {
	pagination
	Collections []SearchCollection `json:"results"`
}

// SearchCollectionsOptions represents the available options for the request.
type SearchCollectionsOptions languagePageOptions

// Collections searches for collections.
func (sr *SearchResource) Collections(query string, opt *SearchCollectionsOptions) (*SearchCollections, *http.Response, error) {
	path := "/search/collection"
	var collections SearchCollections
	resp, err := sr.client.get(path, &collections, WithQueryParam("query", query), WithQueryParams(opt))
	return &collections, resp, errors.Wrap(err, "failed to search for collections")
}

// SearchKeywordsOptions represents the available options for the request.
type SearchKeywordsOptions struct {
	// Specify which page to query.
	Page *int `url:"page,omitempty" json:"page,omitempty"`
}

// SearchKeywords represents keywords in TMDb.
type SearchKeywords struct {
	pagination
	Keywords []Keyword `json:"results"`
}

// Keywords searches for keywords.
func (sr *SearchResource) Keywords(query string, opt *SearchKeywordsOptions) (*SearchKeywords, *http.Response, error) {
	path := "/search/keyword"
	var keywords SearchKeywords
	resp, err := sr.client.get(path, &keywords, WithQueryParam("query", query), WithQueryParams(opt))
	return &keywords, resp, errors.Wrap(err, "failed to search for keywords")
}

// SearchMoviesOptions represents the available options for the request.
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

// SearchMovies represents movies in TMDb.
type SearchMovies paginatedMovies

// Movies searches for movies.
func (sr *SearchResource) Movies(query string, opt *SearchMoviesOptions) (*SearchMovies, *http.Response, error) {
	path := "/search/movie"
	var movies SearchMovies
	resp, err := sr.client.get(path, &movies, WithQueryParam("query", query), WithQueryParams(opt))
	return &movies, resp, errors.Wrap(err, "failed to search for movies")
}

// SearchPeopleOptions represents the available options for the request.
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

// MovieOrTV can be either a movie or a tv show object in TMDb.
type MovieOrTV map[string]interface{}

// GetMediaType retrieves the media type from a movie or tv show.
func (mt MovieOrTV) GetMediaType() string {
	return mt["media_type"].(string)
}

// ToMovie converts the data to a movie.
func (mt MovieOrTV) ToMovie() (*Movie, error) {
	return convertToMovie(mt)
}

// ToTVShow converts the data to a tv show.
func (mt MovieOrTV) ToTVShow() (*TVShow, error) {
	return convertToTVShow(mt)
}

// SearchPerson represents a person in TMDb.
type SearchPerson struct {
	Adult              bool        `json:"adult"`
	Gender             int         `json:"gender"`
	ID                 int         `json:"id"`
	KnownFor           []MovieOrTV `json:"known_for"`
	KnownForDepartment string      `json:"known_for_department"`
	Name               string      `json:"name"`
	Popularity         float64     `json:"popularity"`
	ProfilePath        *string     `json:"profile_path"`
}

// SearchPeople represents people in TMDb.
type SearchPeople struct {
	pagination
	People []SearchPerson `json:"results"`
}

// People searches for people.
func (sr *SearchResource) People(query string, opt *SearchPeopleOptions) (*SearchPeople, *http.Response, error) {
	path := "/search/person"
	var people SearchPeople
	resp, err := sr.client.get(path, &people, WithQueryParam("query", query), WithQueryParams(opt))
	return &people, resp, errors.Wrap(err, "failed to search for people")
}

// SearchTVShowsOptions represents the available options for the request.
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
	// Can be used in conjunction with the "include_null_first_air_dates" filter to include items with no air date.
	FirstAirDateYear *int `url:"first_air_date_year,omitempty" json:"first_air_date_year,omitempty"`
}

// SearchTVShows represents tv shows in TMDb.
type SearchTVShows paginatedTVShows

// TVShows searches for TV shows.
func (sr *SearchResource) TVShows(query string, opt *SearchTVShowsOptions) (*SearchTVShows, *http.Response, error) {
	path := "/search/tv"
	var tvShows SearchTVShows
	resp, err := sr.client.get(path, &tvShows, WithQueryParam("query", query), WithQueryParams(opt))
	return &tvShows, resp, errors.Wrap(err, "failed to search for tv shows")
}

// SearchMultiOptions represents the available options for the request.
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

// SearchMultiPerson represents a person in TMDb.
type SearchMultiPerson struct {
	Adult              bool        `json:"adult"`
	Gender             int         `json:"gender"`
	ID                 int         `json:"id"`
	KnownFor           []MovieOrTV `json:"known_for"`
	KnownForDepartment string      `json:"known_for_department"`
	Name               string      `json:"name"`
	Popularity         float64     `json:"popularity"`
	ProfilePath        *string     `json:"profile_path"`
	MediaType          string      `json:"media_type"`
}

// SearchMultiResults represents results for a multi search in TMDb.
type SearchMultiResults map[string]interface{}

// SearchMulti represents a multi search in TMDb.
type SearchMulti struct {
	pagination
	Results []SearchMultiResults `json:"results"`
}

// Multi searches multiple models in a single request.
// Multi search currently supports searching for movies, tv shows and people in a single request.
func (sr *SearchResource) Multi(query string, opt *SearchTVShowsOptions) (*SearchMulti, *http.Response, error) {
	path := "/search/multi"
	var multi SearchMulti
	resp, err := sr.client.get(path, &multi, WithQueryParam("query", query), WithQueryParams(opt))
	return &multi, resp, errors.Wrap(err, "failed to search multi media")
}

// GetMediaType retrieves the media type from search multi results.
func (sr SearchMultiResults) GetMediaType() string {
	return sr["media_type"].(string)
}

// ToMovie converts the data to a movie.
func (sr SearchMultiResults) ToMovie() (*Movie, error) {
	return convertToMovie(sr)
}

// ToTVShow converts the data to a tv show.
func (sr SearchMultiResults) ToTVShow() (*TVShow, error) {
	return convertToTVShow(sr)
}

// ToPerson converts the data to a person.
func (sr SearchMultiResults) ToPerson() (*SearchMultiPerson, error) {
	var person SearchMultiPerson
	err := convert("person", sr, &person)
	return &person, err
}
