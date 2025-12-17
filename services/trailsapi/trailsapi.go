package trailsapi

import (
	"fmt"
	"net/http"
	"strings"
)

type Options struct {
	// JWTAuthToken is an optional JWT token to authenticate with the api service.
	JWTAuthToken string

	// ProjectAccessKey is an optional project access key to authenticate with the api service.
	// This may be required if no `JWTAuthToken` is provided.
	ProjectAccessKey string

	// TrailsAPIServiceURL is an optional custom URL for the Trails API service.
	// If not provided, the default URL in `DefaultAPIServiceURL` will be used.
	TrailsAPIServiceURL string

	// HTTPClient is an optional custom HTTP client to use for communicating with the
	// api service.
	HTTPClient HTTPClient
}

const DefaultAPIServiceURL = "https://trails-api.sequence.app"

// NewClient creates a new Trails API client instance. Please see https://sequence.build to
// get a `projectAccessKey`, which is your project's access key used to communicate
// with Sequence services.
//
// NOTE: the `projectAccessKey` may be optional if you're using a JWT auth token
// passed in via the `clientOptions`.
func NewClient(clientOptions ...Options) TrailsClient {
	opts := Options{}
	if len(clientOptions) > 0 {
		opts = clientOptions[0]
	}

	c := &httpClient{
		client:           opts.HTTPClient,
		projectAccessKey: "",
	}
	if opts.HTTPClient == nil {
		c.client = http.DefaultClient
	}
	if opts.JWTAuthToken != "" {
		c.jwtAuthHeader = fmt.Sprintf("BEARER %s", opts.JWTAuthToken)
	} else if opts.ProjectAccessKey != "" {
		c.projectAccessKey = opts.ProjectAccessKey
	}

	serviceURL := DefaultAPIServiceURL
	if opts.TrailsAPIServiceURL != "" {
		serviceURL = opts.TrailsAPIServiceURL
	}
	return NewTrailsClient(strings.TrimSuffix(serviceURL, "/"), c)
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
