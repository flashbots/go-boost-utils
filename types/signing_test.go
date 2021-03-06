package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/stretchr/testify/require"

	"github.com/flashbots/go-boost-utils/bls"
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

func genValidatorRegistration(t require.TestingT, domain Domain) *SignedValidatorRegistration {
	sk, pk, err := bls.GenerateNewKeypair()
	require.NoError(t, err)

	var pubKey PublicKey
	pubKey.FromSlice(pk.Compress())

	msg := &RegisterValidatorRequestMessage{
		FeeRecipient: Address{0x42},
		GasLimit:     15_000_000,
		Timestamp:    1652369368,
		Pubkey:       pubKey,
	}

	signature, err := SignMessage(msg, domain, sk)
	require.NoError(t, err)
	return &SignedValidatorRegistration{
		Message:   msg,
		Signature: signature,
	}
}

func BenchmarkSignatureVerification(b *testing.B) {
	domain := ComputeDomain(DomainType{0x01, 0x00, 0x00, 0x00}, ForkVersion{}, Root{})
	reg := genValidatorRegistration(b, domain)

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		ok, err := VerifySignature(reg.Message, domain, reg.Message.Pubkey[:], reg.Signature[:])
		require.NoError(b, err)
		require.True(b, ok)
	}
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

func bytesTo4(bytes []byte) (res [4]byte) {
	copy(res[:], bytes[:4])
	return
}

func TestComputeDomainVector(t *testing.T) {
	for _, tc := range []struct {
		DomainType            string
		ForkVersion           string
		GenesisValidatorsRoot string
		ExpectedDomain        string
	}{
		{"0x07000000", "0x01000000", "0x0a08c27fe4ece2483f9e581f78c66379a06f96e9c24cd1390594ff939b26f95b", "0x07000000b503183cf3d26841cf4499d79f4387520811f5ed97776f0d5317f086"},
	} {
		var genesisValidatorsRoot Root
		genesisValidatorsRoot.FromSlice(hexutil.MustDecode(tc.GenesisValidatorsRoot))
		var expectedDomain [32]byte
		copy(expectedDomain[:], hexutil.MustDecode(tc.ExpectedDomain)[:32])
		require.Equal(t, expectedDomain, ComputeDomain(bytesTo4(hexutil.MustDecode(tc.DomainType)), bytesTo4(hexutil.MustDecode(tc.ForkVersion)), genesisValidatorsRoot))
	}
}

func _ComputeDomain(domainType DomainType, forkVersionHex string, genesisValidatorsRootHex string) (domain Domain, err error) {
	forkVersionBytes, err := hexutil.Decode(forkVersionHex)
	if err != nil || len(forkVersionBytes) > 4 {
		err = errors.New("invalid fork version passed")
		return domain, err
	}
	var forkVersion [4]byte
	copy(forkVersion[:], forkVersionBytes[:4])

	genesisValidatorsRoot := Root(common.HexToHash(genesisValidatorsRootHex))
	return ComputeDomain(domainType, forkVersion, genesisValidatorsRoot), nil
}

func Test_ComputeDomain(t *testing.T) {
	builderDomainKiln, err := _ComputeDomain(DomainTypeAppBuilder, GenesisForkVersionKiln, Root{}.String())
	require.NoError(t, err)
	require.Equal(t, "0x000000017acd69a9ede79f3eb3eaa814c09159eeaa3d004be62f3372d9b31e9c", hexutil.Encode(builderDomainKiln[:]))

	beaconProposerDomainKiln, err := _ComputeDomain(DomainTypeBeaconProposer, BellatrixForkVersionKiln, GenesisValidatorsRootKiln)
	require.NoError(t, err)
	require.Equal(t, "0x00000000e7acb21061790987fa1c1e745cccfb358370b33e8af2b2c18938e6c2", hexutil.Encode(beaconProposerDomainKiln[:]))
}

func TestKilnSignedBlindedBeaconBlockSignature(t *testing.T) {
	jsonFile, err := os.Open("../testdata/kiln-signedBlindedBeaconBlock-899730.json")
	require.NoError(t, err)
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	require.NoError(t, err)

	payload := new(SignedBlindedBeaconBlock)

	dec := json.NewDecoder(bytes.NewReader(byteValue))
	dec.DisallowUnknownFields()
	err = dec.Decode(payload)
	require.NoError(t, err)

	root, err := payload.Message.HashTreeRoot()
	require.NoError(t, err)
	require.Equal(t, "4ada338ce48190c9d2169ca2484de507140b1bfeff2a8f820611dedcdc102d63", common.Bytes2Hex(root[:]))

	pk, err := HexToPubkey("0xa04fe993de82bc878039bba5212a9fa750abf2293195cd55cbbce4827f56799cc67b5f66cf33bb1cec92dabcbcc0a0a9")
	require.NoError(t, err)
	require.Equal(t, "0xa04fe993de82bc878039bba5212a9fa750abf2293195cd55cbbce4827f56799cc67b5f66cf33bb1cec92dabcbcc0a0a9", pk.String())

	domain, err := _ComputeDomain(DomainTypeBeaconProposer, BellatrixForkVersionKiln, GenesisValidatorsRootKiln)
	require.NoError(t, err)
	ok, err := VerifySignature(payload.Message, domain, pk[:], payload.Signature[:])
	require.NoError(t, err)
	require.True(t, ok)
}
