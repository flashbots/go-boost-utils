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

// Stolen from Teku's BLSTest class:
// https://github.com/ConsenSys/teku/blob/ce6e13e32f2c5029aa05895f400c00a9c72a3f38/infrastructure/bls/src/test/java/tech/pegasys/teku/bls/BLSTest.java#L318-L332
func TestSignatureVerifyRealValues(t *testing.T) {
    signingRootBytes := hexutil.MustDecode("0x95b8e2ba063ab62f68ebe7db0a9669ab9e7906aa4e060e1cc0b67b294ce8c5e4")
    sigBytes := hexutil.MustDecode("0xab51f352e90509ca5085ec43af9ad3ea4ae42bf30c91af7dcdc113ef79cfc8601b756f18d8cf634436d8b6b0095fc5680066f382eb3728a7090c55c9afb66e8f94b44d2682db8ef5de4b89928d1744824df174e0c800b9e934b0ad14e6388163")
    pkBytes := hexutil.MustDecode("0xb5e8f551c28abd6ef8253581ffad0834bfd8fafa9948d09b337c9c5f21d6e7fd6065a1ee35ac5146ac17344f97490301")

    result, err := VerifySignatureBytes(signingRootBytes, sigBytes, pkBytes)
    require.NoError(t, err)
    require.Equal(t, result, true)
}
