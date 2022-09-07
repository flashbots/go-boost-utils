package types

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/trie"
)

func CalculateHash(payload *ExecutionPayload) Hash {
	return [32]byte(ExecutionPayloadToHeader(payload).Hash())
}

func ExecutionPayloadToHeader(payload *ExecutionPayload) *types.Header {
	transactionData := make([]*types.Transaction, len(payload.Transactions))
	for i, encTx := range payload.Transactions {
		var tx types.Transaction
		tx.UnmarshalBinary(encTx)
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
	}
}
