package relayer

import (
	"net/http"

	"github.com/0xsequence/go-sequence"
	"github.com/0xsequence/go-sequence/relayer/proto"
)

type httpClient struct {
	client           proto.HTTPClient
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
