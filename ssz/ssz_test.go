package ssz

import (
	"bytes"
	"errors"
	"io"
	"os"
	"strings"
	"testing"
	"time"

	builderApiBellatrix "github.com/attestantio/go-builder-client/api/bellatrix"
	builderApiV1 "github.com/attestantio/go-builder-client/api/v1"
	apiV1Bellatrix "github.com/attestantio/go-eth2-client/api/v1/bellatrix"
	"github.com/attestantio/go-eth2-client/spec/bellatrix"
	"github.com/attestantio/go-eth2-client/spec/phase0"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/flashbots/go-boost-utils/bls"
	"github.com/flashbots/go-boost-utils/types"
	"github.com/flashbots/go-boost-utils/utils"
	"github.com/stretchr/testify/require"
)

var ErrInvalidForkVersion = errors.New("invalid fork version passed")

func TestVerifySignature(t *testing.T) {
	sk, pk, err := bls.GenerateNewKeypair()
	require.NoError(t, err)
	msg := &builderApiV1.ValidatorRegistration{
		FeeRecipient: bellatrix.ExecutionAddress{0x42},
		GasLimit:     15_000_000,
		Timestamp:    time.Now(),
		Pubkey:       phase0.BLSPubKey{0x0d},
	}
	domain := ComputeDomain(phase0.DomainType{0x01, 0x00, 0x00, 0x00}, phase0.Version{}, phase0.Root{})
	root, err := ComputeSigningRoot(msg, domain)
	require.NoError(t, err)

	sig := bls.Sign(sk, root[:])
	sig2, err := SignMessage(msg, domain, sk)
	require.NoError(t, err)
	require.Equal(t, bls.SignatureToBytes(sig), sig2[:])

	ok, err := VerifySignature(msg, domain, bls.PublicKeyToBytes(pk), bls.SignatureToBytes(sig))
	require.NoError(t, err)
	require.True(t, ok)
}

func TestVerifySignatureRoot(t *testing.T) {
	sk, pk, err := bls.GenerateNewKeypair()
	require.NoError(t, err)
	msg := &builderApiV1.ValidatorRegistration{
		FeeRecipient: bellatrix.ExecutionAddress{0x42},
		GasLimit:     15_000_000,
		Timestamp:    time.Now(),
		Pubkey:       phase0.BLSPubKey{0x0d},
	}
	domain := ComputeDomain(phase0.DomainType{0x01, 0x00, 0x00, 0x00}, phase0.Version{}, phase0.Root{})
	root, err := ComputeSigningRoot(msg, domain)
	require.NoError(t, err)

	sig := bls.Sign(sk, root[:])
	sig2, err := SignMessage(msg, domain, sk)
	require.NoError(t, err)
	require.Equal(t, bls.SignatureToBytes(sig), sig2[:])
	root2, err := msg.HashTreeRoot()
	require.NoError(t, err)

	ok, err := VerifySignatureRoot(root2, domain, bls.PublicKeyToBytes(pk), bls.SignatureToBytes(sig))
	require.NoError(t, err)
	require.True(t, ok)
}

func genValidatorRegistration(t require.TestingT, domain phase0.Domain) *builderApiV1.SignedValidatorRegistration {
	sk, pk, err := bls.GenerateNewKeypair()
	require.NoError(t, err)

	var pubKey phase0.BLSPubKey
	pubKeyBytes := bls.PublicKeyToBytes(pk)
	copy(pubKey[:], pubKeyBytes)

	msg := &builderApiV1.ValidatorRegistration{
		FeeRecipient: bellatrix.ExecutionAddress{0x42},
		GasLimit:     15_000_000,
		Timestamp:    time.Now(),
		Pubkey:       pubKey,
	}

	signature, err := SignMessage(msg, domain, sk)
	require.NoError(t, err)
	return &builderApiV1.SignedValidatorRegistration{
		Message:   msg,
		Signature: signature,
	}
}

func BenchmarkSignatureVerification(b *testing.B) {
	domain := ComputeDomain(phase0.DomainType{0x01, 0x00, 0x00, 0x00}, phase0.Version{}, phase0.Root{})
	reg := genValidatorRegistration(b, domain)

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		ok, err := VerifySignature(reg.Message, domain, reg.Message.Pubkey[:], reg.Signature[:])
		require.NoError(b, err)
		require.True(b, ok)
	}
}

