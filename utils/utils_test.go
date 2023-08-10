package utils

import (
	"os"
	"testing"

	"github.com/attestantio/go-builder-client/api"
	"github.com/attestantio/go-eth2-client/spec"
	"github.com/attestantio/go-eth2-client/spec/capella"
	"github.com/stretchr/testify/require"
)

func TestHexToHash(t *testing.T) {
	_, err := HexToHash("0x01")
	require.Error(t, err)

	a, err := HexToHash("0x0100000000000000000000000000000000000000000000000000000000000000")
	require.NoError(t, err)
	require.Equal(t, "0x0100000000000000000000000000000000000000000000000000000000000000", a.String())
}

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

func TestComputeHash(t *testing.T) {
	t.Run("Should compute capella hash", func(t *testing.T) {
		jsonFile, err := os.Open("../testdata/executionpayload/capella-case0.json")
		require.NoError(t, err)
		defer jsonFile.Close()

		payload := new(capella.ExecutionPayload)
		require.NoError(t, DecodeJSON(jsonFile, payload))
		versionedPayload := &api.VersionedExecutionPayload{
			Version: spec.DataVersionCapella,
			Capella: payload,
		}

		hash, err := ComputeBlockHash(versionedPayload)
		require.NoError(t, err)
		require.Equal(t, "0x08751ea2076d3ecc606231495a90ba91a66a9b8fb1a2b76c333f1957a1c667c3", hash.String())
	})

	// 4844 deneb execution payload and header field still unstable to write tests for
	// t.Run("Should compute deneb hash", func(t *testing.T) {
	// 	jsonFile, err := os.Open("../testdata/executionpayload/deneb-case0.json")
	// 	require.NoError(t, err)
	// 	defer jsonFile.Close()

	// 	payload := new(deneb.ExecutionPayload)
	// 	require.NoError(t, DecodeJSON(jsonFile, payload))
	// 	versionedPayload := &api.VersionedExecutionPayload{
	// 		Version: spec.DataVersionDeneb,
	// 		Deneb:   payload,
	// 	}

	// 	hash, err := ComputeBlockHash(versionedPayload)
	// 	require.NoError(t, err)
	// 	require.Equal(t, "0x6fdbdbce83765e4196c1cc104aaad70385bdda458c11703a7bc1495504191fa0", hash.String())
	// })

	t.Run("Should error on unknown version", func(t *testing.T) {
		payload := new(capella.ExecutionPayload)
		versionedPayload := &api.VersionedExecutionPayload{
			Version: spec.DataVersionAltair,
			Capella: payload,
		}

		_, err := ComputeBlockHash(versionedPayload)
		require.Error(t, err)
	})
}
