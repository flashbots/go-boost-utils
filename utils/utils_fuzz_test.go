package utils

import (
	"testing"
)

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
