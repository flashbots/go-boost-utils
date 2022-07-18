package types

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/flashbots/go-boost-utils/bls"
)

type PubkeyHex string

func NewPubkeyHex(pk string) PubkeyHex {
	return PubkeyHex(strings.ToLower(pk))
}

func (p PublicKey) PubkeyHex() PubkeyHex {
	return NewPubkeyHex(p.String())
}

func (p PubkeyHex) String() string {
	return string(p)
}

func IntToU256(i uint64) (ret U256Str) {
	s := fmt.Sprint(i)
	ret.UnmarshalText([]byte(s))
	return
}

func (n *U256Str) BigInt() *big.Int {
	return new(big.Int).SetBytes(reverse(n[:]))
}

// Cmp compares one U256Str to another and returns an integer indicating whether a > b.
// The result will be 0 if a == b, -1 if a < b, and +1 if a > b.
func (n *U256Str) Cmp(b *U256Str) int {
	_a := n.BigInt()
	_b := b.BigInt()
	return _a.Cmp(_b)
}

func BlsPublicKeyToPublicKey(blsPubKey *bls.PublicKey) (ret PublicKey) {
	ret.FromSlice(blsPubKey.Compress())
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
