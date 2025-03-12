package relayer

import (
	"fmt"
	"net/http"

	"github.com/0xsequence/go-sequence/core/v1v2"
	"github.com/0xsequence/go-sequence/core/v1v2/relayer/proto"
)

type Options struct {
	HTTPClient   proto.HTTPClient
	JWTAuthToken string
}

// NewRpcRelayer creates a new Sequence Relayer client instance. See https://docs.sequence.xyz for a list of
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

func MetaTxnStatusFromString(s string) v1v2.MetaTxnStatus {
	var ethTxnStatus proto.ETHTxnStatus
	ethTxnStatus.UnmarshalText([]byte(s))

	switch ethTxnStatus {
	case proto.ETHTxnStatus_UNKNOWN:
		return v1v2.MetaTxnStatusUnknown
	case proto.ETHTxnStatus_DROPPED:
		return v1v2.MetaTxnStatusUnknown
	case proto.ETHTxnStatus_QUEUED:
		return v1v2.MetaTxnStatusUnknown
	case proto.ETHTxnStatus_SENT:
		return v1v2.MetaTxnStatusUnknown
	case proto.ETHTxnStatus_SUCCEEDED:
		return v1v2.MetaTxnExecuted
	case proto.ETHTxnStatus_PARTIALLY_FAILED:
		return v1v2.MetaTxnFailed
	case proto.ETHTxnStatus_FAILED:
		return v1v2.MetaTxnFailed
	default:
		return v1v2.MetaTxnStatusUnknown
	}
}
