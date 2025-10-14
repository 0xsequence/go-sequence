package keymachine

import (
	"fmt"
	"net/http"
	"strings"
)

type Options struct {
	// ProjectAccessKey is an optional project access key to authenticate with
	// the keymachine service.
	ProjectAccessKey string

	// JWTAuthToken is an optional JWT token to authenticate with the keymachine service.
	JWTAuthToken string

	// KeymachineServiceURL is an optional custom URL for the Sequence Keymachine service.
	// If not provided, the default URL in `DefaultKeymachineServiceURL` will be used.
	KeymachineServiceURL string

	// HTTPClient is an optional custom HTTP client to use for communicating with the
	// keymachine service.
	HTTPClient HTTPClient
}

const DefaultKeymachineServiceURL = "https://sessions.sequence.app"

// NewClient creates a new Sequence Keymachine client instance.
func NewClient(clientOptions ...Options) SessionsClient {
	opts := Options{}
	if len(clientOptions) > 0 {
		opts = clientOptions[0]
	}

	c := &httpClient{
		client:           opts.HTTPClient,
		projectAccessKey: opts.ProjectAccessKey,
	}
	if opts.HTTPClient == nil {
		c.client = http.DefaultClient
	}
	if opts.JWTAuthToken != "" {
		c.jwtAuthHeader = fmt.Sprintf("BEARER %s", opts.JWTAuthToken)
	}

	serviceURL := DefaultKeymachineServiceURL
	if opts.KeymachineServiceURL != "" {
		serviceURL = opts.KeymachineServiceURL
	}
	return NewSessionsClient(strings.TrimSuffix(serviceURL, "/"), c)
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
