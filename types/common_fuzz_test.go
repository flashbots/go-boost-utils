package types

import (
	"testing"
	"encoding/json"

	"github.com/stretchr/testify/require"
	"github.com/trailofbits/go-fuzz-utils"
)

func FuzzReverse(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		require.Equal(t, data, reverse(reverse(data)))
	})
}

func GetTypeProvider(data []byte) (*go_fuzz_utils.TypeProvider, error) {
	tp, err := go_fuzz_utils.NewTypeProvider(data)
	if err != nil {
		return nil, err
	}

	// Prevent nil fields.
	err = tp.SetParamsBiases(0, 0, 0, 0)
	if err != nil {
		return nil, err
	}

	return tp, nil
}

func FuzzRoundTripUint64StringSlice(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var value, decJSON Uint64StringSlice
		tp, err := GetTypeProvider(data)
		if err != nil {
			return
		}
		err = tp.Fill(&value)
		if err != nil {
			return
		}

		encJSON, err := json.Marshal(value)
		require.NoError(t, err)
		err = json.Unmarshal(encJSON, &decJSON)
		require.NoError(t, err)
		require.Equal(t, value, decJSON)
	})
}

func FuzzRoundTripSignature(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var value, decJSON Signature
		tp, err := GetTypeProvider(data)
		if err != nil {
			return
		}
		err = tp.Fill(&value)
		if err != nil {
			return
		}

		encJSON, err := json.Marshal(value)
		require.NoError(t, err)
		err = json.Unmarshal(encJSON, &decJSON)
		require.NoError(t, err)
		require.Equal(t, value, decJSON)
	})
}

func FuzzRoundTripPublicKey(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var value, decJSON PublicKey
		tp, err := GetTypeProvider(data)
		if err != nil {
			return
		}
		err = tp.Fill(&value)
		if err != nil {
			return
		}

		encJSON, err := json.Marshal(value)
		require.NoError(t, err)
		err = json.Unmarshal(encJSON, &decJSON)
		require.NoError(t, err)
		require.Equal(t, value, decJSON)
	})
}

func FuzzRoundTripAddress(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var value, decJSON Address
		tp, err := GetTypeProvider(data)
		if err != nil {
			return
		}
		err = tp.Fill(&value)
		if err != nil {
			return
		}

		encJSON, err := json.Marshal(value)
		require.NoError(t, err)
		err = json.Unmarshal(encJSON, &decJSON)
		require.NoError(t, err)
		require.Equal(t, value, decJSON)
	})
}

func FuzzRoundTripHash(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var value, decJSON Hash
		tp, err := GetTypeProvider(data)
		if err != nil {
			return
		}
		err = tp.Fill(&value)
		if err != nil {
			return
		}

		encJSON, err := json.Marshal(value)
		require.NoError(t, err)
		err = json.Unmarshal(encJSON, &decJSON)
		require.NoError(t, err)
		require.Equal(t, value, decJSON)
	})
}

func FuzzRoundTripRoot(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var value, decJSON Root
		tp, err := GetTypeProvider(data)
		if err != nil {
			return
		}
		err = tp.Fill(&value)
		if err != nil {
			return
		}

		encJSON, err := json.Marshal(value)
		require.NoError(t, err)
		err = json.Unmarshal(encJSON, &decJSON)
		require.NoError(t, err)
		require.Equal(t, value, decJSON)
	})
}

func FuzzRoundTripCommitteeBits(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var value, decJSON CommitteeBits
		tp, err := GetTypeProvider(data)
		if err != nil {
			return
		}
		err = tp.Fill(&value)
		if err != nil {
			return
		}

		encJSON, err := json.Marshal(value)
		require.NoError(t, err)
		err = json.Unmarshal(encJSON, &decJSON)
		require.NoError(t, err)
		require.Equal(t, value, decJSON)
	})
}

func FuzzRoundTripBloom(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var value, decJSON Bloom
		tp, err := GetTypeProvider(data)
		if err != nil {
			return
		}
		err = tp.Fill(&value)
		if err != nil {
			return
		}

		encJSON, err := json.Marshal(value)
		require.NoError(t, err)
		err = json.Unmarshal(encJSON, &decJSON)
		require.NoError(t, err)
		require.Equal(t, value, decJSON)
	})
}

func FuzzRoundTripU256Str(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var value, decJSON U256Str
		tp, err := GetTypeProvider(data)
		if err != nil {
			return
		}
		err = tp.Fill(&value)
		if err != nil {
			return
		}

		encJSON, err := json.Marshal(value)
		require.NoError(t, err)
		err = json.Unmarshal(encJSON, &decJSON)
		require.NoError(t, err)
		require.Equal(t, value, decJSON)
	})
}

func FuzzRoundTripExtraData(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var value, decJSON ExtraData
		tp, err := GetTypeProvider(data)
		if err != nil {
			return
		}
		err = tp.Fill(&value)
		if err != nil {
			return
		}

		encJSON, err := json.Marshal(value)
		require.NoError(t, err)
		err = json.Unmarshal(encJSON, &decJSON)
		require.NoError(t, err)
		require.Equal(t, value, decJSON)
	})
}
