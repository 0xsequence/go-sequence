package marketplace

import (
	"cmp"
	"fmt"
	"net/http"
	"strings"
)

type Options struct {
	MarketplaceAPIURL string
	JWTAuthToken      string
	HTTPClient        HTTPClient
}

func NewMarketplaceAdmin(projectAccessKey string, options ...Options) AdminClient {
	return NewAdminClient(newMarketplaceClient(projectAccessKey, options...))
}

func NewMarketplaceMarketplace(projectAccessKey string, options ...Options) MarketplaceClient {
	return NewMarketplaceClient(newMarketplaceClient(projectAccessKey, options...))
}

func newMarketplaceClient(projectAccessKey string, options ...Options) (string, *httpClient) {
	o := Options{}
	if len(options) > 0 {
		o = options[0]
	}

	c := &httpClient{
		client:           cmp.Or[HTTPClient](o.HTTPClient, http.DefaultClient),
		projectAccessKey: projectAccessKey,
	}

	if o.JWTAuthToken != "" {
		c.jwtAuthHeader = fmt.Sprintf("BEARER %s", o.JWTAuthToken)
	}

	apiURL := cmp.Or(o.MarketplaceAPIURL, "https://marketplace-api.sequence.app")

	return strings.TrimSuffix(apiURL, "/"), c
}

type httpClient struct {
	client           HTTPClient
	projectAccessKey string
	jwtAuthHeader    string
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
