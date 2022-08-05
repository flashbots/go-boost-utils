package types

import (
	"encoding/json"
	"fmt"
	"math/big"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/common/hexutil"
)

var (
	// GenesisValidatorsRootMainnet = ""
	GenesisValidatorsRootKiln    = "0x99b09fcd43e5905236c370f184056bec6e6638cfc31a323b304fc4aa789cb4ad"
	GenesisValidatorsRootRopsten = "0x44f1e56283ca88b35c789f7f449e52339bc1fefe3a45913a43a6d16edcd33cf1"
	GenesisValidatorsRootSepolia = "0xd8ea171f3c94aea21ebc42a1ed61052acf3f9209c00e4efbaaddac09ed9b8078"
	GenesisValidatorsRootGoerli  = "0x043db0d9a83813551ee2f33450d23797757d430911a9320530ad8a0eabc43efb"

	// GenesisForkVersionMainnet = "0x00000000"
	GenesisForkVersionKiln    = "0x70000069"
	GenesisForkVersionRopsten = "0x80000069"
	GenesisForkVersionSepolia = "0x90000069"
	GenesisForkVersionGoerli  = "0x00001020"

	// BellatrixForkVersionMainner = ""
	BellatrixForkVersionKiln    = "0x70000071"
	BellatrixForkVersionRopsten = "0x80000071"
	BellatrixForkVersionSepolia = "0x90000071"
	BellatrixForkVersionGoerli  = "0x02001020"

	ErrLength = fmt.Errorf("incorrect byte length")
)

type Uint64StringSlice []uint64

func (slice Uint64StringSlice) MarshalJSON() ([]byte, error) {
	values := make([]string, len(slice))
	for i, value := range []uint64(slice) {
		values[i] = fmt.Sprintf(`"%v"`, value)
	}

	return []byte(fmt.Sprintf("[%v]", strings.Join(values, ","))), nil
}

func (slice *Uint64StringSlice) UnmarshalJSON(b []byte) error {
	// Try array of strings first.
	var values []string
	err := json.Unmarshal(b, &values)
	if err != nil {
		// Fall back to array of integers:
		var values []uint64
		if err := json.Unmarshal(b, &values); err != nil {
			return err
		}
		*slice = values
		return nil
	}
	*slice = make([]uint64, len(values))
	for i, value := range values {
		value, err := strconv.ParseUint(value, 10, 64)
		if err != nil {
			return err
		}
		(*slice)[i] = value
	}
	return nil
}

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
	if err := b.UnmarshalJSON(input); err != nil {
		return err
	}
	return p.FromSlice(b)
}

func (p *PublicKey) UnmarshalText(input []byte) error {
	b := hexutil.Bytes(p[:])
	if err := b.UnmarshalText(input); err != nil {
		return err
	}
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
	if err := b.UnmarshalJSON(input); err != nil {
		return err
	}
	return a.FromSlice(b)
}

func (a *Address) UnmarshalText(input []byte) error {
	b := hexutil.Bytes(a[:])
	if err := b.UnmarshalText(input); err != nil {
		return err
	}
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
	if err := b.UnmarshalJSON(input); err != nil {
		return err
	}
	return h.FromSlice(b)
}

func (h *Hash) UnmarshalText(input []byte) error {
	b := hexutil.Bytes(h[:])
	if err := b.UnmarshalText(input); err != nil {
		return err
	}
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
	if err := b.UnmarshalJSON(input); err != nil {
		return err
	}
	return c.FromSlice(b)
}

func (c *CommitteeBits) UnmarshalText(input []byte) error {
	b := hexutil.Bytes(c[:])
	if err := b.UnmarshalText(input); err != nil {
		return err
	}
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
	if err := buf.UnmarshalJSON(input); err != nil {
		return err
	}
	return b.FromSlice(buf)
}

func (b *Bloom) UnmarshalText(input []byte) error {
	buf := hexutil.Bytes(b[:])
	if err := buf.UnmarshalText(input); err != nil {
		return err
	}
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
	if err := buf.UnmarshalJSON(input); err != nil {
		return err
	}
	return e.FromSlice(buf)
}

func (e *ExtraData) UnmarshalText(input []byte) error {
	var buf hexutil.Bytes
	if err := buf.UnmarshalText(input); err != nil {
		return err
	}
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

type VersionString string

func (s VersionString) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(s))
}

func (s *VersionString) UnmarshalJSON(b []byte) error {
	if len(b) >= 2 {
		return ErrLength
	}
	*s = VersionString(b[1 : len(b)-1])
	return nil
}
