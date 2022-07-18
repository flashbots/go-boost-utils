package types

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func FuzzRoundTripEth1Data(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var value, decSSZ, decJSON Eth1Data
		tp, err := GetTypeProvider(data)
		if err != nil {
			return
		}
		err = tp.Fill(&value)
		if err != nil {
			return
		}

		encSSZ, err := value.MarshalSSZ()
		require.NoError(t, err)
		err = decSSZ.UnmarshalSSZ(encSSZ)
		require.NoError(t, err)
		require.Equal(t, value, decSSZ)

		encJSON, err := json.Marshal(value)
		require.NoError(t, err)
		err = json.Unmarshal(encJSON, &decJSON)
		require.NoError(t, err)
		require.Equal(t, value, decJSON)
	})
}

func FuzzRoundTripBeaconBlockHeader(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var value, decSSZ, decJSON BeaconBlockHeader
		tp, err := GetTypeProvider(data)
		if err != nil {
			return
		}
		err = tp.Fill(&value)
		if err != nil {
			return
		}

		encSSZ, err := value.MarshalSSZ()
		require.NoError(t, err)
		err = decSSZ.UnmarshalSSZ(encSSZ)
		require.NoError(t, err)
		require.Equal(t, value, decSSZ)

		encJSON, err := json.Marshal(value)
		require.NoError(t, err)
		err = json.Unmarshal(encJSON, &decJSON)
		require.NoError(t, err)
		require.Equal(t, value, decJSON)
	})
}

func FuzzRoundTripSignedBeaconBlockHeader(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var value, decSSZ, decJSON SignedBeaconBlockHeader
		tp, err := GetTypeProvider(data)
		if err != nil {
			return
		}
		err = tp.Fill(&value)
		if err != nil {
			return
		}

		encSSZ, err := value.MarshalSSZ()
		require.NoError(t, err)
		err = decSSZ.UnmarshalSSZ(encSSZ)
		require.NoError(t, err)
		require.Equal(t, value, decSSZ)

		encJSON, err := json.Marshal(value)
		require.NoError(t, err)
		err = json.Unmarshal(encJSON, &decJSON)
		require.NoError(t, err)
		require.Equal(t, value, decJSON)
	})
}

func FuzzRoundTripProposerSlashing(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var value, decSSZ, decJSON ProposerSlashing
		tp, err := GetTypeProvider(data)
		if err != nil {
			return
		}
		err = tp.Fill(&value)
		if err != nil {
			return
		}

		encSSZ, err := value.MarshalSSZ()
		require.NoError(t, err)
		err = decSSZ.UnmarshalSSZ(encSSZ)
		require.NoError(t, err)
		require.Equal(t, value, decSSZ)

		encJSON, err := json.Marshal(value)
		require.NoError(t, err)
		err = json.Unmarshal(encJSON, &decJSON)
		require.NoError(t, err)
		require.Equal(t, value, decJSON)
	})
}

func FuzzRoundTripCheckpoint(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var value, decSSZ, decJSON Checkpoint
		tp, err := GetTypeProvider(data)
		if err != nil {
			return
		}
		err = tp.Fill(&value)
		if err != nil {
			return
		}

		encSSZ, err := value.MarshalSSZ()
		require.NoError(t, err)
		err = decSSZ.UnmarshalSSZ(encSSZ)
		require.NoError(t, err)
		require.Equal(t, value, decSSZ)

		encJSON, err := json.Marshal(value)
		require.NoError(t, err)
		err = json.Unmarshal(encJSON, &decJSON)
		require.NoError(t, err)
		require.Equal(t, value, decJSON)
	})
}

func FuzzRoundTripAttestationData(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var value, decSSZ, decJSON AttestationData
		tp, err := GetTypeProvider(data)
		if err != nil {
			return
		}
		err = tp.Fill(&value)
		if err != nil {
			return
		}

		encSSZ, err := value.MarshalSSZ()
		require.NoError(t, err)
		err = decSSZ.UnmarshalSSZ(encSSZ)
		require.NoError(t, err)
		require.Equal(t, value, decSSZ)

		encJSON, err := json.Marshal(value)
		require.NoError(t, err)
		err = json.Unmarshal(encJSON, &decJSON)
		require.NoError(t, err)
		require.Equal(t, value, decJSON)
	})
}

