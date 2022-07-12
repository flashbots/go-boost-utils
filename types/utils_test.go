package types

import (
	"testing"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/flashbots/go-boost-utils/bls"
	"github.com/stretchr/testify/require"
)

var (
	builderSigningDomain = Domain([32]byte{0, 0, 0, 1, 245, 165, 253, 66, 209, 106, 32, 48, 39, 152, 239, 110, 211, 9, 151, 155, 67, 0, 61, 35, 32, 217, 240, 232, 234, 152, 49, 169})
)

func TestHexToAddress(t *testing.T) {
	_, err := HexToAddress("0x01")
	require.Error(t, err)

	a, err := HexToAddress("0x0100000000000000000000000000000000000000")
	require.NoError(t, err)
	require.Equal(t, "0x0100000000000000000000000000000000000000", a.String())
}

func TestHexToPubkey(t *testing.T) {
	_, err := HexToPubkey("0x01")
	require.Error(t, err)

	a, err := HexToPubkey("0xed7f862045422bd51ba732730ce993c94d2545e5db1112102026343904fcdf6f5cf37926a3688444703772ed80fa223f")
	require.NoError(t, err)
	require.Equal(t, "0xed7f862045422bd51ba732730ce993c94d2545e5db1112102026343904fcdf6f5cf37926a3688444703772ed80fa223f", a.String())
}

func TestHexToSignature(t *testing.T) {
	_, err := HexToSignature("0x01")
	require.Error(t, err)

	a, err := HexToSignature("0xb8f03e639b91fa8e9892f66c798f07f6e7b3453234f643b2c06a35c5149cf6d85e4e1572c33549fe749292445fbff9e0739c78159324c35dc1a90e5745ca70c8caf1b63fb6678d81bd2d5cb6baeb1462df7a93877d0e22a31dd6438334536d9a")
	require.NoError(t, err)
	require.Equal(t, "0xb8f03e639b91fa8e9892f66c798f07f6e7b3453234f643b2c06a35c5149cf6d85e4e1572c33549fe749292445fbff9e0739c78159324c35dc1a90e5745ca70c8caf1b63fb6678d81bd2d5cb6baeb1462df7a93877d0e22a31dd6438334536d9a", a.String())
}

func TestBuilderBlockRequestToSignedBuilderBid(t *testing.T) {
	builderPk, err := HexToPubkey("0xf9716c94aab536227804e859d15207aa7eaaacd839f39dcbdb5adc942842a8d2fb730f9f49fc719fdb86f1873e0ed1c2")
	require.NoError(t, err)

	builderSk, err := HexToSignature("0x8209b5391cd69f392b1f02dbc03bab61f574bb6bb54bf87b59e2a85bdc0756f7db6a71ce1b41b727a1f46ccc77b213bf0df1426177b5b29926b39956114421eaa36ec4602969f6f6370a44de44a6bce6dae2136e5fb594cce2a476354264d1ea")
	require.NoError(t, err)

	reqPayload := BuilderSubmitBlockRequest{
		ExecutionPayload: ExecutionPayload{
			ParentHash:    Hash{0x01},
			FeeRecipient:  Address{0x02},
			StateRoot:     Root{0x03},
			ReceiptsRoot:  Root{0x04},
			LogsBloom:     Bloom{0x05},
			Random:        Hash{0x06},
			BlockNumber:   5001,
			GasLimit:      5002,
			GasUsed:       5003,
			Timestamp:     5004,
			ExtraData:     []byte{0x07},
			BaseFeePerGas: IntToU256(123),
			BlockHash:     Hash{0x09},
			Transactions:  []hexutil.Bytes{},
		},
		Message: BidTraceMessage{
			Slot:                 1,
			ParentHash:           Hash{0x01},
			BlockHash:            Hash{0x09},
			BuilderPubkey:        builderPk,
			ProposerPubkey:       PublicKey{0x03},
			ProposerFeeRecipient: Address{0x04},
			Value:                IntToU256(123),
		},
		Signature: builderSk,
	}

	sk, _, err := bls.GenerateNewKeypair()
	require.NoError(t, err)
	signedBuilderBid, err := BuilderSubmitBlockRequestToSignedBuilderBid(&reqPayload, sk, builderSigningDomain)
	require.NoError(t, err)

	require.Equal(t, 0, signedBuilderBid.Message.Value.Cmp(&reqPayload.Message.Value))
	require.Equal(t, reqPayload.Message.BlockHash, signedBuilderBid.Message.Header.BlockHash)
}
