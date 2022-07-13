package types

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common/hexutil"
)

var (
	ErrLength = fmt.Errorf("incorrect byte length")
)

type Signature [96]byte

func (s Signature) MarshalText() ([]byte, error) {
	return hexutil.Bytes(s[:]).MarshalText()
}

func (s *Signature) UnmarshalJSON(input []byte) error {
	b := hexutil.Bytes(s[:])
	err := b.UnmarshalJSON(input)
	if err != nil {
		return err
	}
	return s.FromSlice(b)
}

func (s *Signature) UnmarshalText(input []byte) error {
	b := hexutil.Bytes(s[:])
	err := b.UnmarshalText(input)
	if err != nil {
		return err
	}
	return s.FromSlice(b)
}

func (s Signature) String() string {
	return hexutil.Bytes(s[:]).String()
}

func (s *Signature) FromSlice(x []byte) error {
	if len(x) != 96 {
		return ErrLength
	}
	copy(s[:], x)
	return nil
}

type PublicKey [48]byte

func (p PublicKey) MarshalText() ([]byte, error) {
	return hexutil.Bytes(p[:]).MarshalText()
}

func (p *PublicKey) UnmarshalJSON(input []byte) error {
	b := hexutil.Bytes(p[:])
	b.UnmarshalJSON(input)
	return p.FromSlice(b)
}

func (p *PublicKey) UnmarshalText(input []byte) error {
	b := hexutil.Bytes(p[:])
	b.UnmarshalText(input)
	return p.FromSlice(b)
}

func (p PublicKey) String() string {
	return hexutil.Bytes(p[:]).String()
}

func (p *PublicKey) FromSlice(x []byte) error {
	if len(x) != 48 {
		return ErrLength
	}
	copy(p[:], x)
	return nil
}

type Address [20]byte

func (a Address) MarshalText() ([]byte, error) {
	return hexutil.Bytes(a[:]).MarshalText()
}

func (a *Address) UnmarshalJSON(input []byte) error {
	b := hexutil.Bytes(a[:])
	b.UnmarshalJSON(input)
	return a.FromSlice(b)
}

func (a *Address) UnmarshalText(input []byte) error {
	b := hexutil.Bytes(a[:])
	b.UnmarshalText(input)
	return a.FromSlice(b)
}

func (a Address) String() string {
	return hexutil.Bytes(a[:]).String()
}

func (a *Address) FromSlice(x []byte) error {
	if len(x) != 20 {
		return ErrLength
	}
	copy(a[:], x)
	return nil
}

type Hash [32]byte
type Root = Hash

func (h Hash) MarshalText() ([]byte, error) {
	return hexutil.Bytes(h[:]).MarshalText()
}

func (h *Hash) UnmarshalJSON(input []byte) error {
	b := hexutil.Bytes(h[:])
	b.UnmarshalJSON(input)
	return h.FromSlice(b)
}

func (h *Hash) UnmarshalText(input []byte) error {
	b := hexutil.Bytes(h[:])
	b.UnmarshalText(input)
	return h.FromSlice(b)
}

func (h *Hash) FromSlice(x []byte) error {
	if len(x) != 32 {
		return ErrLength
	}
	copy(h[:], x)
	return nil
}

func (h Hash) String() string {
	return hexutil.Bytes(h[:]).String()
}

type CommitteeBits [64]byte

func (c CommitteeBits) MarshalText() ([]byte, error) {
	return hexutil.Bytes(c[:]).MarshalText()
}

func (c *CommitteeBits) UnmarshalJSON(input []byte) error {
	b := hexutil.Bytes(c[:])
	b.UnmarshalJSON(input)
	return c.FromSlice(b)
}

func (c *CommitteeBits) UnmarshalText(input []byte) error {
	b := hexutil.Bytes(c[:])
	b.UnmarshalText(input)
	return c.FromSlice(b)

}

func (c CommitteeBits) String() string {
	return hexutil.Bytes(c[:]).String()
}

func (c *CommitteeBits) FromSlice(x []byte) error {
	if len(x) != 64 {
		return ErrLength
	}
	copy(c[:], x)
	return nil
}

type Bloom [256]byte

func (b Bloom) MarshalText() ([]byte, error) {
	return hexutil.Bytes(b[:]).MarshalText()
}

func (b *Bloom) UnmarshalJSON(input []byte) error {
	buf := hexutil.Bytes(b[:])
	buf.UnmarshalJSON(input)
	return b.FromSlice(buf)
}

func (b *Bloom) UnmarshalText(input []byte) error {
	buf := hexutil.Bytes(b[:])
	buf.UnmarshalText(input)
	return b.FromSlice(buf)
}

func (b Bloom) String() string {
	return hexutil.Bytes(b[:]).String()
}

func (b *Bloom) FromSlice(x []byte) error {
	if len(b) != 256 {
		return ErrLength
	}
	copy(b[:], x)
	return nil
}

type U256Str Hash // encodes/decodes to string, not hex

func reverse(src []byte) []byte {
	dst := make([]byte, len(src))
	copy(dst, src)
	for i := len(dst)/2 - 1; i >= 0; i-- {
		opp := len(dst) - 1 - i
		dst[i], dst[opp] = dst[opp], dst[i]
	}
	return dst
}

func (n U256Str) MarshalText() ([]byte, error) {
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
	copy(n[:], reverse(x.FillBytes(n[:])))
	return nil
}

type ExtraData []byte

func (e ExtraData) MarshalText() ([]byte, error) {
	return hexutil.Bytes(e).MarshalText()
}

func (e *ExtraData) UnmarshalJSON(input []byte) error {
	var buf hexutil.Bytes
	buf.UnmarshalJSON(input)
	return e.FromSlice(buf)
}

func (e *ExtraData) UnmarshalText(input []byte) error {
	var buf hexutil.Bytes
	buf.UnmarshalText(input)
	return e.FromSlice(buf)
}

func (e ExtraData) String() string {
	return hexutil.Bytes(e).String()
}

func (e *ExtraData) FromSlice(x []byte) error {
	if len(x) > 32 {
		return ErrLength
	}
	tmp := make([]byte, len(x))
	copy(tmp, x)
	*e = tmp
	return nil
}
