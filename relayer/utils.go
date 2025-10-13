package relayer

import (
	"math/big"
	"net/http"

	"github.com/0xsequence/ethkit"
	"github.com/0xsequence/ethkit/go-ethereum/common"
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

func convFeeTokenToRelayerFeeToken(token *proto.FeeToken) sequence.RelayerFeeToken {
	var contractAddress *common.Address
	if token.ContractAddress != nil {
		contractAddress = ethkit.ToPtr(common.HexToAddress(*token.ContractAddress))
	}

	return sequence.RelayerFeeToken{
		ChainID:         big.NewInt(0).SetUint64(token.ChainId),
		Name:            token.Name,
		Symbol:          token.Symbol,
		Type:            sequence.RelayerFeeTokenType(token.Type),
		Decimals:        token.Decimals,
		LogoURL:         token.LogoURL,
		ContractAddress: contractAddress,
	}
}

func convFeeOptionToRelayerFeeOption(option *proto.FeeOption) *sequence.RelayerFeeOption {
	value, _ := big.NewInt(0).SetString(option.Value, 10)
	return &sequence.RelayerFeeOption{
		Token:    convFeeTokenToRelayerFeeToken(option.Token),
		To:       common.HexToAddress(option.To),
		Value:    value,
		GasLimit: big.NewInt(0).SetUint64(uint64(option.GasLimit)),
	}
}

func convFeeOptionsToRelayerFeeOptions(options []*proto.FeeOption) []*sequence.RelayerFeeOption {
	var feeOptions []*sequence.RelayerFeeOption
	for _, option := range options {
		feeOptions = append(feeOptions, convFeeOptionToRelayerFeeOption(option))
	}
	return feeOptions
}
