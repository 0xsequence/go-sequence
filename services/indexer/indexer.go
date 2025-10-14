package indexer

import (
	"fmt"
	"net/http"
	"strings"
)

type Options struct {
	// JWTAuthToken is an optional JWT token to authenticate with the indexer service.
	JWTAuthToken string

	// HTTPClient is an optional custom HTTP client to use for communicating with the
	// indexer service.
	HTTPClient HTTPClient
}

// NewClient creates a new Sequence Indexer client instance for a specific chain.
// Please see https://sequence.build to get a `projectAccessKey`, which is your
// project's access key used to communicate with Sequence services.
//
// NOTE: the `projectAccessKey` may be optional if you're using a JWT auth token
// passed in via the `clientOptions`.
//
// The `indexerURL` is the URL of the indexer service to connect to, for example:
// https://mainnet-indexer.sequence.app for Ethereum mainnet and https://polygon-indexer.sequence.app
// for Polygon mainnet. See https://docs.sequence.xyz for a complete list of indexer urls.
//
// Additionally, you may be interested in the `NewGatewayClient` which is a single
// client that connects to multiple indexers for many different chains at once.
//
// Finally, you may pass in optional `clientOptions` to configure the indexer client
// with jwt-based authentication, a receipts listener, and a custom HTTP client.
func NewClient(indexerURL string, projectAccessKey string, clientOptions ...Options) IndexerClient {
	opts := Options{}
	if len(clientOptions) > 0 {
		opts = clientOptions[0]
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

	return NewIndexerClient(strings.TrimSuffix(indexerURL, "/"), client)
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
