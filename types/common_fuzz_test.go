package types

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/trailofbits/go-fuzz-utils"
)

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

func Fill[T interface{}](data []byte, value *T) bool {
	tp, err := GetTypeProvider(data)
	if err != nil {
		return false
	}
	err = tp.Fill(value)
	return err == nil
}

type MarshalableSSZ interface {
	MarshalSSZ() ([]byte, error)
	UnmarshalSSZ([]byte) error
}

func RoundTripSSZ[V MarshalableSSZ](t *testing.T, data []byte, decSSZ V) {
	value := *new(V)
	if !Fill(data, &value) {
		return
	}
	encSSZ, err := value.MarshalSSZ()
	require.NoError(t, err)
	err = decSSZ.UnmarshalSSZ(encSSZ)
	require.NoError(t, err)
	require.Equal(t, value, decSSZ)
}

func RoundTripJSON[V any](t *testing.T, data []byte, decJSON V) {
	value := *new(V)
	if !Fill(data, &value) {
		return
	}
	encJSON, err := json.Marshal(value)
	require.NoError(t, err)
	err = json.Unmarshal(encJSON, &decJSON)
	require.NoError(t, err)
	require.Equal(t, value, decJSON)
}

func FuzzReverse(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		require.Equal(t, data, reverse(reverse(data)))
	})
}

func FuzzRoundTripUint64StringSlice(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		RoundTripJSON(t, data, &Uint64StringSlice{})
	})
}

func FuzzRoundTripSignature(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		RoundTripJSON(t, data, &Signature{})
	})
}

func FuzzFromSliceSignature(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		new(Signature).FromSlice(data)
	})
}

func FuzzRoundTripPublicKey(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		RoundTripJSON(t, data, &PublicKey{})
	})
}

func FuzzFromSlicePublicKey(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		new(PublicKey).FromSlice(data)
	})
}

func FuzzRoundTripAddress(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		RoundTripJSON(t, data, &Address{})
	})
}

func FuzzFromSliceAddress(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		new(Address).FromSlice(data)
	})
}

func FuzzRoundTripHash(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		RoundTripJSON(t, data, &Hash{})
	})
}

func FuzzFromSliceHash(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		new(Hash).FromSlice(data)
	})
}

func FuzzRoundTripRoot(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		RoundTripJSON(t, data, &Root{})
	})
}

func FuzzFromSliceRoot(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		new(Root).FromSlice(data)
	})
}

func FuzzRoundTripCommitteeBits(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		RoundTripJSON(t, data, &CommitteeBits{})
	})
}

func FuzzFromSliceCommitteeBits(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		new(CommitteeBits).FromSlice(data)
	})
}

func FuzzRoundTripBloom(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		RoundTripJSON(t, data, &Bloom{})
	})
}

func FuzzFromSliceBloom(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		new(Bloom).FromSlice(data)
	})
}

func FuzzRoundTripU256Str(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		RoundTripJSON(t, data, &U256Str{})
	})
}

func FuzzFromSliceU256Str(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		new(U256Str).FromSlice(data)
	})
}

func FuzzRoundTripExtraData(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		RoundTripJSON(t, data, &ExtraData{})
	})
}

func FuzzFromSliceExtraData(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		new(ExtraData).FromSlice(data)
	})
}
