package metadata

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"path/filepath"
)

type CollectionRest struct {
	httpClient  *httpclient
	metadataURL string
}

// NewCollectionREST creates a new Sequence Metadata client instance for REST Collections endpoints.
func NewCollectionREST(projectAccessKey string, options ...Options) CollectionRest {
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

	return CollectionRest{
		httpClient:  client,
		metadataURL: metadataServiceURL,
	}
}

func (c *CollectionRest) AssetUpload(projectID, collectionID, tokenID, assetID string, fileContent io.Reader) (*http.Response, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	fileWriter, err := writer.CreateFormFile("file", filepath.Base(assetID))
	if err != nil {
		return nil, fmt.Errorf("create form file: %w", err)
	}

	_, err = io.Copy(fileWriter, fileContent)
	if err != nil {
		return nil, fmt.Errorf("copy bytes: %w", err)
	}

	writer.Close()

	endpointURL := fmt.Sprintf("%s/collections/%s/%s/%s/upload/%s", c.metadataURL, projectID, collectionID, tokenID, assetID)
	req, err := http.NewRequest(http.MethodPut, endpointURL, body)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())

	resp, err := c.httpClient.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do: %w", err)
	}

	return resp, nil
}
