package types

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestForkChoiceResponse(t *testing.T) {
	input := `{
		"payloadStatus": {
		  "status": "VALID",
		  "latestValidHash": "0x3b8fb240d288781d4aac94d3fd16809ee413bc99294a085798a589dae51ddd4a",
		  "validationError": null
		},
		"payloadId": "0xa247243752eb10b4"
	}`
	var result ForkChoiceResponse
	require.NoError(t, DecodeJSON(strings.NewReader(input), &result))
}
