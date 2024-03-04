package metadata

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"path/filepath"
)

type CollectionsService struct {
	Collections
	options          Options
	projectAccessKey string
	httpclient       HTTPClient
}

// NewCollections creates a new Sequence Metadata Collections client instance. Please see
// https://sequence.build to get a `projectAccessKey` and service-level account JWTAuthToken.
func NewCollections(projectAccessKey string, options ...Options) CollectionsService {
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
	} else {
		opts.MetadataServiceURL = metadataServiceURL
	}

	serviceClient := NewCollectionsClient(metadataServiceURL, client)

	return CollectionsService{
		Collections:      serviceClient,
		options:          opts,
		projectAccessKey: projectAccessKey,
		httpclient:       client,
	}
}

func (c *CollectionsService) UploadAsset(projectID, collectionID, tokenID, assetFilename string, fileContent io.Reader) (*http.Response, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	fileWriter, err := writer.CreateFormFile("file", filepath.Base(assetFilename))
	if err != nil {
		return nil, fmt.Errorf("create form file: %w", err)
	}

	_, err = io.Copy(fileWriter, fileContent)
	if err != nil {
		return nil, fmt.Errorf("copy bytes: %w", err)
	}

	writer.Close()

	endpointURL := fmt.Sprintf("%s/collections/%s/%s/%s/upload/%s", c.options.MetadataServiceURL, projectID, collectionID, tokenID, assetFilename)
	req, err := http.NewRequest(http.MethodPut, endpointURL, body)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())

	resp, err := c.httpclient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do: %w", err)
	}

	return resp, nil
}
