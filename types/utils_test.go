package types

import (
	"testing"

	"github.com/stretchr/testify/require"
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
