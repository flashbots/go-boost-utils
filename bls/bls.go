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

// PublicKeyFromBytes creates a BLS public key from a  BigEndian byte slice.
func PublicKeyFromBytes(pkBytes []byte) (*PublicKey, error) {
	if len(pkBytes) != BLSPublicKeyLength {
		return nil, errors.New("invalid pubkey length")
	}

	pk := new(blst.P1Affine).Uncompress(pkBytes)
	if pk == nil {
		return nil, errors.New("could not uncompress public key from bytes")
	}

	if !pk.KeyValidate() {
		return nil, errors.New("invalid pubkey")
	}

	return pk, nil
}

func PublicKeyFromSecretKey(sk *SecretKey) *PublicKey {
	return new(PublicKey).From(sk)
}

func SecretKeyFromBytes(skBytes []byte) (*SecretKey, error) {
	if len(skBytes) != BLSSecretKeyLength {
		return nil, errors.New("invalid secret key length")
	}
	secretKey := new(blst.SecretKey).Deserialize(skBytes)
	if secretKey == nil {
		return nil, errors.New("could not deserialize secret key from bytes")
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

func GenerateNewKeypair() (*PublicKey, *SecretKey, error) {
	sk, err := GenerateRandomSecretKey()
	if err != nil {
		return nil, nil, err
	}
	return PublicKeyFromSecretKey(sk), sk, nil
}

func Sign(sk *SecretKey, msg []byte) *Signature {
	return new(Signature).Sign(sk, msg, dst)
}

// SignatureFromBytes creates a BLS signature from a LittleEndian byte slice.
func SignatureFromBytes(sigBytes []byte) (*Signature, error) {
	if len(sigBytes) != BLSSignatureLength {
		return nil, errors.New("invalid signature length")
	}

	sig := new(blst.P2Affine).Uncompress(sigBytes)
	if sig == nil {
		return nil, errors.New("could not uncompress signature from bytes")
	}

	if !sig.SigValidate(false) {
		return nil, errors.New("invalid signature")
	}
	return sig, nil
}

func VerifySignature(sig *Signature, pk *PublicKey, msg []byte) bool {
	return sig.Verify(false, pk, false, msg, dst)
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
