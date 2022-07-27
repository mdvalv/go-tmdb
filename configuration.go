package tmdb

import (
	"net/http"

	"github.com/pkg/errors"
)

// ConfigurationResource handles configuration-related requests of TMDb API.
type ConfigurationResource struct {
	client *Client
}

type imagesConfiguration struct {
	BaseUrl       string   `json:"base_url"`
	SecureBaseUrl string   `json:"secure_base_url"`
	BackdropSizes []string `json:"backdrop_sizes"`
	LogoSizes     []string `json:"logo_sizes"`
	PosterSizes   []string `json:"poster_sizes"`
	ProfileSizes  []string `json:"profile_sizes"`
	StillSizes    []string `json:"still_sizes"`
}

type Configuration struct {
	Images     imagesConfiguration `json:"images"`
	ChangeKeys []string            `json:"change_keys"`
}

// Get the system wide configuration information.
// Some elements of the API require some knowledge of this configuration data.
// The purpose of this is to try and keep the actual API responses as light as possible.
// It is recommended you cache this data within your application and check for updates every few days.
// This method currently holds the data relevant to building image URLs as well as the change key map.
// To build an image URL, you will need 3 pieces of data. The `base_url`, `size` and `file_path`.
// Simply combine them all and you will have a fully qualified URL. Here’s an example URL:
// https://image.tmdb.org/t/p/w500/8uO0gUM8aNqYLs1OsTBQiXu0fEv.jpg
// The configuration method also contains the list of change keys which can be useful
// if you are building an app that consumes data from the change feed.
func (cr *ConfigurationResource) GetAPIConfiguration() (*Configuration, *http.Response, error) {
	path := "/configuration"
	var configuration Configuration
	resp, err := cr.client.getResource(path, nil, &configuration)
	return &configuration, resp, errors.Wrap(err, "failed to get API configuration")
}

type Countries []struct {
	Iso31661    string `json:"iso_3166_1"`
	EnglishName string `json:"english_name"`
}

// Get the list of countries (ISO 3166-1 tags) used throughout TMDB.
func (cr *ConfigurationResource) GetCountries() (Countries, *http.Response, error) {
	path := "/configuration/countries"
	var countries Countries
	resp, err := cr.client.getResource(path, nil, &countries)
	return countries, resp, errors.Wrap(err, "failed to get countries")
}

type Jobs []struct {
	Department string   `json:"department"`
	Jobs       []string `json:"jobs"`
}

// Get a list of the jobs and departments we use on TMDB.
func (cr *ConfigurationResource) GetJobs() (Jobs, *http.Response, error) {
	path := "/configuration/jobs"
	var jobs Jobs
	resp, err := cr.client.getResource(path, nil, &jobs)
	return jobs, resp, errors.Wrap(err, "failed to get jobs")
}

type Languages []struct {
	ISO6391     string `json:"iso_639_1"`
	EnglishName string `json:"english_name"`
	Name        string `json:"name"`
}

// Get the list of languages (ISO 639-1 tags) used throughout TMDB.
func (cr *ConfigurationResource) GetLanguages() (Languages, *http.Response, error) {
	path := "/configuration/languages"
	var languages Languages
	resp, err := cr.client.getResource(path, nil, &languages)
	return languages, resp, errors.Wrap(err, "failed to get languages")
}

type PrimaryTranslations []string

// Get a list of the officially supported translations on TMDB.
func (cr *ConfigurationResource) GetPrimaryTranslations() (PrimaryTranslations, *http.Response, error) {
	path := "/configuration/primary_translations"
	var translations PrimaryTranslations
	resp, err := cr.client.getResource(path, nil, &translations)
	return translations, resp, errors.Wrap(err, "failed to get primary translations")
}

type Timezones []struct {
	ISO31661 string   `json:"iso_3166_1"`
	Zones    []string `json:"zones"`
}

// Get the list of timezones used throughout TMDB.
func (cr *ConfigurationResource) GetTimezones() (Timezones, *http.Response, error) {
	path := "/configuration/timezones"
	var timezones Timezones
	resp, err := cr.client.getResource(path, nil, &timezones)
	return timezones, resp, errors.Wrap(err, "failed to get timezones")
}
