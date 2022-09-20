package types

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func FuzzNewPubkeyHex(f *testing.F) {
	f.Fuzz(func(t *testing.T, value string) {
		_ = NewPubkeyHex(value)
	})
}

func FuzzIntToU256(f *testing.F) {
	f.Fuzz(func(t *testing.T, value uint64) {
		u256 := IntToU256(value)
		bigInt := *(&u256).BigInt()
		require.True(t, bigInt.IsUint64())
		require.Equal(t, bigInt.Uint64(), value)
	})
}

func FuzzU256StrCmp(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		tp, err := GetTypeProvider(data)
		if err != nil {
			return
		}
		var a U256Str
		err = tp.Fill(&a)
		if err != nil {
			return
		}
		var b U256Str
		err = tp.Fill(&b)
		if err != nil {
			return
		}
		(&a).Cmp(&b)
	})
}

func FuzzHexToAddress(f *testing.F) {
	f.Fuzz(func(t *testing.T, str string) {
		_, _ = HexToAddress(str)
	})
}

func FuzzHexToPubkey(f *testing.F) {
	f.Fuzz(func(t *testing.T, str string) {
		_, _ = HexToPubkey(str)
	})
}

func FuzzHexToSignature(f *testing.F) {
	f.Fuzz(func(t *testing.T, str string) {
		_, _ = HexToSignature(str)
	})
}
