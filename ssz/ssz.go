package ssz

import (
	"errors"

	"github.com/attestantio/go-eth2-client/spec/phase0"
	"github.com/flashbots/go-boost-utils/bls"
)

var (
	ErrLength = errors.New("invalid length")

	DomainBuilder = ComputeDomain(DomainTypeAppBuilder, phase0.Version{}, phase0.Root{})

	DomainTypeBeaconProposer = phase0.DomainType{0x00, 0x00, 0x00, 0x00}
	DomainTypeBlobSidecar    = phase0.DomainType{0x0B, 0x00, 0x00, 0x00}
	DomainTypeAppBuilder     = phase0.DomainType{0x00, 0x00, 0x00, 0x01}
)

type ObjWithHashTreeRoot interface {
	HashTreeRoot() ([32]byte, error)
}

func ComputeDomain(dt phase0.DomainType, forkVersion phase0.Version, genesisValidatorsRoot phase0.Root) [32]byte {
	forkDataRoot, _ := (&phase0.ForkData{
		CurrentVersion:        forkVersion,
		GenesisValidatorsRoot: genesisValidatorsRoot,
	}).HashTreeRoot()

	var domain [32]byte
	copy(domain[0:4], dt[:])
	copy(domain[4:], forkDataRoot[0:28])

	return domain
}

func ComputeSigningRoot(obj ObjWithHashTreeRoot, d phase0.Domain) ([32]byte, error) {
	var zero [32]byte
	root, err := obj.HashTreeRoot()
	if err != nil {
		return zero, err
	}
	signingData := phase0.SigningData{ObjectRoot: root, Domain: d}
	msg, err := signingData.HashTreeRoot()
	if err != nil {
		return zero, err
	}
	return msg, nil
}

func SignMessage(obj ObjWithHashTreeRoot, d phase0.Domain, sk *bls.SecretKey) (phase0.BLSSignature, error) {
	root, err := ComputeSigningRoot(obj, d)
	if err != nil {
		return phase0.BLSSignature{}, err
	}

	signatureBytes := bls.SignatureToBytes(bls.Sign(sk, root[:]))

	var signature phase0.BLSSignature
	if len(signatureBytes) != 96 {
		return phase0.BLSSignature{}, ErrLength
	}

	copy(signature[:], signatureBytes)
	if err != nil {
		return phase0.BLSSignature{}, err
	}

	return signature, nil
}

func VerifySignature(obj ObjWithHashTreeRoot, d phase0.Domain, pkBytes, sigBytes []byte) (bool, error) {
	msg, err := ComputeSigningRoot(obj, d)
	if err != nil {
		return false, err
	}

	return bls.VerifySignatureBytes(msg[:], sigBytes, pkBytes)
}

func VerifySignatureRoot(root phase0.Root, d phase0.Domain, pkBytes, sigBytes []byte) (bool, error) {
	signingData := phase0.SigningData{ObjectRoot: root, Domain: d}
	msg, err := signingData.HashTreeRoot()
	if err != nil {
		return false, err
	}

	return bls.VerifySignatureBytes(msg[:], sigBytes, pkBytes)
}
