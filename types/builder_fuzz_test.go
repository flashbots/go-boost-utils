package types

import (
	"testing"
)

func FuzzRoundTripEth1Data(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		RoundTripSSZ(t, data, &Eth1Data{})
		RoundTripJSON(t, data, &Eth1Data{})
	})
}

func FuzzRoundTripBeaconBlockHeader(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		RoundTripSSZ(t, data, &BeaconBlockHeader{})
		RoundTripJSON(t, data, &BeaconBlockHeader{})
	})
}

func FuzzRoundTripSignedBeaconBlockHeader(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		RoundTripSSZ(t, data, &SignedBeaconBlockHeader{})
		RoundTripJSON(t, data, &SignedBeaconBlockHeader{})
	})
}

func FuzzRoundTripProposerSlashing(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		RoundTripSSZ(t, data, &ProposerSlashing{})
		RoundTripJSON(t, data, &ProposerSlashing{})
	})
}

func FuzzRoundTripCheckpoint(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		RoundTripSSZ(t, data, &Checkpoint{})
		RoundTripJSON(t, data, &Checkpoint{})
	})
}

func FuzzRoundTripAttestationData(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		RoundTripSSZ(t, data, &AttestationData{})
		RoundTripJSON(t, data, &AttestationData{})
	})
}

func FuzzRoundTripIndexedAttestation(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		RoundTripSSZ(t, data, &IndexedAttestation{})
		RoundTripJSON(t, data, &IndexedAttestation{})
	})
}

func FuzzRoundTripAttesterSlashing(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		RoundTripSSZ(t, data, &AttesterSlashing{})
		RoundTripJSON(t, data, &AttesterSlashing{})
	})
}

func FuzzRoundTripAttestation(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		RoundTripSSZ(t, data, &Attestation{})
		RoundTripJSON(t, data, &Attestation{})
	})
}

func FuzzRoundTripDeposit(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		RoundTripSSZ(t, data, &Deposit{})
		RoundTripJSON(t, data, &Deposit{})
	})
}

func FuzzRoundTripSyncAggregate(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		RoundTripSSZ(t, data, &SyncAggregate{})
		RoundTripJSON(t, data, &SyncAggregate{})
	})
}

func FuzzRoundTripVoluntaryExit(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		RoundTripSSZ(t, data, &VoluntaryExit{})
		RoundTripJSON(t, data, &VoluntaryExit{})
	})
}

func FuzzRoundTripExecutionPayloadHeader(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		RoundTripSSZ(t, data, &ExecutionPayloadHeader{})
		RoundTripJSON(t, data, &ExecutionPayloadHeader{})
	})
}

func FuzzRoundTripExecutionPayload(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		RoundTripJSON(t, data, &ExecutionPayload{})
	})
}

func FuzzRoundTripBlindedBeaconBlockBody(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		RoundTripSSZ(t, data, &BlindedBeaconBlockBody{})
		RoundTripJSON(t, data, &BlindedBeaconBlockBody{})
	})
}

func FuzzRoundTripBlindedBeaconBlock(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		RoundTripSSZ(t, data, &BlindedBeaconBlock{})
		RoundTripJSON(t, data, &BlindedBeaconBlock{})
	})
}

func FuzzRoundTripRegisterValidatorRequestMessage(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		RoundTripSSZ(t, data, &RegisterValidatorRequestMessage{})
		RoundTripJSON(t, data, &RegisterValidatorRequestMessage{})
	})
}

func FuzzRoundTripSignedValidatorRegistration(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		RoundTripJSON(t, data, &SignedValidatorRegistration{})
	})
}

func FuzzRoundTripBuilderBid(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		RoundTripSSZ(t, data, &BuilderBid{})
		RoundTripJSON(t, data, &BuilderBid{})
	})
}

func FuzzRoundTripSignedBuilderBid(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		RoundTripSSZ(t, data, &SignedBuilderBid{})
		RoundTripJSON(t, data, &SignedBuilderBid{})
	})
}

func FuzzRoundTripGetHeaderResponse(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		RoundTripJSON(t, data, &GetHeaderResponse{})
	})
}

func FuzzRoundTripSignedBlindedBeaconBlock(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		RoundTripJSON(t, data, &SignedBlindedBeaconBlock{})
	})
}

func FuzzRoundTripGetPayloadResponse(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		RoundTripJSON(t, data, &GetPayloadResponse{})
	})
}

func FuzzRoundTripTransactions(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		RoundTripSSZ(t, data, &Transactions{})
		RoundTripJSON(t, data, &Transactions{})
	})
}

func FuzzRoundTripBuilderGetValidatorsResponseEntry(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		RoundTripJSON(t, data, &BuilderGetValidatorsResponseEntry{})
	})
}

func FuzzRoundTripBidTrace(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		RoundTripSSZ(t, data, &BidTrace{})
		RoundTripJSON(t, data, &BidTrace{})
	})
}

func FuzzRoundTripSignedBidTrace(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		RoundTripJSON(t, data, &SignedBidTrace{})
	})
}

func FuzzRoundTripBuilderSubmitBlockResponseMessage(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		RoundTripSSZ(t, data, &BuilderSubmitBlockResponseMessage{})
		RoundTripJSON(t, data, &BuilderSubmitBlockResponseMessage{})
	})
}

func FuzzRoundTripBuilderSubmitBlockResponse(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		RoundTripJSON(t, data, &BuilderSubmitBlockResponse{})
	})
}

func FuzzPayloadToPayloadHeader(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var payloadHeader ExecutionPayload
		if !Fill(data, &payloadHeader) {
			return
		}
		PayloadToPayloadHeader(&payloadHeader)
	})
}
