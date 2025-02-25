package transaction

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
)

type Transaction struct {
	Sender    string  `json:"sender"`
	Recipient string  `json:"recipient"`
	Amount    float64 `json:"amount"`
	Signature string  `json:"signature"`
}

// NewTransaction cria uma nova transação
func NewTransaction(sender, recipient string, amount float64) *Transaction {
	return &Transaction{
		Sender:    sender,
		Recipient: recipient,
		Amount:    amount,
	}
}

// SignTransaction assina uma transação com a chave privada do remetente
func (t *Transaction) SignTransaction(privateKey *ecdsa.PrivateKey) error {
	transactionData, err := json.Marshal(t)
	if err != nil {
		return err
	}
	hash := sha256.Sum256(transactionData)
	signature, err := ecdsa.SignASN1(rand.Reader, privateKey, hash[:])
	if err != nil {
		return err
	}
	t.Signature = hex.EncodeToString(signature)
	return nil
}

// ValidateTransaction valida a assinatura de uma transação
func (t *Transaction) ValidateTransaction(publicKey *ecdsa.PublicKey) bool {
	transactionData, err := json.Marshal(t)
	if err != nil {
		return false
	}
	hash := sha256.Sum256(transactionData)
	signature, err := hex.DecodeString(t.Signature)
	if err != nil {
		return false
	}
	return ecdsa.VerifyASN1(publicKey, hash[:], signature)
}
