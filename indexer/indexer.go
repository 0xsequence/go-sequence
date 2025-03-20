package indexer

import (
	"fmt"
	"net/http"
)

//go:generate go run go.uber.org/mock/mockgen -destination indexer.mock.go -package indexer . Indexer

type Options struct {
	HTTPClient   HTTPClient
	JWTAuthToken string
}

// NewIndexer creates a new Sequence Indexer client instance. See https://docs.sequence.xyz for a list of
// indexer urls, and please see https://sequence.build to get a `projectAccessKey`.
func NewIndexer(indexerURL string, projectAccessKey string, options ...Options) IndexerClient {
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

	return NewIndexerClient(indexerURL, client)
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
