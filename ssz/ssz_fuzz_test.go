package ssz

import (
	"encoding/json"
	"testing"

	"github.com/attestantio/go-eth2-client/spec/phase0"
	"github.com/flashbots/go-boost-utils/bls"
	"github.com/flashbots/go-boost-utils/fuzzing"
)

func FuzzRoundTripSigningData(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		fuzzing.RoundTripSSZ(t, data, &phase0.SigningData{})
		fuzzing.RoundTripJSON(t, data, &phase0.SigningData{})
	})
}

func FuzzUnmarshalSigningData(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var value phase0.SigningData
		_ = json.Unmarshal(data, &value)
		_ = value.UnmarshalSSZ(data)
	})
}

func FuzzRoundTripForkData(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		fuzzing.RoundTripSSZ(t, data, &phase0.ForkData{})
		fuzzing.RoundTripJSON(t, data, &phase0.ForkData{})
	})
}

func FuzzUnmarshalForkData(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var value phase0.ForkData
		_ = json.Unmarshal(data, &value)
		_ = value.UnmarshalSSZ(data)
	})
}

func FuzzComputeDomain(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		tp, err := fuzzing.GetTypeProvider(data)
		if err != nil {
			return
		}
		var domainType phase0.DomainType
		err = tp.Fill(&domainType)
		if err != nil {
			return
		}
		var forkVersion phase0.Version
		err = tp.Fill(&forkVersion)
		if err != nil {
			return
		}
		var genesisValidatorsRoot phase0.Root
		err = tp.Fill(&genesisValidatorsRoot)
		if err != nil {
			return
		}
		ComputeDomain(domainType, forkVersion, genesisValidatorsRoot)
	})
}

func FuzzComputeSigningRoot(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		tp, err := fuzzing.GetTypeProvider(data)
		if err != nil {
			return
		}
		var forkData phase0.ForkData
		err = tp.Fill(&forkData)
		if err != nil {
			return
		}
		var domain phase0.Domain
		err = tp.Fill(&domain)
		if err != nil {
			return
		}
		_, _ = ComputeSigningRoot(&forkData, domain)
	})
}

func FuzzSignMessage(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		tp, err := fuzzing.GetTypeProvider(data)
		if err != nil {
			return
		}
		var forkData phase0.ForkData
		err = tp.Fill(&forkData)
		if err != nil {
			return
		}
		var domain phase0.Domain
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
		tp, err := fuzzing.GetTypeProvider(data)
		if err != nil {
			return
		}
		var forkData phase0.ForkData
		err = tp.Fill(&forkData)
		if err != nil {
			return
		}
		var domain phase0.Domain
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
