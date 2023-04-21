package utils

import (
	"encoding/json"
	"errors"
	"io"

	"github.com/attestantio/go-eth2-client/spec/bellatrix"
	"github.com/attestantio/go-eth2-client/spec/phase0"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/flashbots/go-boost-utils/bls"
)

var ErrLength = errors.New("invalid length")

func BlsPublicKeyToPublicKey(blsPubKey *bls.PublicKey) (ret phase0.BLSPubKey, err error) {
	return HexToPubkey(hexutil.Encode(bls.PublicKeyToBytes(blsPubKey)))
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

	if err := decoder.Decode(dst); err != nil {
		return err
	}
	return nil
}
