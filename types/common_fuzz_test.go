package types

import (
	"encoding/json"
	"testing"

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

func FuzzFromSliceSignature(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var value Signature
		value.FromSlice(data)
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

func FuzzFromSlicePublicKey(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var value PublicKey
		value.FromSlice(data)
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

func FuzzFromSliceAddress(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var value Address
		value.FromSlice(data)
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

func FuzzFromSliceHash(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var value Hash
		value.FromSlice(data)
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

func FuzzFromSliceRoot(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var value Root
		value.FromSlice(data)
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

func FuzzFromSliceCommitteeBits(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var value CommitteeBits
		value.FromSlice(data)
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

func FuzzFromSliceBloom(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var value Bloom
		value.FromSlice(data)
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

func FuzzFromSliceU256Str(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var value U256Str
		value.FromSlice(data)
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

func FuzzFromSliceExtraData(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var value ExtraData
		value.FromSlice(data)
	})
}
