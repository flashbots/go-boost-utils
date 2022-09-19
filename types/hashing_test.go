package types

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCalculateHash(t *testing.T) {
	payloadJSON := `{
		"parent_hash": "0x89236ba32cb76b3f17cbba7620d956d561a08a42a22145bb5705099ed94eaddf",
		"fee_recipient": "0x0000000000000000000000000000000000000000",
		"state_root": "0x44f451f33692bc78735f7836ad9c25761ba15609155e7bfcb38ded400d95d500",
		"receipts_root": "0x056b23fbba480696b65fe5a59b8f2148a1299103c4f57df839233af2cf4ca2d2",
		"logs_bloom": "0x00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000",
		"prev_randao": "0x0000000000000000000000000000000000000000000000000000000000000000",
		"block_number": "11",
		"gas_limit": "4707788",
		"gas_used": "21000",
		"timestamp": "9155",
		"extra_data": "0x",
		"base_fee_per_gas": "233138867",
		"block_hash": "0x6662fb418aa7b5c5c80e2e8bc87be48db82e799c4704368d34ddeb3b12549655",
		"transactions": [
		  "0xf8670a843b9aca008252089400000000000000000000000000000000000000008203e880820a95a0ee3d06deddd2465aaa24bac5d329d3c40571c156c18d35c09a7c1daef2e95755a063e676889bbbdd27ab4e798b570f14ed8db32e4be22db15ab9f869c353b21f19"
		]
	}`
	payload := new(ExecutionPayload)
	require.NoError(t, DecodeJSON(strings.NewReader(payloadJSON), payload))

	hash := CalculateHash(payload)
	require.Equal(t, "0x6662fb418aa7b5c5c80e2e8bc87be48db82e799c4704368d34ddeb3b12549655", hash.String())
}
