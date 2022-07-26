// Package tmdb is a complete implementation for TMDb API v3.
package tmdb

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"

	"github.com/pkg/errors"
)

const (
	// BaseURL is the base URL for TMDb API.
	BaseURL = "https://api.themoviedb.org/3"
)

// Client handles interaction with TMDb API.
type Client struct {
	// HTTP client used to communicate with the API.
	HTTPClient *resty.Client

	// Available TMDb resources that can be interacted with through the API.
	Account        *AccountResource
	Authentication *AuthenticationResource
	Certifications *CertificationsResource
	Collections    *CollectionsResource
	Companies      *CompaniesResource
	Configuration  *ConfigurationResource
	Credits        *CreditsResource
	Discover       *DiscoverResource
	Find           *FindResource
	Genres         *GenresResource
	GuestSession   *GuestSessionResource
	Keywords       *KeywordsResource
	Lists          *ListsResource
	Movies         *MoviesResource
	Networks       *NetworksResource
	People         *PeopleResource
	Reviews        *ReviewsResource
	Search         *SearchResource
	Trending       *TrendingResource
	TV             *TVResource
	TVEpisodes     *TVEpisodesResource
	TVSeasons      *TVSeasonsResource
	WatchProviders *WatchProvidersResource
}

// getRestyClient adds some custom configuration to the HTTP client used by TMDb client.
func getRestyClient(token, baseURL string) *resty.Client {
	client := resty.New()
	client.SetBaseURL(baseURL)
	client.SetQueryParam("api_key", token)
	client.SetHeader("Accept", "application/json")
	client.OnAfterResponse(func(c *resty.Client, resp *resty.Response) error {
		return checkResponse(resp)
	})
	return client
}

// NewClient returns a new TMDb API client.
func NewClient(token string) (*Client, error) {
	c := &Client{
		HTTPClient: getRestyClient(token, BaseURL),
	}

	c.Account = &AccountResource{client: c}
	c.Authentication = &AuthenticationResource{client: c}
	c.Certifications = &CertificationsResource{client: c}
	c.Collections = &CollectionsResource{client: c}
	c.Companies = &CompaniesResource{client: c}
	c.Configuration = &ConfigurationResource{client: c}
	c.Credits = &CreditsResource{client: c}
	c.Discover = &DiscoverResource{client: c}
	c.Find = &FindResource{client: c}
	c.Genres = &GenresResource{client: c}
	c.GuestSession = &GuestSessionResource{client: c}
	c.Keywords = &KeywordsResource{client: c}
	c.Lists = &ListsResource{client: c}
	c.Movies = &MoviesResource{client: c}
	c.Networks = &NetworksResource{client: c}
	c.People = &PeopleResource{client: c}
	c.Reviews = &ReviewsResource{client: c}
	c.Search = &SearchResource{client: c}
	c.Trending = &TrendingResource{client: c}
	c.TV = &TVResource{client: c}
	c.TVEpisodes = &TVEpisodesResource{client: c}
	c.TVSeasons = &TVSeasonsResource{client: c}
	c.WatchProviders = &WatchProvidersResource{client: c}

	return c, nil
}

// checkResponse checks the API response for errors, and returns them if present.
func checkResponse(resp *resty.Response) error {
	switch resp.StatusCode() {
	case 200, 201, 202, 204, 304:
		return nil
	}

	var err error
	if resp.Body() != nil {
		var raw interface{}
		if err = json.Unmarshal(resp.Body(), &raw); err != nil {
			return errors.Wrap(err, "failed to parse unknown error format")
		}
		return errors.New(parseError(raw))
	}

	return err
}

// parseError parses the error trying to make them more presentable.
func parseError(raw interface{}) string {
	switch raw := raw.(type) {
	case string:
		return raw
	case int, float32, float64, bool:
		return fmt.Sprint(raw)
	case []interface{}:
		var errs []string
		for _, v := range raw {
			errs = append(errs, parseError(v))
		}
		return fmt.Sprintf("[%s]", strings.Join(errs, ", "))
	case map[string]interface{}:
		var errs []string
		for k, v := range raw {
			errs = append(errs, fmt.Sprintf("{%s: %s}", k, parseError(v)))
		}
		sort.Strings(errs)
		return strings.Join(errs, ", ")
	default:
		return fmt.Sprintf("failed to parse unexpected error type: %T", raw)
	}
}

