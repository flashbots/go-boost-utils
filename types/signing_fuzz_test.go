package types

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func FuzzRoundTripSigningData(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var value, decSSZ, decJSON SigningData
		tp, err := GetTypeProvider(data)
		if err != nil {
			return
		}
		err = tp.Fill(&value)
		if err != nil {
			return
		}

		encSSZ, err := value.MarshalSSZ()
		require.NoError(t, err)
		err = decSSZ.UnmarshalSSZ(encSSZ)
		require.NoError(t, err)
		require.Equal(t, value, decSSZ)

		encJSON, err := json.Marshal(value)
		require.NoError(t, err)
		err = json.Unmarshal(encJSON, &decJSON)
		require.NoError(t, err)
		require.Equal(t, value, decJSON)
	})
}

func FuzzRoundTripForkData(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var value, decSSZ, decJSON ForkData
		tp, err := GetTypeProvider(data)
		if err != nil {
			return
		}
		err = tp.Fill(&value)
		if err != nil {
			return
		}

		encSSZ, err := value.MarshalSSZ()
		require.NoError(t, err)
		err = decSSZ.UnmarshalSSZ(encSSZ)
		require.NoError(t, err)
		require.Equal(t, value, decSSZ)

		encJSON, err := json.Marshal(value)
		require.NoError(t, err)
		err = json.Unmarshal(encJSON, &decJSON)
		require.NoError(t, err)
		require.Equal(t, value, decJSON)
	})
}
