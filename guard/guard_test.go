package guard

import (
	"context"
	"math/big"
	"testing"

	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/0xsequence/go-sequence"
	"github.com/0xsequence/go-sequence/signing_service"
	proto_signing_service "github.com/0xsequence/go-sequence/signing_service/proto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

type MockSigningServiceClient struct {
	mock.Mock
}

func (m *MockSigningServiceClient) Ping(ctx context.Context) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MockSigningServiceClient) Version(ctx context.Context) (*proto_signing_service.Version, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MockSigningServiceClient) RuntimeStatus(ctx context.Context) (*proto_signing_service.RuntimeStatus, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MockSigningServiceClient) GetSignerConfig(ctx context.Context, signer string) (*proto_signing_service.WalletConfig, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MockSigningServiceClient) Sign(ctx context.Context, request *proto_signing_service.SignRequest) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MockSigningServiceClient) SignWith(ctx context.Context, signer string, request *proto_signing_service.SignRequest) (string, error) {
	args := m.Called(ctx, signer, request)
	return args.String(0), args.Error(1)
}

var _ proto_signing_service.SigningService = &MockSigningServiceClient{}

func TestGuardSigningServiceSign(t *testing.T) {
	mockSigningServiceClient := &MockSigningServiceClient{}

	guardSigningService := NewGuardSigningService(GuardSigningServiceParams{
		SigningServiceParams: signing_service.SigningServiceParams{
			SignerAddress: common.Address{},
			SignerWeight:  0,
			Client:        mockSigningServiceClient,
		},
	})

	mockSigningServiceClient.On("SignWith", mock.Anything, mock.Anything, mock.Anything).Return("0x01", nil)

	ctx := signing_service.ContextWithSignContext(context.Background(), &proto_signing_service.SignContext{
		Signature: "0x01",
	})

	msg := []byte("hello world")
	digest := sequence.MessageDigest(msg)

	sigBytes, _, err := guardSigningService.SignDigest(ctx, digest, big.NewInt(1))
	require.NoError(t, err)
	assert.NotEmpty(t, sigBytes)
}

func TestGuardSigningServiceSignNoSignContext(t *testing.T) {
	mockSigningServiceClient := &MockSigningServiceClient{}

	guardSigningService := NewGuardSigningService(GuardSigningServiceParams{
		SigningServiceParams: signing_service.SigningServiceParams{
			SignerAddress: common.Address{},
			SignerWeight:  0,
			Client:        mockSigningServiceClient,
		},
	})

	msg := []byte("hello world")
	digest := sequence.MessageDigest(msg)

	sigBytes, sigTyped, err := guardSigningService.SignDigest(context.Background(), digest, big.NewInt(1))
	require.Error(t, err)
	assert.Nil(t, sigBytes)
	assert.Nil(t, sigTyped)
}
