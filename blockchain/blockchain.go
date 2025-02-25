package blockchain

import (
	"github.com/andreyyrson/blockchain/block"
	"github.com/andreyyrson/blockchain/transaction"
)

type Blockchain struct {
	Chain               []*block.Block
	PendingTransactions []*transaction.Transaction
	Difficulty          int
	MiningReward        float64
}

func NewBlockchain(difficulty int, miningReward float64) *Blockchain {
	genesisBlock := block.NewBlock(0, "", []*transaction.Transaction{})
	return &Blockchain{
		Chain:               []*block.Block{genesisBlock},
		PendingTransactions: []*transaction.Transaction{},
		Difficulty:          difficulty,
		MiningReward:        miningReward,
	}
}

func (bc *Blockchain) AddTransaction(tx *transaction.Transaction) {
	bc.PendingTransactions = append(bc.PendingTransactions, tx)
}

func (bc *Blockchain) MinePendingTransactions(miningRewardAddress string) {
	newBlock := block.NewBlock(len(bc.Chain), bc.Chain[len(bc.Chain)-1].Hash, bc.PendingTransactions)
	newBlock.MineBlock(bc.Difficulty)

	bc.Chain = append(bc.Chain, newBlock)

	bc.PendingTransactions = []*transaction.Transaction{}
	bc.AddTransaction(transaction.NewTransaction("", miningRewardAddress, bc.MiningReward))
}

func (bc *Blockchain) IsChainValid() bool {
	for i := 1; i < len(bc.Chain); i++ {
		currentBlock := bc.Chain[i]
		previousBlock := bc.Chain[i-1]

		if currentBlock.Hash != currentBlock.CalculateHash() {
			return false
		}

		if currentBlock.PrevHash != previousBlock.Hash {
			return false
		}
	}
	return true
}

func (bc *Blockchain) GetBalance(address string) float64 {
	balance := 0.0
	for _, block := range bc.Chain {
		for _, tx := range block.Transactions {
			if tx.Sender == address {
				balance -= tx.Amount
			}
			if tx.Recipient == address {
				balance += tx.Amount
			}
		}
	}
	return balance
}
