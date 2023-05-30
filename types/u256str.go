package types

import (
	"errors"
	"fmt"
	"math/big"
)

var (
	ErrSign   = errors.New("invalid sign")
	ErrLength = errors.New("invalid length")
)

type U256Str [32]byte

func reverse(src []byte) []byte {
	dst := make([]byte, len(src))
	copy(dst, src)
	for i := len(dst)/2 - 1; i >= 0; i-- {
		opp := len(dst) - 1 - i
		dst[i], dst[opp] = dst[opp], dst[i]
	}
	return dst
}

func (n *U256Str) MarshalText() ([]byte, error) {
	return []byte(new(big.Int).SetBytes(reverse(n[:])).String()), nil
}

func (n *U256Str) UnmarshalJSON(input []byte) error {
	if len(input) < 2 {
		return ErrLength
	}
	x := new(big.Int)
	err := x.UnmarshalJSON(input[1 : len(input)-1])
	if err != nil {
		return err
	}
	return n.FromBig(x)
}

func (n *U256Str) UnmarshalText(input []byte) error {
	x := new(big.Int)
	err := x.UnmarshalText(input)
	if err != nil {
		return err
	}
	return n.FromBig(x)
}

func (n *U256Str) String() string {
	return new(big.Int).SetBytes(reverse(n[:])).String()
}

func (n *U256Str) FromSlice(x []byte) error {
	if len(x) > 32 {
		return ErrLength
	}
	copy(n[:], x)
	return nil
}

func (n *U256Str) FromBig(x *big.Int) error {
	if x.BitLen() > 256 {
		return ErrLength
	}
	if x.Sign() == -1 {
		return ErrSign
	}
	copy(n[:], reverse(x.FillBytes(n[:])))
	return nil
}

func IntToU256(i uint64) (ret U256Str) {
	s := fmt.Sprint(i)
	_ = ret.UnmarshalText([]byte(s))
	return ret
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
