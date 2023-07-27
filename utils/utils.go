package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/big"

	"github.com/attestantio/go-builder-client/api"
	"github.com/attestantio/go-eth2-client/spec"
	"github.com/attestantio/go-eth2-client/spec/bellatrix"
	"github.com/attestantio/go-eth2-client/spec/capella"
	"github.com/attestantio/go-eth2-client/spec/deneb"
	"github.com/attestantio/go-eth2-client/spec/phase0"
	utilbellatrix "github.com/attestantio/go-eth2-client/util/bellatrix"
	utilcapella "github.com/attestantio/go-eth2-client/util/capella"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/trie"
	"github.com/flashbots/go-boost-utils/bls"
)

var (
	ErrInvalidTransaction = errors.New("invalid transaction")
	ErrLength             = errors.New("invalid length")
	ErrNilPayload         = errors.New("nil payload")
	ErrUnsupportedVersion = errors.New("unsupported version")
	ErrUnknownVersion     = errors.New("unknown version")
)

func BlsPublicKeyToPublicKey(blsPubKey *bls.PublicKey) (ret phase0.BLSPubKey, err error) {
	return HexToPubkey(hexutil.Encode(bls.PublicKeyToBytes(blsPubKey)))
}

func BlsSignatureToSignature(blsSignature *bls.Signature) (ret phase0.BLSSignature, err error) {
	return HexToSignature(hexutil.Encode(bls.SignatureToBytes(blsSignature)))
}

// HexToHash takes a hex string and returns a Hash
func HexToHash(s string) (ret phase0.Hash32, err error) {
	bytes, err := hexutil.Decode(s)
	if len(bytes) != len(ret) {
		return phase0.Hash32{}, ErrLength
	}
	copy(ret[:], bytes)
	return
}

// HexToAddress takes a hex string and returns an Address
func HexToAddress(s string) (ret bellatrix.ExecutionAddress, err error) {
	bytes, err := hexutil.Decode(s)
	if len(bytes) != len(ret) {
		return bellatrix.ExecutionAddress{}, ErrLength
	}
	copy(ret[:], bytes)
	return
}

// HexToPubkey takes a hex string and returns a PublicKey
func HexToPubkey(s string) (ret phase0.BLSPubKey, err error) {
	bytes, err := hexutil.Decode(s)
	if len(bytes) != len(ret) {
		return phase0.BLSPubKey{}, ErrLength
	}
	copy(ret[:], bytes)
	return
}

// HexToSignature takes a hex string and returns a Signature
func HexToSignature(s string) (ret phase0.BLSSignature, err error) {
	bytes, err := hexutil.Decode(s)
	if len(bytes) != len(ret) {
		return phase0.BLSSignature{}, ErrLength
	}
	copy(ret[:], bytes)
	return
}

// DecodeJSON decodes a JSON string into a struct while disallowing unknown fields
func DecodeJSON(r io.Reader, dst any) error {
	decoder := json.NewDecoder(r)
	decoder.DisallowUnknownFields()
	return decoder.Decode(dst)
}

// PayloadToPayloadHeader converts an ExecutionPayload to ExecutionPayloadHeader
func PayloadToPayloadHeader(payload *api.VersionedExecutionPayload) (*api.VersionedExecutionPayloadHeader, error) {
	if payload == nil {
		return nil, ErrNilPayload
	}

	switch payload.Version {
	case spec.DataVersionCapella:
		header, err := capellaPayloadToPayloadHeader(payload.Capella)
		if err != nil {
			return nil, err
		}

		return &api.VersionedExecutionPayloadHeader{
			Version: spec.DataVersionCapella,
			Capella: header,
		}, nil
	case spec.DataVersionDeneb:
		header, err := denebPayloadToPayloadHeader(payload.Deneb)
		if err != nil {
			return nil, err
		}

		return &api.VersionedExecutionPayloadHeader{
			Version: spec.DataVersionDeneb,
			Deneb:   header,
		}, nil
	case spec.DataVersionPhase0, spec.DataVersionAltair, spec.DataVersionBellatrix:
		return nil, fmt.Errorf("%w: %d", ErrUnsupportedVersion, payload.Version)
	default:
		return nil, fmt.Errorf("%w: %d", ErrUnknownVersion, payload.Version)
	}
}

