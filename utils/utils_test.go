package utils

import (
	"os"
	"testing"

	"github.com/attestantio/go-builder-client/api"
	"github.com/attestantio/go-eth2-client/spec"
	"github.com/attestantio/go-eth2-client/spec/bellatrix"
	"github.com/attestantio/go-eth2-client/spec/capella"
	"github.com/attestantio/go-eth2-client/spec/deneb"
	"github.com/attestantio/go-eth2-client/spec/phase0"
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
	t.Run("Should compute bellatrix hash", func(t *testing.T) {
		jsonFile, err := os.Open("../testdata/executionpayload/bellatrix-case0.json")
		require.NoError(t, err)
		defer jsonFile.Close()

		payload := new(bellatrix.ExecutionPayload)
		require.NoError(t, DecodeJSON(jsonFile, payload))
		versionedPayload := &api.VersionedExecutionPayload{
			Version:   spec.DataVersionBellatrix,
			Bellatrix: payload,
		}

		hash, err := ComputeBlockHash(versionedPayload, nil)
		require.NoError(t, err)
		require.Equal(t, "0x6662fb418aa7b5c5c80e2e8bc87be48db82e799c4704368d34ddeb3b12549655", hash.String())
	})

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

		hash, err := ComputeBlockHash(versionedPayload, nil)
		require.NoError(t, err)
		require.Equal(t, "0x08751ea2076d3ecc606231495a90ba91a66a9b8fb1a2b76c333f1957a1c667c3", hash.String())
	})

	t.Run("Should compute deneb hash", func(t *testing.T) {
		jsonFile, err := os.Open("../testdata/executionpayload/deneb-case0.json")
		require.NoError(t, err)
		defer jsonFile.Close()

		payload := new(deneb.ExecutionPayload)
		require.NoError(t, DecodeJSON(jsonFile, payload))
		versionedPayload := &api.VersionedExecutionPayload{
			Version: spec.DataVersionDeneb,
			Deneb:   payload,
		}
		h, _ := HexToHash("0xa119064ee9c03e2c7ad5821b6077606c64f36542eda12ed61a1edc5f898a17fc")
		r := phase0.Root(h)
		hash, err := ComputeBlockHash(versionedPayload, &r)
		require.NoError(t, err)
		require.Equal(t, "0xd9491c8ae79611d0f08806f29b1e2e86cb8f64512aa381e543dcae257dda80d6", hash.String())
	})

	t.Run("Should error on unknown version", func(t *testing.T) {
		payload := new(capella.ExecutionPayload)
		versionedPayload := &api.VersionedExecutionPayload{
			Version: spec.DataVersionAltair,
			Capella: payload,
		}

		_, err := ComputeBlockHash(versionedPayload, nil)
		require.Error(t, err)
	})
}
