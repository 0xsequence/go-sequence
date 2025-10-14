package indexer

import (
	"fmt"
	"net/http"
	"strings"
)

type GatewayOptions struct {
	// JWTAuthToken is an optional JWT token to authenticate with the indexer
	// gateway service.
	JWTAuthToken string

	// IndexerGatewayServiceURL is an optional custom URL for the Sequence Indexer Gateway
	// service. If not provided, the default URL in `DefaultIndexerGatewayServiceURL`
	// will be used.
	IndexerGatewayServiceURL string

	// HTTPClient is an optional custom HTTP client to use for communicating with the
	// indexer gateway service.
	HTTPClient HTTPClient
}

const DefaultIndexerGatewayServiceURL = "https://indexer.sequence.app"

func NewGatewayClient(projectAccessKey string, clientOptions ...GatewayOptions) IndexerGatewayClient {
	opts := GatewayOptions{}
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

	serviceURL := DefaultIndexerGatewayServiceURL
	if opts.IndexerGatewayServiceURL != "" {
		serviceURL = opts.IndexerGatewayServiceURL
	}
	return NewIndexerGatewayClient(strings.TrimSuffix(serviceURL, "/"), c)

}
