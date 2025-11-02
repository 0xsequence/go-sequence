package metadata

import (
	"fmt"
	"net/http"
	"strings"
)

type Options struct {
	// JWTAuthToken is an optional JWT token to authenticate with the metadata service.
	JWTAuthToken string

	// MetadataServiceURL is an optional custom URL for the Sequence Metadata service.
	// If not provided, the default URL in `DefaultMetadataServiceURL` will be used.
	MetadataServiceURL string

	// HTTPClient is an optional custom HTTP client to use for communicating with the
	// metadata service.
	HTTPClient HTTPClient

	// ProjectAccessKey is the access key for the project. (optional)
	ProjectAccessKey string
}

const DefaultMetadataServiceURL = "https://metadata.sequence.app"

// NewClient creates a new Sequence Metadata client instance. Please see https://sequence.build to
// get a `projectAccessKey`, which is your project's access key used to communicate
// with Sequence services.
//
// NOTE: the `projectAccessKey` may be optional if you're using a JWT auth token
// passed in via the `clientOptions`.
func NewClient(projectAccessKey string, clientOptions ...Options) MetadataClient {
	opts := Options{}
	if len(clientOptions) > 0 {
		opts = clientOptions[0]
	}
	if opts.ProjectAccessKey != "" && projectAccessKey == "" {
		projectAccessKey = opts.ProjectAccessKey
	}

	client := &httpClient{
		client:           opts.HTTPClient,
		projectAccessKey: projectAccessKey,
	}
	if opts.HTTPClient == nil {
		client.client = http.DefaultClient
	}
	if opts.JWTAuthToken != "" {
		client.jwtAuthHeader = fmt.Sprintf("BEARER %s", opts.JWTAuthToken)
	}

	serviceURL := DefaultMetadataServiceURL
	if opts.MetadataServiceURL != "" {
		serviceURL = opts.MetadataServiceURL
	}
	return NewMetadataClient(strings.TrimSuffix(serviceURL, "/"), client)
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
