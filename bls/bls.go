package bls

import (
	"crypto/rand"
	"errors"

	bls12381 "github.com/kilic/bls12-381"
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
	PublicKey = bls12381.PointG1
	SecretKey = bls12381.Fr
	Signature = bls12381.PointG2
)

var (
	ErrInvalidPubkeyLength    = errors.New("invalid pubkey length")
	ErrInvalidSecretKeyLength = errors.New("invalid secret key length")
	ErrInvalidSignatureLength = errors.New("invalid signature length")
	ErrSecretKeyIsZero        = errors.New("invalid secret key is zero")
)

func PublicKeyToBytes(pk *PublicKey) []byte {
	return bls12381.NewG1().ToCompressed(pk)
}

func SecretKeyToBytes(sk *SecretKey) []byte {
	return sk.ToBytes()
}

func SignatureToBytes(sig *Signature) []byte {
	return bls12381.NewG2().ToCompressed(sig)
}

func PublicKeyFromBytes(pkBytes []byte) (*PublicKey, error) {
	if len(pkBytes) != PublicKeyLength {
		return nil, ErrInvalidPubkeyLength
	}
	return bls12381.NewG1().FromCompressed(pkBytes)
}

func PublicKeyFromSecretKey(sk *SecretKey) (*PublicKey, error) {
	if sk.IsZero() {
		return nil, ErrSecretKeyIsZero
	}
	pk := new(PublicKey)
	g1 := bls12381.NewG1()
	g1.MulScalar(pk, &bls12381.G1One, sk)
	return pk, nil
}

func SecretKeyFromBytes(skBytes []byte) (*SecretKey, error) {
	if len(skBytes) != SecretKeyLength {
		return nil, ErrInvalidSecretKeyLength
	}
	sk := bls12381.NewFr().FromBytes(skBytes)
	if sk.IsZero() {
		return nil, ErrSecretKeyIsZero
	}
	return sk, nil
}

func GenerateRandomSecretKey() (*SecretKey, error) {
	return new(SecretKey).Rand(rand.Reader)
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
	g2 := bls12381.NewG2()
	Q, err := g2.HashToCurve(msg, domain)
	if err != nil {
		panic(err)
	}
	var signature bls12381.PointG2
	g2.MulScalar(&signature, Q, sk)
	return &signature
}

func SignatureFromBytes(sigBytes []byte) (*Signature, error) {
	if len(sigBytes) != SignatureLength {
		return nil, ErrInvalidSignatureLength
	}
	return bls12381.NewG2().FromCompressed(sigBytes)
}

func VerifySignatureBytes(msg, sigBytes, pkBytes []byte) (bool, error) {
	xP, err := bls12381.NewG1().FromCompressed(pkBytes)
	if err != nil {
		return false, err
	}
	Q, err := bls12381.NewG2().HashToCurve(msg, domain)
	if err != nil {
		return false, err
	}
	R, err := bls12381.NewG2().FromCompressed(sigBytes)
	if err != nil {
		return false, err
	}
	P := &bls12381.G1One

	pairingEngine := bls12381.NewEngine()
	pairingEngine.AddPair(xP, Q)
	pairingEngine.AddPairInv(P, R)
	return pairingEngine.Check(), nil
}