func FuzzRoundTripIndexedAttestation(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var value, decSSZ, decJSON IndexedAttestation
		tp, err := GetTypeProvider(data)
		if err != nil {
			return
		}
		err = tp.Fill(&value)
		if err != nil {
			return
		}

		encSSZ, err := value.MarshalSSZ()
		require.NoError(t, err)
		err = decSSZ.UnmarshalSSZ(encSSZ)
		require.NoError(t, err)
		require.Equal(t, value, decSSZ)

		encJSON, err := json.Marshal(value)
		require.NoError(t, err)
		err = json.Unmarshal(encJSON, &decJSON)
		require.NoError(t, err)
		require.Equal(t, value, decJSON)
	})
}

func FuzzRoundTripAttesterSlashing(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var value, decSSZ, decJSON AttesterSlashing
		tp, err := GetTypeProvider(data)
		if err != nil {
			return
		}
		err = tp.Fill(&value)
		if err != nil {
			return
		}

		encSSZ, err := value.MarshalSSZ()
		require.NoError(t, err)
		err = decSSZ.UnmarshalSSZ(encSSZ)
		require.NoError(t, err)
		require.Equal(t, value, decSSZ)

		encJSON, err := json.Marshal(value)
		require.NoError(t, err)
		err = json.Unmarshal(encJSON, &decJSON)
		require.NoError(t, err)
		require.Equal(t, value, decJSON)
	})
}

func FuzzRoundTripAttestation(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var value, decSSZ, decJSON Attestation
		tp, err := GetTypeProvider(data)
		if err != nil {
			return
		}
		err = tp.Fill(&value)
		if err != nil {
			return
		}

		encSSZ, err := value.MarshalSSZ()
		require.NoError(t, err)
		err = decSSZ.UnmarshalSSZ(encSSZ)
		if err == nil {
			require.Equal(t, value, decSSZ)
		}

		encJSON, err := json.Marshal(value)
		require.NoError(t, err)
		err = json.Unmarshal(encJSON, &decJSON)
		require.NoError(t, err)
		require.Equal(t, value, decJSON)
	})
}

func FuzzRoundTripDeposit(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var value, decSSZ, decJSON Deposit
		tp, err := GetTypeProvider(data)
		if err != nil {
			return
		}
		err = tp.Fill(&value)
		if err != nil {
			return
		}

		encSSZ, err := value.MarshalSSZ()
		require.NoError(t, err)
		err = decSSZ.UnmarshalSSZ(encSSZ)
		require.NoError(t, err)
		require.Equal(t, value, decSSZ)

		encJSON, err := json.Marshal(value)
		require.NoError(t, err)
		err = json.Unmarshal(encJSON, &decJSON)
		require.NoError(t, err)
		require.Equal(t, value, decJSON)
	})
}

func FuzzRoundTripSyncAggregate(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var value, decSSZ, decJSON SyncAggregate
		tp, err := GetTypeProvider(data)
		if err != nil {
			return
		}
		err = tp.Fill(&value)
		if err != nil {
			return
		}

		encSSZ, err := value.MarshalSSZ()
		require.NoError(t, err)
		err = decSSZ.UnmarshalSSZ(encSSZ)
		require.NoError(t, err)
		require.Equal(t, value, decSSZ)

		encJSON, err := json.Marshal(value)
		require.NoError(t, err)
		err = json.Unmarshal(encJSON, &decJSON)
		require.NoError(t, err)
		require.Equal(t, value, decJSON)
	})
}

