package proto

import (
	"fmt"
	"net/http"

	"github.com/0xsequence/go-sequence"
)

type Options struct {
	HTTPClient   HTTPClient
	JWTAuthToken string
}

// NewRpcRelayer creates a new Sequence Relayer client instance. See https://docs.sequence.xyz for a list of
// relayer urls, and please see https://sequence.build to get a `projectAccessKey`.
func NewRelayer(relayerServiceURL string, projectAccessKey string, options ...Options) Relayer {
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

	return NewRelayerClient(relayerServiceURL, client)
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

func MetaTxnStatusFromString(s string) sequence.MetaTxnStatus {
	var ethTxnStatus ETHTxnStatus
	ethTxnStatus.UnmarshalJSON([]byte(s))

	switch ethTxnStatus {
	case ETHTxnStatus_UNKNOWN:
		return sequence.MetaTxnStatusUnknown
	case ETHTxnStatus_DROPPED:
		return sequence.MetaTxnStatusUnknown
	case ETHTxnStatus_QUEUED:
		return sequence.MetaTxnStatusUnknown
	case ETHTxnStatus_SENT:
		return sequence.MetaTxnStatusUnknown
	case ETHTxnStatus_SUCCEEDED:
		return sequence.MetaTxnExecuted
	case ETHTxnStatus_PARTIALLY_FAILED:
		return sequence.MetaTxnFailed
	case ETHTxnStatus_FAILED:
		return sequence.MetaTxnFailed
	default:
		return sequence.MetaTxnStatusUnknown
	}
}
