package api

import (
	"fmt"
	"net/http"
	"strings"
)

type Options struct {
	// JWTAuthToken is an optional JWT token to authenticate with the api service.
	JWTAuthToken string

	// APIServiceURL is an optional custom URL for the Sequence API service.
	// If not provided, the default URL in `DefaultAPIServiceURL` will be used.
	APIServiceURL string

	// HTTPClient is an optional custom HTTP client to use for communicating with the
	// api service.
	HTTPClient HTTPClient
}

const DefaultAPIServiceURL = "https://api.sequence.app"

// NewClient creates a new Sequence API client instance. Please see https://sequence.build to
// get a `projectAccessKey`, which is your project's access key used to communicate
// with Sequence services.
//
// NOTE: the `projectAccessKey` may be optional if you're using a JWT auth token
// passed in via the `clientOptions`.
func NewClient(projectAccessKey string, clientOptions ...Options) APIClient {
	opts := Options{}
	if len(clientOptions) > 0 {
		opts = clientOptions[0]
	}

	c := &httpClient{
		client:           opts.HTTPClient,
		projectAccessKey: projectAccessKey,
	}
	if opts.HTTPClient == nil {
		c.client = http.DefaultClient
	}
	if opts.JWTAuthToken != "" {
		c.jwtAuthHeader = fmt.Sprintf("BEARER %s", opts.JWTAuthToken)
	}

	serviceURL := DefaultAPIServiceURL
	if opts.APIServiceURL != "" {
		serviceURL = opts.APIServiceURL
	}
	return NewAPIClient(strings.TrimSuffix(serviceURL, "/"), c)
}

type httpClient struct {
	client           HTTPClient
	jwtAuthHeader    string
	projectAccessKey string
}

func (c *httpClient) Do(req *http.Request) (*http.Response, error) {
	if c.projectAccessKey != "" {
		req.Header.Set("X-Access-Key", c.projectAccessKey)
	}
	if c.jwtAuthHeader != "" {
		req.Header.Set("Authorization", c.jwtAuthHeader)
	}
	return c.client.Do(req)
}