func FuzzRoundTripVoluntaryExit(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var value, decSSZ, decJSON VoluntaryExit
		tp, err := GetTypeProvider(data)
		if err != nil {
			return
		}
		err = tp.Fill(&value)
		if err != nil {
			return
		}

		encSSZ, err := value.MarshalSSZ()
		require.NoError(t, err)
		err = decSSZ.UnmarshalSSZ(encSSZ)
		require.NoError(t, err)
		require.Equal(t, value, decSSZ)

		encJSON, err := json.Marshal(value)
		require.NoError(t, err)
		err = json.Unmarshal(encJSON, &decJSON)
		require.NoError(t, err)
		require.Equal(t, value, decJSON)
	})
}

func FuzzRoundTripExecutionPayloadHeader(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var value, decSSZ, decJSON ExecutionPayloadHeader
		tp, err := GetTypeProvider(data)
		if err != nil {
			return
		}
		err = tp.Fill(&value)
		if err != nil {
			return
		}

		encSSZ, err := value.MarshalSSZ()
		require.NoError(t, err)
		err = decSSZ.UnmarshalSSZ(encSSZ)
		require.NoError(t, err)
		require.Equal(t, value, decSSZ)

		encJSON, err := json.Marshal(value)
		require.NoError(t, err)
		err = json.Unmarshal(encJSON, &decJSON)
		require.NoError(t, err)
		require.Equal(t, value, decJSON)
	})
}

func FuzzRoundTripExecutionPayload(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var value, decJSON ExecutionPayload
		tp, err := GetTypeProvider(data)
		if err != nil {
			return
		}
		err = tp.Fill(&value)
		if err != nil {
			return
		}

		encJSON, err := json.Marshal(value)
		require.NoError(t, err)
		err = json.Unmarshal(encJSON, &decJSON)
		require.NoError(t, err)
		require.Equal(t, value, decJSON)
	})
}

func FuzzRoundTripBlindedBeaconBlockBody(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var value, decSSZ, decJSON BlindedBeaconBlockBody
		tp, err := GetTypeProvider(data)
		if err != nil {
			return
		}
		err = tp.Fill(&value)
		if err != nil {
			return
		}

		encSSZ, err := value.MarshalSSZ()
		require.NoError(t, err)
		err = decSSZ.UnmarshalSSZ(encSSZ)
		require.NoError(t, err)
		require.Equal(t, value, decSSZ)

		encJSON, err := json.Marshal(value)
		require.NoError(t, err)
		err = json.Unmarshal(encJSON, &decJSON)
		require.NoError(t, err)
		require.Equal(t, value, decJSON)
	})
}

func FuzzRoundTripBlindedBeaconBlock(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var value, decSSZ, decJSON BlindedBeaconBlock
		tp, err := GetTypeProvider(data)
		if err != nil {
			return
		}
		err = tp.Fill(&value)
		if err != nil {
			return
		}

		encSSZ, err := value.MarshalSSZ()
		require.NoError(t, err)
		err = decSSZ.UnmarshalSSZ(encSSZ)
		require.NoError(t, err)
		require.Equal(t, value, decSSZ)

		encJSON, err := json.Marshal(value)
		require.NoError(t, err)
		err = json.Unmarshal(encJSON, &decJSON)
		require.NoError(t, err)
		require.Equal(t, value, decJSON)
	})
}

func FuzzRoundTripRegisterValidatorRequestMessage(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var value, decSSZ, decJSON RegisterValidatorRequestMessage
		tp, err := GetTypeProvider(data)
		if err != nil {
			return
		}
		err = tp.Fill(&value)
		if err != nil {
			return
		}

		encSSZ, err := value.MarshalSSZ()
		require.NoError(t, err)
		err = decSSZ.UnmarshalSSZ(encSSZ)
		require.NoError(t, err)
		require.Equal(t, value, decSSZ)

		encJSON, err := json.Marshal(value)
		require.NoError(t, err)
		err = json.Unmarshal(encJSON, &decJSON)
		require.NoError(t, err)
		require.Equal(t, value, decJSON)
	})
}