func TestVerifySignatureManualPk(t *testing.T) {
	msg := &builderApiV1.ValidatorRegistration{
		FeeRecipient: bellatrix.ExecutionAddress{0x42},
		GasLimit:     15_000_000,
		Timestamp:    time.Unix(1652369368, 0),
		Pubkey:       phase0.BLSPubKey{0x0d},
	}
	root2, err := msg.HashTreeRoot()
	require.NoError(t, err)

	// Verify expected signature with manual pk
	pkBytes := make([]byte, 32)
	pkBytes[0] = 0x01
	sk, err := bls.SecretKeyFromBytes(pkBytes)
	require.NoError(t, err)

	sig := bls.SignatureToBytes(bls.Sign(sk, root2[:]))
	sigHex := hexutil.Encode(sig)
	require.Equal(t, "0x8e09a0ae7af113da2043001cc19fb1b3b24bbe022c1b8050ba2297ad1186f4217dd7095edad1d16d83d10f3297883d9e1674c81da95f10d3358c5afdb2500279e720b32879219c9a3b33415239bf46a66cd92b9d1750a6dd7cc7ec936a357128", sigHex)
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
		var genesisValidatorsRoot phase0.Root
		genesisValidatorsRootBytes := hexutil.MustDecode(tc.GenesisValidatorsRoot)
		copy(genesisValidatorsRoot[:], genesisValidatorsRootBytes[:])
		var expectedDomain [32]byte
		copy(expectedDomain[:], hexutil.MustDecode(tc.ExpectedDomain)[:32])
		require.Equal(t, expectedDomain, ComputeDomain(bytesTo4(hexutil.MustDecode(tc.DomainType)), bytesTo4(hexutil.MustDecode(tc.ForkVersion)), genesisValidatorsRoot))
	}
}

func _ComputeDomain(domainType phase0.DomainType, forkVersionHex, genesisValidatorsRootHex string) (domain phase0.Domain, err error) {
	forkVersionBytes, err := hexutil.Decode(forkVersionHex)
	if err != nil || len(forkVersionBytes) > 4 {
		return domain, ErrInvalidForkVersion
	}
	var forkVersion [4]byte
	copy(forkVersion[:], forkVersionBytes[:4])

	genesisValidatorsRoot := phase0.Root(common.HexToHash(genesisValidatorsRootHex))
	return ComputeDomain(domainType, forkVersion, genesisValidatorsRoot), nil
}

func Test_ComputeDomain(t *testing.T) {
	builderDomainKiln, err := _ComputeDomain(DomainTypeAppBuilder, types.GenesisForkVersionKiln, phase0.Root{}.String())
	require.NoError(t, err)
	require.Equal(t, "0x000000017acd69a9ede79f3eb3eaa814c09159eeaa3d004be62f3372d9b31e9c", hexutil.Encode(builderDomainKiln[:]))

	beaconProposerDomainKiln, err := _ComputeDomain(DomainTypeBeaconProposer, types.BellatrixForkVersionKiln, types.GenesisValidatorsRootKiln)
	require.NoError(t, err)
	require.Equal(t, "0x00000000e7acb21061790987fa1c1e745cccfb358370b33e8af2b2c18938e6c2", hexutil.Encode(beaconProposerDomainKiln[:]))
}

func TestVerifySignedBuilderBidSignature(t *testing.T) {
	// SignedBuilderBid from Kiln
	bidStr := `{"message":{"header":{"parent_hash":"0x0544e2170998060d9560fdbf8f263a08c0a209211569a0560138522b84805abc","fee_recipient":"0x0000000000000000000000000000000000000000","state_root":"0xcded53d652660a91bfe6f5dfb017204a4cdd1598a07116b2cdea1586d603d01c","receipts_root":"0x56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421","logs_bloom":"0x00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000","prev_randao":"0xd60955dc7f0cc7bf28d7e6c6f4859081f3a6df5ef4f70e05d70d8282bac20c6c","block_number":"960335","gas_limit":"30000000","gas_used":"0","timestamp":"1659720144","extra_data":"0x466c617368626f747320666c617368626c6f636b","base_fee_per_gas":"7","block_hash":"0xea33078b00e6b2926f45ed6d3190a3a6ada75cee342f600cf22fa02a9a2edcb7","transactions_root":"0x7ffe241ea60187fdb0187bfa22de35d1f9bed7ab061d9401fd47e34a54fbede1"},"value":"0","pubkey":"0xb5246e299aeb782fbc7c91b41b3284245b1ed5206134b0028b81dfb974e5900616c67847c2354479934fc4bb75519ee1"},"signature":"0xa775df980d589a87b234cf36b94fbcd40540ab1dffb752a013c02f636d85db60023f7e9d883de8cfdbfd94e0e3b598c01429fee50a5cb8d9fce50557baec2e9f81268f14f4f044b44b1238b7945201f036036d1a25d60e681f3737d4ef3b54b6"}`

	// Decode the bid
	bid := new(builderApiBellatrix.SignedBuilderBid)
	require.NoError(t, utils.DecodeJSON(strings.NewReader(bidStr), bid))

	// Verify signature
	builderDomainKiln, err := _ComputeDomain(DomainTypeAppBuilder, types.GenesisForkVersionKiln, phase0.Root{}.String())
	require.NoError(t, err)
	ok, err := VerifySignature(bid.Message, builderDomainKiln, bid.Message.Pubkey[:], bid.Signature[:])
	require.NoError(t, err)
	require.True(t, ok)
}

