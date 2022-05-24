package types

import "bytes"

// Cmp compares one U256Str to another and returns an integer indicating whether a > b.
// The result will be 0 if a == b, -1 if a < b, and +1 if a > b.
func (n *U256Str) Cmp(b *U256Str) int {
	return bytes.Compare(n[:], b[:])
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
