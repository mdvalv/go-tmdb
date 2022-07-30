package tmdb

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
	"strings"

	"github.com/google/go-querystring/query"
	retryablehttp "github.com/hashicorp/go-retryablehttp"
	"github.com/pkg/errors"
)

const (
	BASE_URL = "https://api.themoviedb.org/3"
)

// Client handles interaction with TMDb API.
type Client struct {
	// HTTP client used to communicate with the API.
	// By default, the HTTP client retries communication with the server on some cases,
	// check https://pkg.go.dev/github.com/hashicorp/go-retryablehttp@v0.7.1#DefaultRetryPolicy for more information.
	// Still, some configuration for this client can be changed by using the HTTPClientOptionFunc functions.
	httpClient *retryablehttp.Client

	// Base URL for API requests.
	baseURL *url.URL

	// Token used for API authentication in TMDb.
	token string

	// Available TMDb resources that can be interacted with through the API.
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

// getRetryableClient adds some custom configuration to the HTTP client used by TMDb client.
func getRetryableClient(options []HTTPClientOptionFunc) *retryablehttp.Client {
	retryClient := retryablehttp.NewClient()

	// Disable retryClient logs, because the default is DEBUG
	retryClient.Logger = nil

	for _, fn := range options {
		if fn == nil {
			continue
		}
		fn(retryClient)
	}

	return retryClient
}

// NewClient returns a new TMDb API client.
func NewClient(token string, options ...HTTPClientOptionFunc) (*Client, error) {
	baseUrl, err := url.Parse(BASE_URL)
	if err != nil {
		err = errors.Wrap(err, "failed to parse base URL")
		return nil, err
	}

	retryClient := getRetryableClient(options)
	c := &Client{
		httpClient: retryClient,
		baseURL:    baseUrl,
		token:      token,
	}

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

// newRequest creates a new API request. The path should always be specified without a preceding slash.
func (c *Client) newRequest(method, path string, opt interface{}) (*retryablehttp.Request, error) {
	u := *c.baseURL
	u.Path = c.baseURL.Path + path

	reqHeaders := make(http.Header)
	reqHeaders.Set("Accept", "application/json")

	var body []byte
	var err error
	q := url.Values{}
	switch {
	case method == http.MethodPost || method == http.MethodPut || method == http.MethodDelete:
		reqHeaders.Set("Content-Type", "application/json")
		if opt != nil {
			body, err = json.Marshal(opt)
			if err != nil {
				err = errors.Wrap(err, "failed to prepare request body")
				return nil, err
			}
		}
	case opt != nil:
		q, err = query.Values(opt)
		if err != nil {
			err = errors.Wrap(err, "failed to prepare request body")
			return nil, err
		}
	}
	q.Add("api_key", c.token)
	u.RawQuery = q.Encode()

	req, err := retryablehttp.NewRequest(method, u.String(), body)
	if err != nil {
		err = errors.Wrap(err, "failed to get new retryable request")
		return nil, err
	}

	for k, v := range reqHeaders {
		req.Header[k] = v
	}

	return req, nil
}

// do sends an API request and returns the API response. The API response is JSON decoded and stored in the value
// pointed to by resource, or returned as an error if an API error has occurred. If resource implements the io.Writer
// interface, the raw response body will be written to resource, without attempting to first decode it.
func (c *Client) do(req *retryablehttp.Request, resource interface{}) (*http.Response, error) {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		err = errors.Wrap(err, "failed to execute retryable request")
		return nil, err
	}

	err = checkResponse(resp)
	if err != nil {
		return resp, err
	}

	if resource != nil {
		if w, ok := resource.(io.Writer); ok {
			_, err = io.Copy(w, resp.Body)
			err = errors.Wrap(err, "failed to copy request body")
		} else {
			err = json.NewDecoder(resp.Body).Decode(resource)
			err = errors.Wrap(err, "failed to decode response body")
		}
	}

	return resp, err
}

// checkResponse checks the API response for errors, and returns them if present.
func checkResponse(resp *http.Response) error {
	switch resp.StatusCode {
	case 200, 201, 202, 204, 304:
		return nil
	}

	var err error
	data, err := ioutil.ReadAll(resp.Body)
	if err == nil && data != nil {
		var raw interface{}
		if err := json.Unmarshal(data, &raw); err != nil {
			err = errors.Wrap(err, "failed to parse unknown error format")
			return err
		} else {
			return errors.New(parseError(raw))
		}
	}

	return err
}

// parseError parses the error trying to make them more presentable
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

// createResource creates a resource in TMDb.
func (c *Client) createResource(basePath string, opt, resource interface{}) (*http.Response, error) {
	req, err := c.newRequest(http.MethodPost, basePath, opt)
	if err != nil {
		err = errors.Wrap(err, "failed to get new request")
		return nil, err
	}

	resp, err := c.do(req, resource)
	return resp, errors.Wrap(err, "failed to execute request")
}

// getResource retrieves a resource from TMDb.
// If the request should have url values/parameters, they should be included in opt
func (c *Client) getResource(resourcePath string, opt, resource interface{}) (*http.Response, error) {
	req, err := c.newRequest(http.MethodGet, resourcePath, opt)
	if err != nil {
		err = errors.Wrap(err, "failed to get new request")
		return nil, err
	}
	resp, err := c.do(req, resource)
	return resp, errors.Wrap(err, "failed to execute request")
}

// deleteResource deletes a resource from TMDb.
func (c *Client) deleteResource(resourcePath string, opt, resource interface{}) (*http.Response, error) {
	req, err := c.newRequest(http.MethodDelete, resourcePath, opt)
	if err != nil {
		err = errors.Wrap(err, "failed to get new request")
		return nil, err
	}

	resp, err := c.do(req, resource)
	return resp, errors.Wrap(err, "failed to execute request")
}
