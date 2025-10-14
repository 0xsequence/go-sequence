package marketplace

import (
	"cmp"
	"fmt"
	"net/http"
	"strings"
)

type Options struct {
	// JWTAuthToken is an optional JWT token to authenticate with the marketplace service.
	JWTAuthToken string

	// MarketplaceServiceURL is an optional custom URL for the Sequence Marketplace service.
	// If not provided, the default URL in `DefaultMarketplaceServiceURL` will be used.
	MarketplaceServiceURL string

	// HTTPClient is an optional custom HTTP client to use for communicating with the
	// marketplace service.
	HTTPClient HTTPClient
}

const DefaultMarketplaceServiceURL = "https://marketplace-api.sequence.app"

func NewClient(projectAccessKey string, clientOptions ...Options) MarketplaceClient {
	return NewMarketplaceClient(newMarketplaceClient(projectAccessKey, clientOptions...))
}

func NewMarketplaceAdminClient(projectAccessKey string, clientOptions ...Options) AdminClient {
	return NewAdminClient(newMarketplaceClient(projectAccessKey, clientOptions...))
}

func newMarketplaceClient(projectAccessKey string, clientOptions ...Options) (string, *httpClient) {
	o := Options{}
	if len(clientOptions) > 0 {
		o = clientOptions[0]
	}

	c := &httpClient{
		client:           cmp.Or[HTTPClient](o.HTTPClient, http.DefaultClient),
		projectAccessKey: projectAccessKey,
	}
	if o.JWTAuthToken != "" {
		c.jwtAuthHeader = fmt.Sprintf("BEARER %s", o.JWTAuthToken)
	}

	serviceURL := cmp.Or(o.MarketplaceServiceURL, DefaultMarketplaceServiceURL)

	return strings.TrimSuffix(serviceURL, "/"), c
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
