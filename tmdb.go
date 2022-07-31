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
	BASE_URL = "https://api.themoviedb.org/3"
)

// Client handles interaction with TMDb API.
type Client struct {
	// HTTP client used to communicate with the API.
	HttpClient *resty.Client

	// Available TMDb resources that can be interacted with through the API.
	Authentication *AuthenticationResource
	Certifications *CertificationsResource
	Collections    *CollectionsResource
	Companies      *CompaniesResource
	Configuration  *ConfigurationResource
	Credits        *CreditsResource
	Genres         *GenresResource
	Keywords       *KeywordsResource
	Networks       *NetworksResource
	Reviews        *ReviewsResource
	Trending       *TrendingResource
	WatchProviders *WatchProvidersResource
}

// getRestyClient adds some custom configuration to the HTTP client used by TMDb client.
func getRestyClient(token, baseUrl string) *resty.Client {
	client := resty.New()
	client.SetBaseURL(baseUrl)
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
		HttpClient: getRestyClient(token, BASE_URL),
	}

	c.Authentication = &AuthenticationResource{client: c}
	c.Certifications = &CertificationsResource{client: c}
	c.Collections = &CollectionsResource{client: c}
	c.Companies = &CompaniesResource{client: c}
	c.Configuration = &ConfigurationResource{client: c}
	c.Credits = &CreditsResource{client: c}
	c.Genres = &GenresResource{client: c}
	c.Keywords = &KeywordsResource{client: c}
	c.Networks = &NetworksResource{client: c}
	c.Reviews = &ReviewsResource{client: c}
	c.Trending = &TrendingResource{client: c}
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
			err = errors.Wrap(err, "failed to parse unknown error format")
			return err
		} else {
			return errors.New(parseError(raw))
		}
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

// requestOptionFn can be used to customize the request fields.
type requestOptionFn func(*resty.Request) error

// WithBody can be used to set a custom request body.
func WithBody(body interface{}) requestOptionFn {
	return func(r *resty.Request) error {
		if body != nil {
			r.SetHeader("Content-Type", "application/json")
			r.SetBody(body)
		}
		return nil
	}
}

// WithQueryParams can be used to set custom query parameters to the request.
func WithQueryParams(params interface{}) requestOptionFn {
	return func(r *resty.Request) error {
		q, err := query.Values(params)
		if err != nil {
			return errors.Wrap(err, "failed to prepare request query params")
		}
		r.SetQueryParamsFromValues(q)
		return nil
	}
}

// newRequest prepares a new resty.Request.
func (c *Client) newRequest(resource interface{}, options ...requestOptionFn) (*resty.Request, error) {
	req := c.HttpClient.NewRequest().SetResult(resource)
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
func (c *Client) get(path string, resource interface{}, options ...requestOptionFn) (*http.Response, error) {
	req, err := c.newRequest(resource, options...)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get request")
	}
	resp, err := req.Get(path)
	return resp.RawResponse, errors.Wrap(err, "failed to execute request")
}

// delete performs a delete request.
func (c *Client) delete(path string, resource interface{}, options ...requestOptionFn) (*http.Response, error) {
	req, err := c.newRequest(resource, options...)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get request")
	}
	resp, err := req.Delete(path)
	return resp.RawResponse, errors.Wrap(err, "failed to execute request")
}

// post performs post request.
func (c *Client) post(path string, resource interface{}, options ...requestOptionFn) (*http.Response, error) {
	req, err := c.newRequest(resource, options...)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get request")
	}
	resp, err := req.Post(path)
	return resp.RawResponse, errors.Wrap(err, "failed to execute request")
}
