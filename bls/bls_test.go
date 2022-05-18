package bls

import (
	"testing"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/stretchr/testify/require"
)

func TestSecretToPubkey(t *testing.T) {
	for _, tc := range []struct {
		Input  []byte
		Output []byte
	}{
		{hexutil.MustDecode("0x263dbd792f5b1be47ed85f8938c0f29586af0d3ac7b977f21c278fe1462040e3"), hexutil.MustDecode("0xa491d1b0ecd9bb917989f0e74f0dea0422eac4a873e5e2644f368dffb9a6e20fd6e10c1b77654d067c0618f6e5a7f79a")},
		{hexutil.MustDecode("0x47b8192d77bf871b62e87859d653922725724a5c031afeabc60bcef5ff665138"), hexutil.MustDecode("0xb301803f8b5ac4a1133581fc676dfedc60d891dd5fa99028805e5ea5b08d3491af75d0707adab3b70c6a6a580217bf81")},
		{hexutil.MustDecode("0x328388aff0d4a5b7dc9205abd374e7e98f3cd9f3418edb4eafda5fb16473d216"), hexutil.MustDecode("0xb53d21a4cfd562c469cc81514d4ce5a6b577d8403d32a394dc265dd190b47fa9f829fdd7963afdf972e5e77854051f6f")},
	} {
		sk, err := SecretKeyFromBytes(tc.Input)
		require.NoError(t, err)
		pk := PublicKeyFromSecretKey(sk)
		require.Equal(t, pk.Compress(), tc.Output)
	}
}
