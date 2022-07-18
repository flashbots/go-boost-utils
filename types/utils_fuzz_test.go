package types

import (
	"testing"
)

func FuzzHexToAddress(f *testing.F) {
	f.Fuzz(func(t *testing.T, str string) {
		HexToAddress(str)
	})
}

func FuzzHexToPubkey(f *testing.F) {
	f.Fuzz(func(t *testing.T, str string) {
		HexToPubkey(str)
	})
}

func FuzzHexToSignature(f *testing.F) {
	f.Fuzz(func(t *testing.T, str string) {
		HexToSignature(str)
	})
}
