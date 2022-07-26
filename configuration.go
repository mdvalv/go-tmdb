package tmdb

import (
	"net/http"

	"github.com/pkg/errors"
)

// ConfigurationResource handles configuration-related requests of TMDb API.
type ConfigurationResource struct {
	client *Client
}

// ConfigurationImages represents configuration images in TMDb.
type ConfigurationImages struct {
	BaseURL       string   `json:"base_url"`
	SecureBaseURL string   `json:"secure_base_url"`
	BackdropSizes []string `json:"backdrop_sizes"`
	LogoSizes     []string `json:"logo_sizes"`
	PosterSizes   []string `json:"poster_sizes"`
	ProfileSizes  []string `json:"profile_sizes"`
	StillSizes    []string `json:"still_sizes"`
}

// Configuration represents the configuration in TMDb.
type Configuration struct {
	Images     ConfigurationImages `json:"images"`
	ChangeKeys []string            `json:"change_keys"`
}

// GetAPIConfiguration retrieves the system wide configuration information.
// Some elements of the API require some knowledge of this configuration data.
// The purpose of this is to try and keep the actual API responses as light as possible.
// It is recommended to cache this data within the application and check for updates every few days.
// This method currently holds the data relevant to building image URLs as well as the change key map.
// To build an image URL, 3 pieces of data are needed. The `base_url`, `size` and `file_path`.
// Simply combine them all for a fully qualified URL. Here’s an example URL:
// https://image.tmdb.org/t/p/w500/8uO0gUM8aNqYLs1OsTBQiXu0fEv.jpg
// The configuration method also contains the list of change keys which can be useful
// if building an app that consumes data from the change feed.
func (cr *ConfigurationResource) GetAPIConfiguration() (*Configuration, *http.Response, error) {
	path := "/configuration"
	var configuration Configuration
	resp, err := cr.client.get(path, &configuration)
	return &configuration, resp, errors.Wrap(err, "failed to get API configuration")
}

// Countries represents countries in TMDb.
type Countries []struct {
	ISO31661    string `json:"iso_3166_1"`
	EnglishName string `json:"english_name"`
}

// GetCountries retrieves the list of countries (ISO 3166-1 tags) used throughout TMDB.
func (cr *ConfigurationResource) GetCountries() (Countries, *http.Response, error) {
	path := "/configuration/countries"
	var countries Countries
	resp, err := cr.client.get(path, &countries)
	return countries, resp, errors.Wrap(err, "failed to get countries")
}

// Jobs represents jobs in TMDb.
type Jobs []struct {
	Department string   `json:"department"`
	Jobs       []string `json:"jobs"`
}

// GetJobs retrieves a list of the jobs and departments used on TMDB.
func (cr *ConfigurationResource) GetJobs() (Jobs, *http.Response, error) {
	path := "/configuration/jobs"
	var jobs Jobs
	resp, err := cr.client.get(path, &jobs)
	return jobs, resp, errors.Wrap(err, "failed to get jobs")
}

// Languages represents languages in TMDb.
type Languages []struct {
	ISO6391     string `json:"iso_639_1"`
	EnglishName string `json:"english_name"`
	Name        string `json:"name"`
}

// GetLanguages retrieves the list of languages (ISO 639-1 tags) used throughout TMDB.
func (cr *ConfigurationResource) GetLanguages() (Languages, *http.Response, error) {
	path := "/configuration/languages"
	var languages Languages
	resp, err := cr.client.get(path, &languages)
	return languages, resp, errors.Wrap(err, "failed to get languages")
}

// PrimaryTranslations represents primary translations in TMDb.
type PrimaryTranslations []string

// GetPrimaryTranslations retrieves a list of the officially supported translations on TMDB.
func (cr *ConfigurationResource) GetPrimaryTranslations() (PrimaryTranslations, *http.Response, error) {
	path := "/configuration/primary_translations"
	var translations PrimaryTranslations
	resp, err := cr.client.get(path, &translations)
	return translations, resp, errors.Wrap(err, "failed to get primary translations")
}

// Timezones represents timezones in TMDb.
type Timezones []struct {
	ISO31661 string   `json:"iso_3166_1"`
	Zones    []string `json:"zones"`
}

// GetTimezones retrieves the list of timezones used throughout TMDB.
func (cr *ConfigurationResource) GetTimezones() (Timezones, *http.Response, error) {
	path := "/configuration/timezones"
	var timezones Timezones
	resp, err := cr.client.get(path, &timezones)
	return timezones, resp, errors.Wrap(err, "failed to get timezones")
}
