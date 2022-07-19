package bls

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/trailofbits/go-fuzz-utils"
)

func GetTypeProvider(data []byte) (*go_fuzz_utils.TypeProvider, error) {
	tp, err := go_fuzz_utils.NewTypeProvider(data)
	if err != nil {
		return nil, err
	}

	// Prevent nil fields.
	err = tp.SetParamsBiases(0, 0, 0, 0)
	if err != nil {
		return nil, err
	}

	return tp, nil
}

func FuzzPublicKeyFromBytes(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		PublicKeyFromBytes(data)
	})
}

func FuzzPublicKeyFromBytesValidLength(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		tp, err := GetTypeProvider(data)
		if err != nil {
			return
		}
		bytes, err := tp.GetNBytes(BLSPublicKeyLength)
		if err != nil {
			return
		}
		PublicKeyFromBytes(bytes)
	})
}

func FuzzPublicKeyFromSecretKey(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		tp, err := GetTypeProvider(data)
		if err != nil {
			return
		}

		var sk SecretKey
		err = tp.Fill(&sk)
		if err != nil {
			return
		}

		pk := PublicKeyFromSecretKey(&sk)
		require.NotNil(t, pk)
	})
}

func FuzzSecretKeyFromBytes(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		SecretKeyFromBytes(data)
	})
}

func FuzzSecretKeyFromBytesValidLength(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		tp, err := GetTypeProvider(data)
		if err != nil {
			return
		}
		bytes, err := tp.GetNBytes(BLSSecretKeyLength)
		if err != nil {
			return
		}
		SecretKeyFromBytes(bytes)
	})
}

func FuzzSignatureFromBytes(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		SignatureFromBytes(data)
	})
}

func FuzzSignatureFromBytesValidLength(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		tp, err := GetTypeProvider(data)
		if err != nil {
			return
		}
		bytes, err := tp.GetNBytes(BLSSignatureLength)
		if err != nil {
			return
		}
		SignatureFromBytes(bytes)
	})
}

func FuzzSign(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		tp, err := GetTypeProvider(data)
		if err != nil {
			return
		}

		var sk SecretKey
		err = tp.Fill(&sk)
		if err != nil {
			return
		}
		msg, err := tp.GetBytes()
		if err != nil {
			return
		}

		Sign(&sk, msg)
	})
}

func FuzzVerifySignature(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		tp, err := GetTypeProvider(data)
		if err != nil {
			return
		}

		var sig Signature
		err = tp.Fill(&sig)
		if err != nil {
			return
		}
		var pk PublicKey
		err = tp.Fill(&pk)
		if err != nil {
			return
		}
		msg, err := tp.GetBytes()
		if err != nil {
			return
		}

		VerifySignature(&sig, &pk, msg)
	})
}

func FuzzVerifySignatureBytes(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		tp, err := GetTypeProvider(data)
		if err != nil {
			return
		}

		sigBytes, err := tp.GetBytes()
		if err != nil {
			return
		}
		pkBytes, err := tp.GetBytes()
		if err != nil {
			return
		}
		msg, err := tp.GetBytes()
		if err != nil {
			return
		}

		VerifySignatureBytes(msg, sigBytes, pkBytes)
	})
}

func FuzzVerifySignatureBytesValidLength(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		tp, err := GetTypeProvider(data)
		if err != nil {
			return
		}

		sigBytes, err := tp.GetNBytes(BLSSignatureLength)
		if err != nil {
			return
		}
		pkBytes, err := tp.GetNBytes(BLSPublicKeyLength)
		if err != nil {
			return
		}
		msg, err := tp.GetBytes()
		if err != nil {
			return
		}

		VerifySignatureBytes(msg, sigBytes, pkBytes)
	})
}
