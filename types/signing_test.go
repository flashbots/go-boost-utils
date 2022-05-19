package types

import (
	"testing"

	"github.com/flashbots/go-boost-utils/bls"
	"github.com/stretchr/testify/require"
)

func TestVerifySignature(t *testing.T) {
	sk, pk, err := bls.GenerateNewKeypair()
	require.NoError(t, err)
	msg := &RegisterValidatorRequestMessage{
		FeeRecipient: Address{0x42},
		GasLimit:     15_000_000,
		Timestamp:    1652369368,
		Pubkey:       PublicKey{0x0d},
	}
	domain := ComputeDomain(DomainType{0x01, 0x00, 0x00, 0x00}, ForkVersion{}, Root{})
	root, err := ComputeSigningRoot(msg, domain)
	require.NoError(t, err)

	sig := bls.Sign(sk, root[:])
	sig2, err := SignMessage(msg, domain, sk)
	require.NoError(t, err)
	require.Equal(t, sig.Compress(), sig2[:])

	ok, err := VerifySignature(msg, domain, pk.Compress(), sig.Compress())
	require.NoError(t, err)
	require.True(t, ok)
}

func TestVerifySignatureManualPk(t *testing.T) {
	msg2 := RegisterValidatorRequestMessage{
		FeeRecipient: Address{0x42},
		GasLimit:     15_000_000,
		Timestamp:    1652369368,
		Pubkey:       PublicKey{0x0d},
	}
	root2, err := msg2.HashTreeRoot()
	require.NoError(t, err)

	// Verify expected signature with manual pk
	pkBytes := make([]byte, 32)
	pkBytes[0] = 0x01
	sk2, err := bls.SecretKeyFromBytes(pkBytes)
	require.NoError(t, err)
	sig2 := bls.Sign(sk2, root2[:]).Compress()
	var signature2 Signature
	signature2.FromSlice(sig2)
	require.Equal(t, "0x8e09a0ae7af113da2043001cc19fb1b3b24bbe022c1b8050ba2297ad1186f4217dd7095edad1d16d83d10f3297883d9e1674c81da95f10d3358c5afdb2500279e720b32879219c9a3b33415239bf46a66cd92b9d1750a6dd7cc7ec936a357128", signature2.String())
}