func FuzzRoundTripSignedValidatorRegistration(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var value, decJSON SignedValidatorRegistration
		tp, err := GetTypeProvider(data)
		if err != nil {
			return
		}
		err = tp.Fill(&value)
		if err != nil {
			return
		}

		encJSON, err := json.Marshal(value)
		require.NoError(t, err)
		err = json.Unmarshal(encJSON, &decJSON)
		require.NoError(t, err)
		require.Equal(t, value, decJSON)
	})
}

func FuzzRoundTripBuilderBid(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var value, decSSZ, decJSON BuilderBid
		tp, err := GetTypeProvider(data)
		if err != nil {
			return
		}
		err = tp.Fill(&value)
		if err != nil {
			return
		}

		encSSZ, err := value.MarshalSSZ()
		require.NoError(t, err)
		err = decSSZ.UnmarshalSSZ(encSSZ)
		require.NoError(t, err)
		require.Equal(t, value, decSSZ)

		encJSON, err := json.Marshal(value)
		require.NoError(t, err)
		err = json.Unmarshal(encJSON, &decJSON)
		require.NoError(t, err)
		require.Equal(t, value, decJSON)
	})
}

func FuzzRoundTripSignedBuilderBid(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var value, decSSZ, decJSON SignedBuilderBid
		tp, err := GetTypeProvider(data)
		if err != nil {
			return
		}
		err = tp.Fill(&value)
		if err != nil {
			return
		}

		encSSZ, err := value.MarshalSSZ()
		require.NoError(t, err)
		err = decSSZ.UnmarshalSSZ(encSSZ)
		require.NoError(t, err)
		require.Equal(t, value, decSSZ)

		encJSON, err := json.Marshal(value)
		require.NoError(t, err)
		err = json.Unmarshal(encJSON, &decJSON)
		require.NoError(t, err)
		require.Equal(t, value, decJSON)
	})
}

func FuzzRoundTripGetHeaderResponse(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var value, decJSON GetHeaderResponse
		tp, err := GetTypeProvider(data)
		if err != nil {
			return
		}
		err = tp.Fill(&value)
		if err != nil {
			return
		}

		encJSON, err := json.Marshal(value)
		require.NoError(t, err)
		err = json.Unmarshal(encJSON, &decJSON)
		require.NoError(t, err)
		require.Equal(t, value, decJSON)
	})
}

func FuzzRoundTripSignedBlindedBeaconBlock(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var value, decJSON SignedBlindedBeaconBlock
		tp, err := GetTypeProvider(data)
		if err != nil {
			return
		}
		err = tp.Fill(&value)
		if err != nil {
			return
		}

		encJSON, err := json.Marshal(value)
		require.NoError(t, err)
		err = json.Unmarshal(encJSON, &decJSON)
		require.NoError(t, err)
		require.Equal(t, value, decJSON)
	})
}

func FuzzRoundTripGetPayloadResponse(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var value, decJSON GetPayloadResponse
		tp, err := GetTypeProvider(data)
		if err != nil {
			return
		}
		err = tp.Fill(&value)
		if err != nil {
			return
		}

		encJSON, err := json.Marshal(value)
		require.NoError(t, err)
		err = json.Unmarshal(encJSON, &decJSON)
		require.NoError(t, err)
		require.Equal(t, value, decJSON)
	})
}

func FuzzRoundTripTransactions(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var value, decSSZ, decJSON Transactions
		tp, err := GetTypeProvider(data)
		if err != nil {
			return
		}
		err = tp.Fill(&value)
		if err != nil {
			return
		}

		encSSZ, err := value.MarshalSSZ()
		require.NoError(t, err)
		err = decSSZ.UnmarshalSSZ(encSSZ)
		require.NoError(t, err)
		require.Equal(t, value, decSSZ)

		encJSON, err := json.Marshal(value)
		require.NoError(t, err)
		err = json.Unmarshal(encJSON, &decJSON)
		require.NoError(t, err)
		require.Equal(t, value, decJSON)
	})
}
