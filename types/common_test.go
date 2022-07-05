package types

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/stretchr/testify/require"
)

func TestJSONSerialization(t *testing.T) {
	a := Signature{0x01}
	b, err := json.Marshal(a)
	require.NoError(t, err)

	expectedHex := `0x010000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000`
	expectedJSON := fmt.Sprintf(`"%s"`, expectedHex)
	require.JSONEq(t, expectedJSON, string(b))

	ax := new(hexutil.Bytes)
	err = ax.UnmarshalJSON([]byte(expectedJSON))
	require.NoError(t, err)
	require.Equal(t, expectedHex, ax.String())

	a2 := Signature{}
	err = a2.UnmarshalJSON([]byte(expectedJSON))
	require.NoError(t, err)
	require.Equal(t, a, a2)
}

func TestU256Str(t *testing.T) {
	a := U256Str(common.HexToHash("0100000000000000000000000000000000000000000000000000000000000000"))
	require.Equal(t, "1", a.String())

	b, err := json.Marshal(a)
	require.NoError(t, err)

	expectedStr := `1`
	expectedJSON := fmt.Sprintf(`"%s"`, expectedStr)
	require.JSONEq(t, expectedJSON, string(b))

	// UnmarshalText
	a2 := U256Str{}
	err = a2.UnmarshalText([]byte(expectedStr))
	require.NoError(t, err)
	require.Equal(t, a, a2)

	// UnmarshalJSON
	a3 := U256Str{}
	err = a3.UnmarshalJSON([]byte(expectedJSON))
	require.NoError(t, err)
	require.Equal(t, a, a3)

	// IntToU256
	u := IntToU256(123)
	require.Equal(t, "123", u.String())
}

func TestU256StrCmp(t *testing.T) {
	one := IntToU256(1)
	two := IntToU256(2)
	bigger := IntToU256(256)
	require.Equal(t, "1", one.String())
	require.Equal(t, "2", two.String())
	require.Equal(t, "256", bigger.String())

	require.Equal(t, 0, one.Cmp(&one))
	require.Equal(t, 0, bigger.Cmp(&bigger))

	require.Equal(t, -1, one.Cmp(&two))
	require.Equal(t, -1, one.Cmp(&bigger))

	require.Equal(t, 1, bigger.Cmp(&two))
	require.Equal(t, 1, bigger.Cmp(&one))
}
