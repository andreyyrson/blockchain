package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
)

// Wallet
type Wallet struct {
	PrivateKey *ecdsa.PrivateKey
	PublicKey  *ecdsa.PublicKey
	Address    string
}

func NewWallet() *Wallet {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		panic(err)
	}

	publicKey := &privateKey.PublicKey
	address := generateAddress(publicKey)

	return &Wallet{
		PrivateKey: privateKey,
		PublicKey:  publicKey,
		Address:    address,
	}
}

func generateAddress(publicKey *ecdsa.PublicKey) string {
	publicKeyBytes := elliptic.Marshal(publicKey.Curve, publicKey.X, publicKey.Y)
	hash := sha256.Sum256(publicKeyBytes)
	return hex.EncodeToString(hash[:])
}

func (w *Wallet) SignData(data []byte) (string, error) {
	hash := sha256.Sum256(data)
	signature, err := ecdsa.SignASN1(rand.Reader, w.PrivateKey, hash[:])
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(signature), nil
}

func VerifySignature(publicKey *ecdsa.PublicKey, data []byte, signature string) bool {
	signatureBytes, err := hex.DecodeString(signature)
	if err != nil {
		return false
	}
	hash := sha256.Sum256(data)
	return ecdsa.VerifyASN1(publicKey, hash[:], signatureBytes)
}
