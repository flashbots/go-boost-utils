package types

import (
	"encoding/json"
	"testing"
)

func FuzzRoundTripEth1Data(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		RoundTripSSZ(t, data, &Eth1Data{})
		RoundTripJSON(t, data, &Eth1Data{})
	})
}

func FuzzUnmarshalEth1Data(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var value Eth1Data
		_ = json.Unmarshal(data, &value)
		_ = value.UnmarshalSSZ(data)
	})
}

func FuzzRoundTripBeaconBlockHeader(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		RoundTripSSZ(t, data, &BeaconBlockHeader{})
		RoundTripJSON(t, data, &BeaconBlockHeader{})
	})
}

func FuzzUnmarshalBeaconBlockHeader(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var value BeaconBlockHeader
		_ = json.Unmarshal(data, &value)
		_ = value.UnmarshalSSZ(data)
	})
}

func FuzzRoundTripSignedBeaconBlockHeader(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		RoundTripSSZ(t, data, &SignedBeaconBlockHeader{})
		RoundTripJSON(t, data, &SignedBeaconBlockHeader{})
	})
}

func FuzzUnmarshalSignedBeaconBlockHeader(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var value SignedBeaconBlockHeader
		_ = json.Unmarshal(data, &value)
		_ = value.UnmarshalSSZ(data)
	})
}

func FuzzRoundTripProposerSlashing(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		RoundTripSSZ(t, data, &ProposerSlashing{})
		RoundTripJSON(t, data, &ProposerSlashing{})
	})
}

func FuzzUnmarshalProposerSlashing(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var value ProposerSlashing
		_ = json.Unmarshal(data, &value)
		_ = value.UnmarshalSSZ(data)
	})
}

func FuzzRoundTripCheckpoint(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		RoundTripSSZ(t, data, &Checkpoint{})
		RoundTripJSON(t, data, &Checkpoint{})
	})
}

func FuzzUnmarshalCheckpoint(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var value Checkpoint
		_ = json.Unmarshal(data, &value)
		_ = value.UnmarshalSSZ(data)
	})
}

func FuzzRoundTripAttestationData(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		RoundTripSSZ(t, data, &AttestationData{})
		RoundTripJSON(t, data, &AttestationData{})
	})
}

func FuzzUnmarshalAttestationData(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var value AttestationData
		_ = json.Unmarshal(data, &value)
		_ = value.UnmarshalSSZ(data)
	})
}

func FuzzRoundTripIndexedAttestation(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		RoundTripSSZ(t, data, &IndexedAttestation{})
		RoundTripJSON(t, data, &IndexedAttestation{})
	})
}

func FuzzUnmarshalIndexedAttestation(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var value IndexedAttestation
		_ = json.Unmarshal(data, &value)
		_ = value.UnmarshalSSZ(data)
	})
}

func FuzzRoundTripAttesterSlashing(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		RoundTripSSZ(t, data, &AttesterSlashing{})
		RoundTripJSON(t, data, &AttesterSlashing{})
	})
}

func FuzzUnmarshalAttesterSlashing(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var value AttesterSlashing
		_ = json.Unmarshal(data, &value)
		_ = value.UnmarshalSSZ(data)
	})
}

func FuzzRoundTripAttestation(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		RoundTripSSZ(t, data, &Attestation{})
		RoundTripJSON(t, data, &Attestation{})
	})
}

func FuzzUnmarshalAttestation(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var value Attestation
		_ = json.Unmarshal(data, &value)
		_ = value.UnmarshalSSZ(data)
	})
}

func FuzzRoundTripDeposit(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		RoundTripSSZ(t, data, &Deposit{})
		RoundTripJSON(t, data, &Deposit{})
	})
}

func FuzzUnmarshalDeposit(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var value Deposit
		_ = json.Unmarshal(data, &value)
		_ = value.UnmarshalSSZ(data)
	})
}

