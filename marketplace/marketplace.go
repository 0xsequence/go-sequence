package marketplace

import (
	"fmt"
	"net/http"
)

type Options struct {
	MarketplaceAPIURL string
	JWTAuthToken      string
	HTTPClient        HTTPClient
}

func NewMarketplaceAdmin(projectAccessKey string, options ...Options) AdminClient {
	opts := Options{}
	if len(options) > 0 {
		opts = options[0]
	}

	client := &httpclient{
		client:           opts.HTTPClient,
		projectAccessKey: projectAccessKey,
	}

	if client.client == nil {
		client.client = http.DefaultClient
	}

	if opts.JWTAuthToken != "" {
		client.jwtAuthHeader = fmt.Sprintf("BEARER %s", opts.JWTAuthToken)
	}

	// prod: https://marketplace-api.sequence.app
	// dev: https://dev-marketplace-api.sequence.app
	apiURL := fmt.Sprintf("https://marketplace-api.sequence.app/")
	if opts.MarketplaceAPIURL != "" {
		apiURL = opts.MarketplaceAPIURL
	}

	return NewAdminClient(apiURL, client)
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
