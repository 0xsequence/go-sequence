package relayer

import (
	"fmt"
	"net/http"

	"github.com/0xsequence/go-sequence"
	"github.com/0xsequence/go-sequence/relayer/proto"
)

type Options struct {
	HTTPClient   proto.HTTPClient
	JWTAuthToken string
}

// NewRelayer creates a new Sequence Relayer client instance. See https://docs.sequence.xyz for a list of
// relayer urls, and please see https://sequence.build to get a `projectAccessKey`.
func NewRelayer(relayerURL string, projectAccessKey string, options ...Options) proto.RelayerClient {
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

	return proto.NewRelayerClient(relayerURL, client)
}

type httpclient struct {
	client           proto.HTTPClient
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
	var ethTxnStatus proto.ETHTxnStatus
	ethTxnStatus.UnmarshalText([]byte(s))

	switch ethTxnStatus {
	case proto.ETHTxnStatus_UNKNOWN:
		return sequence.MetaTxnStatusUnknown
	case proto.ETHTxnStatus_DROPPED:
		return sequence.MetaTxnStatusUnknown
	case proto.ETHTxnStatus_QUEUED:
		return sequence.MetaTxnStatusUnknown
	case proto.ETHTxnStatus_SENT:
		return sequence.MetaTxnStatusUnknown
	case proto.ETHTxnStatus_SUCCEEDED:
		return sequence.MetaTxnExecuted
	case proto.ETHTxnStatus_PARTIALLY_FAILED:
		return sequence.MetaTxnFailed
	case proto.ETHTxnStatus_FAILED:
		return sequence.MetaTxnFailed
	default:
		return sequence.MetaTxnStatusUnknown
	}
}
