package types

import (
	"fmt"
)

// IntToU256 takes a uint64 and returns a U256 type
func IntToU256(i uint64) (ret U256Str) {
	s := fmt.Sprint(i)
	ret.UnmarshalText([]byte(s))
	return
}

// HexToAddress takes a hex string and returns an Address
func HexToAddress(s string) (ret Address, err error) {
	err = ret.UnmarshalText([]byte(s))
	return ret, err
}

// HexToPubkey takes a hex string and returns a PublicKey
func HexToPubkey(s string) (ret PublicKey, err error) {
	err = ret.UnmarshalText([]byte(s))
	return ret, err
}

// HexToSignature takes a hex string and returns a Signature
func HexToSignature(s string) (ret Signature, err error) {
	err = ret.UnmarshalText([]byte(s))
	return ret, err
}
