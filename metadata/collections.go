package metadata

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
)

type CollectionsService struct {
	Collections
	options    Options
	httpclient HTTPClient
}

// NewCollections creates a new Sequence Metadata Collections client instance. Please see
// https://sequence.build to get a `projectServiceJWTToken` service-level account jwt token.
func NewCollections(projectServiceJWTToken string, options ...Options) CollectionsService {
	opts := Options{}
	if len(options) > 0 {
		opts = options[0]
	}

	client := &httpclient{client: opts.HTTPClient}
	if opts.HTTPClient == nil {
		client.client = http.DefaultClient
	}
	if opts.JWTAuthToken == "" {
		opts.JWTAuthToken = projectServiceJWTToken
	}
	client.jwtAuthHeader = fmt.Sprintf("BEARER %s", opts.JWTAuthToken)

	metadataServiceURL := "https://metadata.sequence.app"
	if opts.MetadataServiceURL != "" {
		metadataServiceURL = opts.MetadataServiceURL
	} else {
		opts.MetadataServiceURL = metadataServiceURL
	}

	serviceClient := NewCollectionsClient(metadataServiceURL, client)

	return CollectionsService{
		Collections: serviceClient,
		options:     opts,
		httpclient:  client,
	}
}

func (c *CollectionsService) UploadAsset(ctx context.Context, projectID, collectionID, assetID uint64, assetContent io.Reader) (*Asset, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	fileWriter, err := writer.CreateFormFile("file", "asset")
	if err != nil {
		return nil, fmt.Errorf("create form file: %w", err)
	}

	_, err = io.Copy(fileWriter, assetContent)
	if err != nil {
		return nil, fmt.Errorf("copy bytes: %w", err)
	}

	writer.Close()

	endpointURL := fmt.Sprintf("%s/projects/%d/collections/%d/assets/%d/upload", c.options.MetadataServiceURL, projectID, collectionID, assetID)
	req, err := http.NewRequestWithContext(ctx, http.MethodPut, endpointURL, body)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())

	resp, err := c.httpclient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("upload asset failed: %s", resp.Status)
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read body: %w", err)

	}

	asset := &Asset{}
	err = json.Unmarshal(respBody, asset)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return asset, nil
}
