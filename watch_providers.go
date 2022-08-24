package tmdb

import (
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

// WatchProvidersResource handles watch providers-related requests of TMDb API.
type WatchProvidersResource struct {
	client *Client
}

// Provider represents a provider in TMDb.
type Provider struct {
	DisplayPriority int     `json:"display_priority"`
	LogoPath        *string `json:"logo_path"`
	ProviderName    string  `json:"provider_name"`
	ProviderID      int     `json:"provider_id"`
}

// providers represents providers in TMDb.
type providers struct {
	Providers []Provider `json:"results"`
}

// ProvidersOptions represents the available options for the request.
type ProvidersOptions struct {
	// Pass a ISO 639-1 value to display translated data for the fields that support it.
	// minLength: 2
	// pattern: ([a-z]{2})-([A-Z]{2})
	// default: en-US
	// If the provided language is wrong, it is ignored.
	Language string `url:"language,omitempty" json:"language,omitempty"`

	// Use the ISO-3166-1 code to filter the providers that are available in a particular country.
	// Example: BR, FR, US
	WatchRegion string `url:"watch_region,omitempty" json:"watch_region,omitempty"`
}

// GetMovieProviders returns a list of the watch provider (OTT/streaming) data TMDb has available for movies.
func (pr *WatchProvidersResource) GetMovieProviders(opt *ProvidersOptions) ([]Provider, *http.Response, error) {
	return pr.getProviders("movie", opt)
}

// GetTVProviders returns a list of the watch provider (OTT/streaming) data TMDb has available for TV series.
func (pr *WatchProvidersResource) GetTVProviders(opt *ProvidersOptions) ([]Provider, *http.Response, error) {
	return pr.getProviders("tv", opt)
}

func (pr *WatchProvidersResource) getProviders(providerType string, opt *ProvidersOptions) ([]Provider, *http.Response, error) {
	path := fmt.Sprintf("/watch/providers/%s", providerType)
	var providers providers
	resp, err := pr.client.get(path, &providers, WithQueryParams(opt))
	return providers.Providers, resp, errors.Wrap(err, fmt.Sprintf("failed to get %s providers", providerType))
}

// ProviderRegion represents a provider region in TMDb.
type ProviderRegion struct {
	EnglishName string `json:"english_name"`
	ISO31661    string `json:"iso_3166_1"`
	NativeName  string `json:"native_name"`
}

// providerRegions represents provider regions in TMDb.
type providerRegions struct {
	ProviderRegions []ProviderRegion `json:"results"`
}

// ProviderRegionsOptions represents the available options for the request.
type ProviderRegionsOptions languageOptions

// GetProviderRegions returns a list of all of the countries TMDb has watch provider (OTT/streaming) data for.
func (pr *WatchProvidersResource) GetProviderRegions(opt *ProviderRegionsOptions) ([]ProviderRegion, *http.Response, error) {
	path := "/watch/providers/regions"
	var providerRegions providerRegions
	resp, err := pr.client.get(path, &providerRegions, WithQueryParams(opt))
	return providerRegions.ProviderRegions, resp, errors.Wrap(err, "failed to get provider regions")
}
