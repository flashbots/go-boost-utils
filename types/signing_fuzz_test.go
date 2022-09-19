package types

import (
	"encoding/json"
	"testing"

	"github.com/flashbots/go-boost-utils/bls"
)

func FuzzRoundTripSigningData(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		RoundTripSSZ(t, data, &SigningData{})
		RoundTripJSON(t, data, &SigningData{})
	})
}

func FuzzUnmarshalSigningData(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var value SigningData
		_ = json.Unmarshal(data, &value)
		_ = value.UnmarshalSSZ(data)
	})
}

func FuzzRoundTripForkData(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		RoundTripSSZ(t, data, &ForkData{})
		RoundTripJSON(t, data, &ForkData{})
	})
}

func FuzzUnmarshalForkData(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var value ForkData
		_ = json.Unmarshal(data, &value)
		_ = value.UnmarshalSSZ(data)
	})
}

func FuzzComputeDomain(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		tp, err := GetTypeProvider(data)
		if err != nil {
			return
		}
		var domainType DomainType
		err = tp.Fill(&domainType)
		if err != nil {
			return
		}
		var forkVersion ForkVersion
		err = tp.Fill(&forkVersion)
		if err != nil {
			return
		}
		var genesisValidatorsRoot Root
		err = tp.Fill(&genesisValidatorsRoot)
		if err != nil {
			return
		}
		ComputeDomain(domainType, forkVersion, genesisValidatorsRoot)
	})
}

func FuzzComputeSigningRoot(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		tp, err := GetTypeProvider(data)
		if err != nil {
			return
		}
		var forkData ForkData
		err = tp.Fill(&forkData)
		if err != nil {
			return
		}
		var domain Domain
		err = tp.Fill(&domain)
		if err != nil {
			return
		}
		_, _ = ComputeSigningRoot(&forkData, domain)
	})
}

func FuzzSignMessage(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		tp, err := GetTypeProvider(data)
		if err != nil {
			return
		}
		var forkData ForkData
		err = tp.Fill(&forkData)
		if err != nil {
			return
		}
		var domain Domain
		err = tp.Fill(&domain)
		if err != nil {
			return
		}
		var pkBytes []byte
		err = tp.Fill(&pkBytes)
		if err != nil {
			return
		}
		var sk bls.SecretKey
		err = tp.Fill(&sk)
		if err != nil {
			return
		}
		_, _ = SignMessage(&forkData, domain, &sk)
	})
}

func FuzzVerifySignature(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		tp, err := GetTypeProvider(data)
		if err != nil {
			return
		}
		var forkData ForkData
		err = tp.Fill(&forkData)
		if err != nil {
			return
		}
		var domain Domain
		err = tp.Fill(&domain)
		if err != nil {
			return
		}
		var pkBytes []byte
		err = tp.Fill(&pkBytes)
		if err != nil {
			return
		}
		var sigBytes []byte
		err = tp.Fill(&sigBytes)
		if err != nil {
			return
		}
		_, _ = VerifySignature(&forkData, domain, pkBytes, sigBytes)
	})
}
