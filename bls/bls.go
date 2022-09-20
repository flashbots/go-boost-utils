package bls

import (
	"crypto/rand"
	"errors"

	blst "github.com/supranational/blst/bindings/go"
)

// From https://github.com/supranational/blst/tree/master/bindings/go

var dst = []byte("BLS_SIG_BLS12381G2_XMD:SHA-256_SSWU_RO_POP_")

const BLSPublicKeyLength int = 48
const BLSSecretKeyLength int = 32
const BLSSignatureLength int = 96

type PublicKey = blst.P1Affine
type SecretKey = blst.SecretKey
type Signature = blst.P2Affine

var (
	ErrDeserializeSecretKey   = errors.New("could not deserialize secret key from bytes")
	ErrInvalidPubkey          = errors.New("invalid pubkey")
	ErrInvalidPubkeyLength    = errors.New("invalid pubkey length")
	ErrInvalidSecretKeyLength = errors.New("invalid secret key length")
	ErrInvalidSignature       = errors.New("invalid signature")
	ErrInvalidSignatureLength = errors.New("invalid signature length")
	ErrUncompressPubkey       = errors.New("could not uncompress public key from bytes")
	ErrUncompressSignature    = errors.New("could not uncompress signature from bytes")
)

func PublicKeyFromBytes(pkBytes []byte) (*PublicKey, error) {
	if len(pkBytes) != BLSPublicKeyLength {
		return nil, ErrInvalidPubkeyLength
	}

	pk := new(PublicKey).Uncompress(pkBytes)
	if pk == nil {
		return nil, ErrUncompressPubkey
	}

	if !pk.KeyValidate() {
		return nil, ErrInvalidPubkey
	}

	return pk, nil
}

func PublicKeyFromSecretKey(sk *SecretKey) *PublicKey {
	return new(PublicKey).From(sk)
}

func SecretKeyFromBytes(skBytes []byte) (*SecretKey, error) {
	if len(skBytes) != BLSSecretKeyLength {
		return nil, ErrInvalidSecretKeyLength
	}
	secretKey := new(SecretKey).Deserialize(skBytes)
	if secretKey == nil {
		return nil, ErrDeserializeSecretKey
	}
	return secretKey, nil
}

func GenerateRandomSecretKey() (*SecretKey, error) {
	var ikm [BLSSecretKeyLength]byte
	_, err := rand.Read(ikm[:])
	if err != nil {
		return nil, err
	}
	sk := blst.KeyGen(ikm[:])
	return sk, nil
}

func GenerateNewKeypair() (*SecretKey, *PublicKey, error) {
	sk, err := GenerateRandomSecretKey()
	if err != nil {
		return nil, nil, err
	}
	return sk, PublicKeyFromSecretKey(sk), nil
}

func Sign(sk *SecretKey, msg []byte) *Signature {
	return new(Signature).Sign(sk, msg, dst)
}

func SignatureFromBytes(sigBytes []byte) (*Signature, error) {
	if len(sigBytes) != BLSSignatureLength {
		return nil, ErrInvalidSignatureLength
	}

	sig := new(Signature).Uncompress(sigBytes)
	if sig == nil {
		return nil, ErrUncompressSignature
	}

	if !sig.SigValidate(false) {
		return nil, ErrInvalidSignature
	}
	return sig, nil
}

func VerifySignature(sig *Signature, pk *PublicKey, msg []byte) bool {
	return sig.Verify(true, pk, false, msg, dst)
}

func VerifySignatureBytes(msg, sigBytes, pkBytes []byte) (bool, error) {
	sig, err := SignatureFromBytes(sigBytes)
	if err != nil {
		return false, err
	}
	pubkey, err := PublicKeyFromBytes(pkBytes)
	if err != nil {
		return false, err
	}
	return VerifySignature(sig, pubkey, msg[:]), nil
}
