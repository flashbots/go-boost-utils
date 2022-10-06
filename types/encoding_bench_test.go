package types

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func BenchmarkJSONvsSSZEncoding(b *testing.B) {
	var err error
	jsonFile, err := os.Open("../testdata/signed-blinded-beacon-block-with-deposit.json")
	require.NoError(b, err)
	defer jsonFile.Close()
	signedBlindedBeaconBlock := new(SignedBlindedBeaconBlock)
	require.NoError(b, DecodeJSON(jsonFile, &signedBlindedBeaconBlock))
	signedBlindedBeaconBlockJSONBytes, err := json.Marshal(signedBlindedBeaconBlock)
	require.NoError(b, err)
	signedBlindedBeaconBlockSSZBytes, err := signedBlindedBeaconBlock.MarshalSSZ()
	require.NoError(b, err)

	b.Run("Encode JSON", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			_, err = json.Marshal(signedBlindedBeaconBlock)
			require.NoError(b, err)
		}
	})

	b.Run("Encode SSZ", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			_, err = signedBlindedBeaconBlock.MarshalSSZ()
			require.NoError(b, err)
		}
	})

	b.Run("Decode JSON", func(b *testing.B) {
		_blindedBeaconBlock := new(SignedBlindedBeaconBlock)
		for n := 0; n < b.N; n++ {
			err = json.Unmarshal(signedBlindedBeaconBlockJSONBytes, _blindedBeaconBlock)
			require.NoError(b, err)
		}
	})

	b.Run("Decode SSZ", func(b *testing.B) {
		_blindedBeaconBlock := new(SignedBlindedBeaconBlock)
		for n := 0; n < b.N; n++ {
			err = _blindedBeaconBlock.UnmarshalSSZ(signedBlindedBeaconBlockSSZBytes)
			require.NoError(b, err)
		}
	})
}