// RequestOptionFn can be used to customize the request fields.
type RequestOptionFn func(*resty.Request) error

// WithBody can be used to set a custom request body.
func WithBody(body interface{}) RequestOptionFn {
	return func(r *resty.Request) error {
		if body != nil {
			r.SetHeader("Content-Type", "application/json")
			r.SetBody(body)
		}
		return nil
	}
}

// WithQueryParams can be used to set custom query parameters to the request.
func WithQueryParams(params interface{}) RequestOptionFn {
	return func(r *resty.Request) error {
		q, err := query.Values(params)
		if err != nil {
			return errors.Wrap(err, "failed to prepare request query params")
		}
		r.SetQueryParamsFromValues(q)
		return nil
	}
}

// WithQueryParam can be used to set a custom query parameter to the request.
func WithQueryParam(param, value string) RequestOptionFn {
	return func(r *resty.Request) error {
		r.SetQueryParam(param, value)
		return nil
	}
}

// WithSessionID can be used to set a session ID to the request.
func WithSessionID(sessionID string) RequestOptionFn {
	return func(r *resty.Request) error {
		r.SetQueryParam("session_id", sessionID)
		return nil
	}
}

// newRequest prepares a new resty.Request.
func (c *Client) newRequest(resource interface{}, options ...RequestOptionFn) (*resty.Request, error) {
	req := c.HTTPClient.NewRequest().SetResult(resource)
	for _, fn := range options {
		if fn != nil {
			if err := fn(req); err != nil {
				return nil, err
			}
		}
	}
	return req, nil
}

// get performs a get request.
func (c *Client) get(path string, resource interface{}, options ...RequestOptionFn) (*http.Response, error) {
	req, err := c.newRequest(resource, options...)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get request")
	}
	resp, err := req.Get(path)
	return resp.RawResponse, errors.Wrap(err, "failed to execute request")
}

// delete performs a delete request.
func (c *Client) delete(path string, resource interface{}, options ...RequestOptionFn) (*http.Response, error) {
	req, err := c.newRequest(resource, options...)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get request")
	}
	resp, err := req.Delete(path)
	return resp.RawResponse, errors.Wrap(err, "failed to execute request")
}

// post performs post request.
func (c *Client) post(path string, resource interface{}, options ...RequestOptionFn) (*http.Response, error) {
	req, err := c.newRequest(resource, options...)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get request")
	}
	resp, err := req.Post(path)
	return resp.RawResponse, errors.Wrap(err, "failed to execute request")
}

type media interface {
	GetMediaType() string
}

func convertToMovie(obj media) (*Movie, error) {
	var movie Movie
	err := convert("movie", obj, &movie)
	return &movie, errors.Wrap(err, "failed to convert object to movie")
}

func convertToTVShow(obj media) (*TVShow, error) {
	var tvShow TVShow
	err := convert("tv", obj, &tvShow)
	return &tvShow, errors.Wrap(err, "failed to convert object to tv show")
}

func convert(expectedMedia string, obj media, to interface{}) error {
	if obj.GetMediaType() != expectedMedia {
		return fmt.Errorf("invalid conversion from %s to %s", obj.GetMediaType(), expectedMedia)
	}
	result, err := json.Marshal(obj)
	if err != nil {
		return errors.Wrap(err, "failed to marshal object")
	}
	return json.Unmarshal(result, to)
}

type statusResponse struct {
	StatusCode    int    `json:"status_code"`
	StatusMessage string `json:"status_message"`
}

type pagination struct {
	Page         int `json:"page"`
	TotalPages   int `json:"total_pages"`
	TotalResults int `json:"total_results"`
}
