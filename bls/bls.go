package bls

import (
	"errors"
	"math/big"

	bls12381 "github.com/consensys/gnark-crypto/ecc/bls12-381"
	"github.com/consensys/gnark-crypto/ecc/bls12-381/fr"
)

const (
	PublicKeyLength = bls12381.SizeOfG1AffineCompressed
	SecretKeyLength = fr.Bytes
	SignatureLength = bls12381.SizeOfG2AffineCompressed
)

type (
	PublicKey = bls12381.G1Affine
	SecretKey = fr.Element
	Signature = bls12381.G2Affine
)

var (
	g1OneJac, _, g1One, _     = bls12381.Generators()
	domain                    = []byte("BLS_SIG_BLS12381G2_XMD:SHA-256_SSWU_RO_POP_")
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

func SignatureFromBytes(sigBytes []byte) (*Signature, error) {
	if len(sigBytes) != SignatureLength {
		return nil, ErrInvalidSignatureLength
	}
	sig := new(Signature)
	err := sig.Unmarshal(sigBytes)
	return sig, err
}

func GenerateRandomSecretKey() (*SecretKey, error) {
	return new(SecretKey).SetRandom()
}

func PublicKeyFromSecretKey(sk *SecretKey) (*PublicKey, error) {
	if sk.IsZero() {
		return nil, ErrSecretKeyIsZero
	}
	skBigInt := new(big.Int)
	sk.ToBigIntRegular(skBigInt)
	pkJac := new(bls12381.G1Jac).ScalarMultiplication(&g1OneJac, skBigInt)
	return new(bls12381.G1Affine).FromJacobian(pkJac), nil
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
	sk.ToBigIntRegular(skBigInt)
	QJac := new(bls12381.G2Jac).FromAffine(&Q)
	sigJac := new(bls12381.G2Jac).ScalarMultiplication(QJac, skBigInt)
	return new(bls12381.G2Affine).FromJacobian(sigJac)
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
