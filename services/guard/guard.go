package guard

import (
	"fmt"
	"net/http"
	"strings"
)

type Options struct {
	// JWTAuthToken is an optional JWT token to authenticate with the guard service.
	JWTAuthToken string

	// GuardServiceURL is an optional custom URL for the Sequence Guard service.
	// If not provided, the default URL in `DefaultGuardServiceURL` will be used.
	GuardServiceURL string

	// HTTPClient is an optional custom HTTP client to use for communicating with the
	// guard service.
	HTTPClient HTTPClient
}

const DefaultGuardServiceURL = "https://guard.sequence.app"

// NewClient creates a new Sequence Guard client instance. Please see https://sequence.build to
// get a `projectAccessKey`, which is your project's access key used to communicate
// with Sequence services.
//
// NOTE: the `projectAccessKey` may be optional if you're using a JWT auth token
// passed in via the `clientOptions`.
func NewClient(projectAccessKey string, clientOptions ...Options) GuardClient {
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

	serviceURL := DefaultGuardServiceURL
	if opts.GuardServiceURL != "" {
		serviceURL = opts.GuardServiceURL
	}
	return NewGuardClient(strings.TrimSuffix(serviceURL, "/"), c)
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
