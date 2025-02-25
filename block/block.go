package block

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"
	"time"

	"github.com/andreyyrson/blockchain/transaction"
)

type Block struct {
	Index        int
	Timestamp    string
	PrevHash     string
	Hash         string
	Transactions []*transaction.Transaction // Agora é um slice de ponteiros
	Nonce        int
}

// NewBlock cria um novo bloco
func NewBlock(index int, prevHash string, transactions []*transaction.Transaction) *Block {
	block := &Block{
		Index:        index,
		Timestamp:    time.Now().String(),
		PrevHash:     prevHash,
		Transactions: transactions,
		Nonce:        0,
	}
	block.Hash = block.CalculateHash()
	return block
}

// CalculateHash calcula o hash do bloco
func (b *Block) CalculateHash() string {
	blockData := fmt.Sprintf("%d%s%s%v%d", b.Index, b.Timestamp, b.PrevHash, b.Transactions, b.Nonce)
	hash := sha256.Sum256([]byte(blockData))
	return hex.EncodeToString(hash[:])
}

// MineBlock realiza a mineração do bloco
func (b *Block) MineBlock(difficulty int) {
	target := strings.Repeat("0", difficulty)
	for b.Hash[:difficulty] != target {
		b.Nonce++
		b.Hash = b.CalculateHash()
	}
	fmt.Printf("Bloco minerado: %s\n", b.Hash)
}

// ValidateBlock valida a integridade do bloco
func (b *Block) ValidateBlock() bool {
	return b.Hash == b.CalculateHash()
}