func FuzzRoundTripSyncAggregate(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		RoundTripSSZ(t, data, &SyncAggregate{})
		RoundTripJSON(t, data, &SyncAggregate{})
	})
}

func FuzzUnmarshalSyncAggregate(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var value SyncAggregate
		_ = json.Unmarshal(data, &value)
		_ = value.UnmarshalSSZ(data)
	})
}

func FuzzRoundTripVoluntaryExit(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		RoundTripSSZ(t, data, &VoluntaryExit{})
		RoundTripJSON(t, data, &VoluntaryExit{})
	})
}

func FuzzUnmarshalVoluntaryExit(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var value VoluntaryExit
		_ = json.Unmarshal(data, &value)
		_ = value.UnmarshalSSZ(data)
	})
}

func FuzzRoundTripExecutionPayloadHeader(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		RoundTripSSZ(t, data, &ExecutionPayloadHeader{})
		RoundTripJSON(t, data, &ExecutionPayloadHeader{})
	})
}

func FuzzUnmarshalExecutionPayloadHeader(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var value ExecutionPayloadHeader
		_ = json.Unmarshal(data, &value)
		_ = value.UnmarshalSSZ(data)
	})
}

func FuzzRoundTripExecutionPayload(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		RoundTripJSON(t, data, &ExecutionPayload{})
	})
}

func FuzzUnmarshalExecutionPayload(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var value ExecutionPayload
		_ = json.Unmarshal(data, &value)
	})
}

func FuzzRoundTripBlindedBeaconBlockBody(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		RoundTripSSZ(t, data, &BlindedBeaconBlockBody{})
		RoundTripJSON(t, data, &BlindedBeaconBlockBody{})
	})
}

func FuzzUnmarshalBlindedBeaconBlockBody(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var value BlindedBeaconBlockBody
		_ = json.Unmarshal(data, &value)
		_ = value.UnmarshalSSZ(data)
	})
}

func FuzzRoundTripBeaconBlockBody(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		RoundTripJSON(t, data, &BeaconBlockBody{})
	})
}

func FuzzUnmarshalBeaconBlockBody(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var value BeaconBlockBody
		_ = json.Unmarshal(data, &value)
	})
}

func FuzzRoundTripBlindedBeaconBlock(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		RoundTripSSZ(t, data, &BlindedBeaconBlock{})
		RoundTripJSON(t, data, &BlindedBeaconBlock{})
	})
}

func FuzzUnmarshalBlindedBeaconBlock(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var value BlindedBeaconBlock
		_ = json.Unmarshal(data, &value)
		_ = value.UnmarshalSSZ(data)
	})
}

func FuzzRoundTripBeaconBlock(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		RoundTripJSON(t, data, &BeaconBlock{})
	})
}

func FuzzUnmarshalBeaconBlock(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var value BeaconBlock
		_ = json.Unmarshal(data, &value)
	})
}

func FuzzRoundTripRegisterValidatorRequestMessage(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		RoundTripSSZ(t, data, &RegisterValidatorRequestMessage{})
		RoundTripJSON(t, data, &RegisterValidatorRequestMessage{})
	})
}

func FuzzUnmarshalRegisterValidatorRequestMessage(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var value RegisterValidatorRequestMessage
		_ = json.Unmarshal(data, &value)
		_ = value.UnmarshalSSZ(data)
	})
}

func FuzzRoundTripSignedValidatorRegistration(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		RoundTripJSON(t, data, &SignedValidatorRegistration{})
	})
}

func FuzzUnmarshalSignedValidatorRegistration(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var value SignedValidatorRegistration
		_ = json.Unmarshal(data, &value)
	})
}

func FuzzRoundTripBuilderBid(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		RoundTripSSZ(t, data, &BuilderBid{})
		RoundTripJSON(t, data, &BuilderBid{})
	})
}

func FuzzUnmarshalBuilderBid(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var value BuilderBid
		_ = json.Unmarshal(data, &value)
		_ = value.UnmarshalSSZ(data)
	})
}