func TestVerifySignedBuilderBidRootSignature(t *testing.T) {
	// SignedBuilderBid from Kiln
	bidStr := `{"message":{"header":{"parent_hash":"0x0544e2170998060d9560fdbf8f263a08c0a209211569a0560138522b84805abc","fee_recipient":"0x0000000000000000000000000000000000000000","state_root":"0xcded53d652660a91bfe6f5dfb017204a4cdd1598a07116b2cdea1586d603d01c","receipts_root":"0x56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421","logs_bloom":"0x00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000","prev_randao":"0xd60955dc7f0cc7bf28d7e6c6f4859081f3a6df5ef4f70e05d70d8282bac20c6c","block_number":"960335","gas_limit":"30000000","gas_used":"0","timestamp":"1659720144","extra_data":"0x466c617368626f747320666c617368626c6f636b","base_fee_per_gas":"7","block_hash":"0xea33078b00e6b2926f45ed6d3190a3a6ada75cee342f600cf22fa02a9a2edcb7","transactions_root":"0x7ffe241ea60187fdb0187bfa22de35d1f9bed7ab061d9401fd47e34a54fbede1"},"value":"0","pubkey":"0xb5246e299aeb782fbc7c91b41b3284245b1ed5206134b0028b81dfb974e5900616c67847c2354479934fc4bb75519ee1"},"signature":"0xa775df980d589a87b234cf36b94fbcd40540ab1dffb752a013c02f636d85db60023f7e9d883de8cfdbfd94e0e3b598c01429fee50a5cb8d9fce50557baec2e9f81268f14f4f044b44b1238b7945201f036036d1a25d60e681f3737d4ef3b54b6"}`

	// Decode the bid
	bid := new(builderApiBellatrix.SignedBuilderBid)
	require.NoError(t, utils.DecodeJSON(strings.NewReader(bidStr), bid))

	// Verify signature
	builderDomainKiln, err := _ComputeDomain(DomainTypeAppBuilder, types.GenesisForkVersionKiln, phase0.Root{}.String())
	require.NoError(t, err)
	root, err := bid.Message.HashTreeRoot()
	require.NoError(t, err)
	ok, err := VerifySignatureRoot(root, builderDomainKiln, bid.Message.Pubkey[:], bid.Signature[:])
	require.NoError(t, err)
	require.True(t, ok)
}

func TestKilnSignedBlindedBeaconBlockSignature(t *testing.T) {
	jsonFile, err := os.Open("../testdata/kiln-signedBlindedBeaconBlock-899730.json")
	require.NoError(t, err)
	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)
	require.NoError(t, err)

	payload := new(apiV1Bellatrix.SignedBlindedBeaconBlock)
	require.NoError(t, utils.DecodeJSON(bytes.NewReader(byteValue), payload))

	root, err := payload.Message.HashTreeRoot()
	require.NoError(t, err)
	require.Equal(t, "4ada338ce48190c9d2169ca2484de507140b1bfeff2a8f820611dedcdc102d63", common.Bytes2Hex(root[:]))

	pk, err := utils.HexToPubkey("0xa04fe993de82bc878039bba5212a9fa750abf2293195cd55cbbce4827f56799cc67b5f66cf33bb1cec92dabcbcc0a0a9")
	require.NoError(t, err)
	require.Equal(t, "0xa04fe993de82bc878039bba5212a9fa750abf2293195cd55cbbce4827f56799cc67b5f66cf33bb1cec92dabcbcc0a0a9", pk.String())

	domain, err := _ComputeDomain(DomainTypeBeaconProposer, types.BellatrixForkVersionKiln, types.GenesisValidatorsRootKiln)
	require.NoError(t, err)
	ok, err := VerifySignature(payload.Message, domain, pk[:], payload.Signature[:])
	require.NoError(t, err)
	require.True(t, ok)
}

func TestKilnSignedBlindedBeaconBlockSignature2(t *testing.T) {
	jsonFile, err := os.Open("../testdata/kiln-signedBlindedBeaconBlock-899730.json")
	require.NoError(t, err)
	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)
	require.NoError(t, err)

	payload := new(apiV1Bellatrix.SignedBlindedBeaconBlock)
	require.NoError(t, utils.DecodeJSON(bytes.NewReader(byteValue), payload))

	proof := make([][]byte, 33)
	for i := 0; i < 33; i++ {
		hash := phase0.Hash32{0x01}
		proof[i] = hash[:]
	}

	withdrawalCredentials := [32]byte{0x03}
	payload.Message.Body.Deposits = []*phase0.Deposit{
		{
			Proof: proof,
			Data: &phase0.DepositData{
				PublicKey:             phase0.BLSPubKey{0x02},
				WithdrawalCredentials: withdrawalCredentials[:],
				Amount:                1234,
				Signature:             phase0.BLSSignature{0x04},
			},
		},
	}

	root, err := payload.Message.HashTreeRoot()
	require.NoError(t, err)
	htrHex := common.Bytes2Hex(root[:])
	require.Equal(t, "da469dcc55560d3f8ae26ea6c3910efce3e3b1c4cecc988c3ebafe71e81ad077", htrHex, htrHex)
}
