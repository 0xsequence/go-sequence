package metadata

import (
	"fmt"
	"net/http"
)

type Options struct {
	HTTPClient         HTTPClient
	JWTAuthToken       string
	MetadataServiceURL string
}

// NewMetadata creates a new Sequence Metadata client instance. Please see
// https://sequence.build to get a `projectAccessKey`.
func NewMetadata(projectAccessKey string, options ...Options) MetadataClient {
	opts := Options{}
	if len(options) > 0 {
		opts = options[0]
	}

	client := &httpclient{
		client:           opts.HTTPClient,
		projectAccessKey: projectAccessKey,
	}
	if opts.HTTPClient == nil {
		client.client = http.DefaultClient
	}
	if opts.JWTAuthToken != "" {
		client.jwtAuthHeader = fmt.Sprintf("BEARER %s", opts.JWTAuthToken)
	}

	metadataServiceURL := "https://metadata.sequence.app"
	if opts.MetadataServiceURL != "" {
		metadataServiceURL = opts.MetadataServiceURL
	}
	return NewMetadataClient(metadataServiceURL, client)
}

type httpclient struct {
	client           HTTPClient
	jwtAuthHeader    string
	projectAccessKey string
}

func (c *httpclient) Do(req *http.Request) (*http.Response, error) {
	if c.projectAccessKey != "" {
		req.Header.Set("X-Access-Key", c.projectAccessKey)
	}
	if c.jwtAuthHeader != "" {
		req.Header.Set("Authorization", c.jwtAuthHeader)
	}
	return c.client.Do(req)
}