func capellaPayloadToPayloadHeader(payload *capella.ExecutionPayload) (*capella.ExecutionPayloadHeader, error) {
	if payload == nil {
		return nil, ErrNilPayload
	}

	txRoot, err := deriveTransactionsRoot(payload.Transactions)
	if err != nil {
		return nil, fmt.Errorf("failed to derive transactions root: %w", err)
	}

	wdRoot, err := deriveWithdrawalsRoot(payload.Withdrawals)
	if err != nil {
		return nil, fmt.Errorf("failed to derive withdrawals root: %w", err)
	}

	return &capella.ExecutionPayloadHeader{
		ParentHash:       payload.ParentHash,
		FeeRecipient:     payload.FeeRecipient,
		StateRoot:        payload.StateRoot,
		ReceiptsRoot:     payload.ReceiptsRoot,
		LogsBloom:        payload.LogsBloom,
		PrevRandao:       payload.PrevRandao,
		BlockNumber:      payload.BlockNumber,
		GasLimit:         payload.GasLimit,
		GasUsed:          payload.GasUsed,
		Timestamp:        payload.Timestamp,
		ExtraData:        payload.ExtraData,
		BaseFeePerGas:    payload.BaseFeePerGas,
		BlockHash:        payload.BlockHash,
		TransactionsRoot: txRoot,
		WithdrawalsRoot:  wdRoot,
	}, nil
}

func denebPayloadToPayloadHeader(payload *deneb.ExecutionPayload) (*deneb.ExecutionPayloadHeader, error) {
	if payload == nil {
		return nil, ErrNilPayload
	}

	txRoot, err := deriveTransactionsRoot(payload.Transactions)
	if err != nil {
		return nil, fmt.Errorf("failed to derive transactions root: %w", err)
	}

	wdRoot, err := deriveWithdrawalsRoot(payload.Withdrawals)
	if err != nil {
		return nil, fmt.Errorf("failed to derive withdrawals root: %w", err)
	}

	return &deneb.ExecutionPayloadHeader{
		ParentHash:       payload.ParentHash,
		FeeRecipient:     payload.FeeRecipient,
		StateRoot:        payload.StateRoot,
		ReceiptsRoot:     payload.ReceiptsRoot,
		LogsBloom:        payload.LogsBloom,
		PrevRandao:       payload.PrevRandao,
		BlockNumber:      payload.BlockNumber,
		GasLimit:         payload.GasLimit,
		GasUsed:          payload.GasUsed,
		Timestamp:        payload.Timestamp,
		ExtraData:        payload.ExtraData,
		BaseFeePerGas:    payload.BaseFeePerGas,
		BlockHash:        payload.BlockHash,
		TransactionsRoot: txRoot,
		WithdrawalsRoot:  wdRoot,
		DataGasUsed:      payload.DataGasUsed,
		ExcessDataGas:    payload.ExcessDataGas,
	}, nil
}

func deriveTransactionsRoot(transactions []bellatrix.Transaction) (phase0.Root, error) {
	txs := utilbellatrix.ExecutionPayloadTransactions{Transactions: transactions}
	txRoot, err := txs.HashTreeRoot()
	if err != nil {
		return phase0.Root{}, err
	}
	return txRoot, nil
}

func deriveWithdrawalsRoot(withdrawals []*capella.Withdrawal) (phase0.Root, error) {
	wd := utilcapella.ExecutionPayloadWithdrawals{Withdrawals: withdrawals}
	wdRoot, err := wd.HashTreeRoot()
	if err != nil {
		return phase0.Root{}, err
	}
	return wdRoot, nil
}

// ComputeBlockHash computes the block hash for a given execution payload.
func ComputeBlockHash(payload *api.VersionedExecutionPayload) (phase0.Hash32, error) {
	switch payload.Version {
	case spec.DataVersionCapella:
		header, err := capellaExecutionPayloadToBlockHeader(payload.Capella)
		if err != nil {
			return phase0.Hash32{}, err
		}
		return phase0.Hash32(header.Hash()), nil
	case spec.DataVersionDeneb:
		header, err := denebExecutionPayloadToBlockHeader(payload.Deneb)
		if err != nil {
			return phase0.Hash32{}, err
		}
		return phase0.Hash32(header.Hash()), nil
	case spec.DataVersionPhase0, spec.DataVersionAltair, spec.DataVersionBellatrix:
		return phase0.Hash32{}, fmt.Errorf("%w: %d", ErrUnsupportedVersion, payload.Version)
	default:
		return phase0.Hash32{}, fmt.Errorf("%w: %d", ErrUnknownVersion, payload.Version)
	}
}

