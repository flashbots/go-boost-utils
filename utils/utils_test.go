package utils

import (
	"fmt"
	"github.com/attestantio/go-eth2-client/spec/electra"
	"github.com/golang/snappy"
	"io/ioutil"
	"net/http"
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

const mainnetTests = "https://github.com/ethereum/consensus-spec-tests/raw/master/tests/mainnet/"

func downloadFile(url string) ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to download file: %s", resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

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
	testCases := []struct {
		name        string
		pubkey      string
		expectedErr string
	}{
		{
			name:   "Valid pubkey",
			pubkey: "0xac6e77dfe25ecd6110b8e780608cce0dab71fdd5ebea22a16c0205200f2f8e2e3ad3b71d3499c54ad14d6c21b41a37ae",
		},
		{
			name:        "Invalid pubkey (wrong length)",
			pubkey:      "0x123456",
			expectedErr: "invalid length",
		},
		{
			name:        "Invalid pubkey (not on the curve)",
			pubkey:      "0xed7f862045422bd51ba732730ce993c94d2545e5db1112102026343904fcdf6f5cf37926a3688444703772ed80fa223f",
			expectedErr: "invalid pubkey",
		},
		{
			name:        "Invalid pubkey (no 0x prefix)",
			pubkey:      "this is not hex",
			expectedErr: "hex string without 0x prefix",
		},
		{
			name:        "Invalid pubkey (not hex)",
			pubkey:      "0xthisisnothex",
			expectedErr: "invalid hex string",
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			result, err := HexToPubkey(tt.pubkey)
			if tt.expectedErr != "" {
				require.EqualError(t, err, tt.expectedErr)
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.pubkey, result.String())
			}
		})
	}
}

func TestHexToSignature(t *testing.T) {
	testCases := []struct {
		name        string
		signature   string
		expectedErr string
	}{
		{
			name:      "Valid signature",
			signature: "0x8069aa021666163aae46d353c348aa913fb5050062e05ab764c8eef99407a8befcd46190b30cc40e40ee8c197356959816799d62e85f640ef76b4be1b08a741949230fbde49589125537daad06c23a66838725d89e3504bc21559a91534f6712",
		},
		{
			name:        "Invalid signature (wrong length)",
			signature:   "0x123456",
			expectedErr: "invalid length",
		},
		{
			name:        "Invalid signature (not on the curve)",
			signature:   "0xb8f03e639b91fa8e9892f66c798f07f6e7b3453234f643b2c06a35c5149cf6d85e4e1572c33549fe749292445fbff9e0739c78159324c35dc1a90e5745ca70c8caf1b63fb6678d81bd2d5cb6baeb1462df7a93877d0e22a31dd6438334536d9a",
			expectedErr: "invalid signature",
		},
		{
			name:        "Invalid signature (no 0x prefix)",
			signature:   "this is not hex",
			expectedErr: "hex string without 0x prefix",
		},
		{
			name:        "Invalid signature (not hex)",
			signature:   "0xthisisnothex",
			expectedErr: "invalid hex string",
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			result, err := HexToSignature(tt.signature)
			if tt.expectedErr != "" {
				require.EqualError(t, err, tt.expectedErr)
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.signature, result.String())
			}
		})
	}
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

	t.Run("compute deneb block hash", func(t *testing.T) {
		body := func() (body deneb.BeaconBlockBody) {
			contentURL := mainnetTests + "deneb/operations/execution_payload/pyspec_tests/success_regular_payload/body.ssz_snappy"
			compressedData, err := downloadFile(contentURL)
			require.NoError(t, err)
			decompressedData, err := snappy.Decode(nil, compressedData)
			require.NoError(t, err)
			err = body.UnmarshalSSZ(decompressedData)
			require.NoError(t, err)
			return
		}()

		versionedPayload := &api.VersionedExecutionPayload{
			Version: spec.DataVersionDeneb,
			Deneb:   body.ExecutionPayload,
		}

		hash, err := ComputeBlockHash(versionedPayload, nil)
		require.NoError(t, err)
		expectedHash, err := versionedPayload.BlockHash()
		require.NoError(t, err)
		require.Equal(t, expectedHash, hash)
	})

	t.Run("compute electra block hash", func(t *testing.T) {
		body := func() (body electra.BeaconBlockBody) {
			contentURL := mainnetTests + "electra/operations/execution_payload/pyspec_tests/success_regular_payload/body.ssz_snappy"
			compressedData, err := downloadFile(contentURL)
			require.NoError(t, err)
			decompressedData, err := snappy.Decode(nil, compressedData)
			require.NoError(t, err)
			err = body.UnmarshalSSZ(decompressedData)
			require.NoError(t, err)
			return
		}()

		versionedPayload := &api.VersionedExecutionPayload{
			Version: spec.DataVersionElectra,
			Electra: body.ExecutionPayload,
		}

		hash, err := ComputeBlockHash(versionedPayload, nil)
		require.NoError(t, err)
		expectedHash, err := versionedPayload.BlockHash()
		require.NoError(t, err)
		require.Equal(t, expectedHash, hash)
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
