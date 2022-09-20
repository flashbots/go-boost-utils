package types

import (
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/trie"
)

var ErrInvalidTransaction = errors.New("invalid transaction")

func CalculateHash(payload *ExecutionPayload) (Hash, error) {
	header, err := ExecutionPayloadToHeader(payload)
	if err != nil {
		return Hash{}, err
	}
	return Hash(header.Hash()), nil
}

func ExecutionPayloadToHeader(payload *ExecutionPayload) (*types.Header, error) {
	transactionData := make([]*types.Transaction, len(payload.Transactions))
	for i, encTx := range payload.Transactions {
		var tx types.Transaction
		err := tx.UnmarshalBinary(encTx)
		if err != nil {
			return nil, ErrInvalidTransaction
		}
		transactionData[i] = &tx
	}

	return &types.Header{
		ParentHash:  common.Hash(payload.ParentHash),
		UncleHash:   types.EmptyUncleHash,
		Coinbase:    common.Address(payload.FeeRecipient),
		Root:        common.Hash(payload.StateRoot),
		TxHash:      types.DeriveSha(types.Transactions(transactionData), trie.NewStackTrie(nil)),
		ReceiptHash: common.Hash(payload.ReceiptsRoot),
		Bloom:       types.Bloom(payload.LogsBloom),
		Difficulty:  common.Big0,
		Number:      new(big.Int).SetUint64(payload.BlockNumber),
		GasLimit:    payload.GasLimit,
		GasUsed:     payload.GasUsed,
		Time:        payload.Timestamp,
		Extra:       payload.ExtraData,
		MixDigest:   common.Hash(payload.Random),
		BaseFee:     payload.BaseFeePerGas.BigInt(),
	}, nil
}
