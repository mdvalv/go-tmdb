package tmdb

import (
	"time"

	"github.com/hashicorp/go-retryablehttp"
)

// HTTPClientOptionFunc can be used to customize the retryable HTTP client used in TMDb client.
type HTTPClientOptionFunc func(*retryablehttp.Client)

// CustomRetryMax can be used to configure a custom number of retries for the request.
// The default value is 4.
func CustomRetryMax(maxRetries int) HTTPClientOptionFunc {
	return func(c *retryablehttp.Client) {
		c.RetryMax = maxRetries
	}
}

// CustomRetryWaitMin can be used to configure a custom minimum time to wait between requests.
// The default value is 1 second.
func CustomRetryWaitMin(retryWaitMin time.Duration) HTTPClientOptionFunc {
	return func(c *retryablehttp.Client) {
		c.RetryWaitMin = retryWaitMin
	}
}

// CustomRetryWaitMax can be used to configure a custom maximum time to wait between requests.
// The default value is 30 seconds.
func CustomRetryWaitMax(retryWaitMax time.Duration) HTTPClientOptionFunc {
	return func(c *retryablehttp.Client) {
		c.RetryWaitMax = retryWaitMax
	}
}

// CustomLeveledLogger can be used to configure a custom retryablehttp leveled logger.
func CustomLeveledLogger(leveledLogger retryablehttp.LeveledLogger) HTTPClientOptionFunc {
	return func(c *retryablehttp.Client) {
		c.Logger = leveledLogger
	}
}

// WithCustomLogger can be used to configure a custom retryablehttp logger.
func CustomLogger(logger retryablehttp.Logger) HTTPClientOptionFunc {
	return func(c *retryablehttp.Client) {
		c.Logger = logger
	}
}