func capellaExecutionPayloadToBlockHeader(payload *capella.ExecutionPayload) (*types.Header, error) {
	transactionHash, err := deriveTransactionsHash(payload.Transactions)
	if err != nil {
		return nil, err
	}
	withdrawalsHash := deriveWithdrawalsHash(payload.Withdrawals)
	baseFeePerGas := deriveBaseFeePerGas(payload.BaseFeePerGas)
	return &types.Header{
		ParentHash:      common.Hash(payload.ParentHash),
		UncleHash:       types.EmptyUncleHash,
		Coinbase:        common.Address(payload.FeeRecipient),
		Root:            payload.StateRoot,
		TxHash:          transactionHash,
		ReceiptHash:     payload.ReceiptsRoot,
		Bloom:           payload.LogsBloom,
		Difficulty:      common.Big0,
		Number:          new(big.Int).SetUint64(payload.BlockNumber),
		GasLimit:        payload.GasLimit,
		GasUsed:         payload.GasUsed,
		Time:            payload.Timestamp,
		Extra:           payload.ExtraData,
		MixDigest:       payload.PrevRandao,
		BaseFee:         baseFeePerGas,
		WithdrawalsHash: &withdrawalsHash,
	}, nil
}

func denebExecutionPayloadToBlockHeader(payload *deneb.ExecutionPayload) (*types.Header, error) {
	transactionHash, err := deriveTransactionsHash(payload.Transactions)
	if err != nil {
		return nil, err
	}
	baseFeePerGas := payload.BaseFeePerGas.ToBig()
	withdrawalsHash := deriveWithdrawalsHash(payload.Withdrawals)
	return &types.Header{
		ParentHash:      common.Hash(payload.ParentHash),
		UncleHash:       types.EmptyUncleHash,
		Coinbase:        common.Address(payload.FeeRecipient),
		Root:            common.Hash(payload.StateRoot),
		TxHash:          transactionHash,
		ReceiptHash:     common.Hash(payload.ReceiptsRoot),
		Bloom:           payload.LogsBloom,
		Difficulty:      common.Big0,
		Number:          new(big.Int).SetUint64(payload.BlockNumber),
		GasLimit:        payload.GasLimit,
		GasUsed:         payload.GasUsed,
		Time:            payload.Timestamp,
		Extra:           payload.ExtraData,
		MixDigest:       payload.PrevRandao,
		BaseFee:         baseFeePerGas,
		WithdrawalsHash: &withdrawalsHash,
		ExcessDataGas:   new(big.Int).SetUint64(payload.ExcessDataGas),
	}, nil
}

func deriveTransactionsHash(transactions []bellatrix.Transaction) (common.Hash, error) {
	transactionData := make([]*types.Transaction, len(transactions))
	for i, encTx := range transactions {
		var tx types.Transaction

		if err := tx.UnmarshalBinary(encTx); err != nil {
			return common.Hash{}, ErrInvalidTransaction
		}
		transactionData[i] = &tx
	}
	return types.DeriveSha(types.Transactions(transactionData), trie.NewStackTrie(nil)), nil
}

func deriveWithdrawalsHash(withdrawals []*capella.Withdrawal) common.Hash {
	withdrawalData := make([]*types.Withdrawal, len(withdrawals))
	for i, w := range withdrawals {
		withdrawalData[i] = &types.Withdrawal{
			Index:     uint64(w.Index),
			Validator: uint64(w.ValidatorIndex),
			Address:   common.Address(w.Address),
			Amount:    uint64(w.Amount),
		}
	}
	return types.DeriveSha(types.Withdrawals(withdrawalData), trie.NewStackTrie(nil))
}

func deriveBaseFeePerGas(baseFeePerGas [32]byte) *big.Int {
	// base fee per gas is stored little-endian but we need it
	// big-endian for big.Int.
	var arr [32]byte
	for i := 0; i < 32; i++ {
		arr[i] = baseFeePerGas[32-1-i]
	}
	return new(big.Int).SetBytes(arr[:])
}
