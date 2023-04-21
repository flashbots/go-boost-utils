package fuzzing

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	go_fuzz_utils "github.com/trailofbits/go-fuzz-utils"
)

func GetTypeProvider(data []byte) (*go_fuzz_utils.TypeProvider, error) {
	tp, err := go_fuzz_utils.NewTypeProvider(data)
	if err != nil {
		return nil, err
	}
	err = tp.SetParamsStringBounds(0, 1024)
	if err != nil {
		return nil, err
	}
	err = tp.SetParamsSliceBounds(0, 4096)
	if err != nil {
		return nil, err
	}
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
	t.Helper()
	value := *new(V)
	if !Fill(data, &value) {
		return
	}
	t.Log(value)

	encSSZ, err := value.MarshalSSZ()
	if err != nil {
		if strings.Contains(err.Error(), "list length is higher than max value") {
			return
		}
		if strings.Contains(err.Error(), "does not have the correct length") {
			return
		}
	}
	require.NoError(t, err)

	err = decSSZ.UnmarshalSSZ(encSSZ)
	if err != nil {
		if strings.Contains(err.Error(), "unexpected number of bytes") {
			return
		}
		if err.Error() == "bitlist empty, it does not have length bit" {
			return
		}
		if err.Error() == "trailing byte is zero" {
			return
		}
		if err.Error() == "too many bits" {
			return
		}
	}
	require.NoError(t, err)
	require.Equal(t, value, decSSZ)
}

func RoundTripJSON[V any](t *testing.T, data []byte, decJSON V) {
	t.Helper()
	value := *new(V)
	if !Fill(data, &value) {
		return
	}
	t.Log(value)

	encJSON, err := json.Marshal(value)
	require.NoError(t, err)

	err = json.Unmarshal(encJSON, &decJSON)
	if err != nil {
		if err.Error() == "incorrect byte length" {
			return
		}
	}
	require.NoError(t, err)
	require.Equal(t, value, decJSON)
}