func FuzzRoundTripSignedBuilderBid(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		RoundTripSSZ(t, data, &SignedBuilderBid{})
		RoundTripJSON(t, data, &SignedBuilderBid{})
	})
}

func FuzzUnmarshalSignedBuilderBid(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var value SignedBuilderBid
		_ = json.Unmarshal(data, &value)
		_ = value.UnmarshalSSZ(data)
	})
}

func FuzzRoundTripGetHeaderResponse(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		RoundTripJSON(t, data, &GetHeaderResponse{})
	})
}

func FuzzUnmarshalGetHeaderResponse(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var value GetHeaderResponse
		_ = json.Unmarshal(data, &value)
	})
}

func FuzzRoundTripSignedBlindedBeaconBlock(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		RoundTripJSON(t, data, &SignedBlindedBeaconBlock{})
	})
}

func FuzzUnmarshalSignedBlindedBeaconBlock(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var value SignedBlindedBeaconBlock
		_ = json.Unmarshal(data, &value)
	})
}

func FuzzRoundTripGetPayloadResponse(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		RoundTripJSON(t, data, &GetPayloadResponse{})
	})
}

func FuzzUnmarshalGetPayloadResponse(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var value GetPayloadResponse
		_ = json.Unmarshal(data, &value)
	})
}

func FuzzRoundTripTransactions(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		RoundTripSSZ(t, data, &Transactions{})
		RoundTripJSON(t, data, &Transactions{})
	})
}

func FuzzUnmarshalTransactions(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var value Transactions
		_ = json.Unmarshal(data, &value)
		_ = value.UnmarshalSSZ(data)
	})
}

func FuzzRoundTripBuilderGetValidatorsResponseEntry(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		RoundTripJSON(t, data, &BuilderGetValidatorsResponseEntry{})
	})
}

func FuzzUnmarshalBuilderGetValidatorsResponseEntry(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var value BuilderGetValidatorsResponseEntry
		_ = json.Unmarshal(data, &value)
	})
}

func FuzzRoundTripBidTrace(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		RoundTripSSZ(t, data, &BidTrace{})
		RoundTripJSON(t, data, &BidTrace{})
	})
}

func FuzzUnmarshalBidTrace(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var value BidTrace
		_ = json.Unmarshal(data, &value)
		_ = value.UnmarshalSSZ(data)
	})
}

func FuzzRoundTripSignedBidTrace(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		RoundTripJSON(t, data, &SignedBidTrace{})
	})
}

func FuzzUnmarshalSignedBidTrace(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var value SignedBidTrace
		_ = json.Unmarshal(data, &value)
	})
}

func FuzzRoundTripBuilderSubmitBlockRequest(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		RoundTripJSON(t, data, &BuilderSubmitBlockRequest{})
	})
}

func FuzzUnmarshalBuilderSubmitBlockRequest(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var value BuilderSubmitBlockRequest
		_ = json.Unmarshal(data, &value)
	})
}

func FuzzRoundTripBuilderSubmitBlockResponseMessage(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		RoundTripSSZ(t, data, &BuilderSubmitBlockResponseMessage{})
		RoundTripJSON(t, data, &BuilderSubmitBlockResponseMessage{})
	})
}

func FuzzUnmarshalBuilderSubmitBlockResponseMessage(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var value BuilderSubmitBlockResponseMessage
		_ = json.Unmarshal(data, &value)
		_ = value.UnmarshalSSZ(data)
	})
}

func FuzzRoundTripBuilderSubmitBlockResponse(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		RoundTripJSON(t, data, &BuilderSubmitBlockResponse{})
	})
}

func FuzzUnmarshalBuilderSubmitBlockResponse(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var value BuilderSubmitBlockResponse
		_ = json.Unmarshal(data, &value)
	})
}

func FuzzPayloadToPayloadHeader(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var payloadHeader ExecutionPayload
		if !Fill(data, &payloadHeader) {
			return
		}
		_, _ = PayloadToPayloadHeader(&payloadHeader)
	})
}
