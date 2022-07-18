package types

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func FuzzReverse(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		require.Equal(t, data, reverse(reverse(data)))
	})
}
