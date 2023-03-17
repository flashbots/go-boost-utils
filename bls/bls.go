package bls

import (
	"errors"
	"math/big"

	bls12381 "github.com/consensys/gnark-crypto/ecc/bls12-381"
	"github.com/consensys/gnark-crypto/ecc/bls12-381/fr"
)

// Heavily inspired by:
// https://github.com/protolambda/bls12-381-util/blob/master/signatures.go
// Thank you for the excellent code.

var domain = []byte("BLS_SIG_BLS12381G2_XMD:SHA-256_SSWU_RO_POP_")

const (
	PublicKeyLength int = 48
	SecretKeyLength int = 32
	SignatureLength int = 96
)

type (
	PublicKey = bls12381.G1Affine
	SecretKey = fr.Element
	Signature = bls12381.G2Affine
)

var (
	_, _, g1One, _            = bls12381.Generators()
	ErrInvalidPubkeyLength    = errors.New("invalid public key length")
	ErrInvalidSecretKeyLength = errors.New("invalid secret key length")
	ErrInvalidSignatureLength = errors.New("invalid signature length")
	ErrSecretKeyIsZero        = errors.New("invalid secret key is zero")
)

func PublicKeyToBytes(pk *PublicKey) []byte {
	pkBytes := pk.Bytes()
	return pkBytes[:]
}

func SecretKeyToBytes(sk *SecretKey) []byte {
	skBytes := sk.Bytes()
	return skBytes[:]
}

func SignatureToBytes(sig *Signature) []byte {
	sigBytes := sig.Bytes()
	return sigBytes[:]
}

func PublicKeyFromBytes(pkBytes []byte) (*PublicKey, error) {
	if len(pkBytes) != PublicKeyLength {
		return nil, ErrInvalidPubkeyLength
	}
	pk := new(PublicKey)
	err := pk.Unmarshal(pkBytes)
	return pk, err
}

func PublicKeyFromSecretKey(sk *SecretKey) (*PublicKey, error) {
	if sk.IsZero() {
		return nil, ErrSecretKeyIsZero
	}
	skBigInt := new(big.Int)
	sk.BigInt(skBigInt)
	return new(bls12381.G1Affine).ScalarMultiplication(&g1One, skBigInt), nil
}

func SecretKeyFromBytes(skBytes []byte) (*SecretKey, error) {
	if len(skBytes) != SecretKeyLength {
		return nil, ErrInvalidSecretKeyLength
	}
	sk := new(SecretKey).SetBytes(skBytes)
	if sk.IsZero() {
		return nil, ErrSecretKeyIsZero
	}
	return sk, nil
}

func GenerateRandomSecretKey() (*SecretKey, error) {
	return new(SecretKey).SetRandom()
}

func GenerateNewKeypair() (*SecretKey, *PublicKey, error) {
	sk, err := GenerateRandomSecretKey()
	if err != nil {
		return nil, nil, err
	}
	pk, err := PublicKeyFromSecretKey(sk)
	if err != nil {
		return nil, nil, err
	}
	return sk, pk, nil
}

func Sign(sk *SecretKey, msg []byte) *Signature {
	Q, err := bls12381.HashToG2(msg, domain)
	if err != nil {
		panic(err)
	}
	skBigInt := new(big.Int)
	sk.BigInt(skBigInt)
	signature := new(bls12381.G2Affine)
	signature.ScalarMultiplication(&Q, skBigInt)
	return signature
}

func SignatureFromBytes(sigBytes []byte) (*Signature, error) {
	if len(sigBytes) != SignatureLength {
		return nil, ErrInvalidSignatureLength
	}
	sig := new(Signature)
	err := sig.Unmarshal(sigBytes)
	return sig, err
}

func VerifySignature(sig *Signature, pk *PublicKey, msg []byte) (bool, error) {
	Q, err := bls12381.HashToG2(msg, domain)
	if err != nil {
		return false, err
	}
	var negP bls12381.G1Affine
	negP.Neg(&g1One)
	return bls12381.PairingCheck(
		[]bls12381.G1Affine{*pk, negP},
		[]bls12381.G2Affine{Q, *sig},
	)
}

func VerifySignatureBytes(msg, sigBytes, pkBytes []byte) (bool, error) {
	pk, err := PublicKeyFromBytes(pkBytes)
	if err != nil {
		return false, err
	}
	sig, err := SignatureFromBytes(sigBytes)
	if err != nil {
		return false, err
	}
	return VerifySignature(sig, pk, msg)
}
