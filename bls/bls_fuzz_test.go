package bls

import (
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

func FuzzPublicKeyFromBytes(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		_, _ = PublicKeyFromBytes(data)
	})
}

func FuzzPublicKeyFromBytesValidLength(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		tp, err := GetTypeProvider(data)
		if err != nil {
			return
		}
		bytes, err := tp.GetNBytes(PublicKeyLength)
		if err != nil {
			return
		}
		_, _ = PublicKeyFromBytes(bytes)
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
		pk, err := PublicKeyFromSecretKey(&sk)
		if err != nil {
			require.NotNil(t, pk)
		}
	})
}

func FuzzSecretKeyFromBytes(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		_, _ = SecretKeyFromBytes(data)
	})
}

func FuzzSecretKeyFromBytesValidLength(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		tp, err := GetTypeProvider(data)
		if err != nil {
			return
		}
		bytes, err := tp.GetNBytes(SecretKeyLength)
		if err != nil {
			return
		}
		_, _ = SecretKeyFromBytes(bytes)
	})
}

func FuzzSignatureFromBytes(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		_, _ = SignatureFromBytes(data)
	})
}

func FuzzSignatureFromBytesValidLength(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		tp, err := GetTypeProvider(data)
		if err != nil {
			return
		}
		bytes, err := tp.GetNBytes(SignatureLength)
		if err != nil {
			return
		}
		_, _ = SignatureFromBytes(bytes)
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
		_, _ = VerifySignature(&sig, &pk, msg)
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
		_, _ = VerifySignatureBytes(msg, sigBytes, pkBytes)
	})
}

func FuzzVerifySignatureBytesValidLength(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		tp, err := GetTypeProvider(data)
		if err != nil {
			return
		}
		sigBytes, err := tp.GetNBytes(SignatureLength)
		if err != nil {
			return
		}
		pkBytes, err := tp.GetNBytes(PublicKeyLength)
		if err != nil {
			return
		}
		msg, err := tp.GetBytes()
		if err != nil {
			return
		}
		_, _ = VerifySignatureBytes(msg, sigBytes, pkBytes)
	})
}
