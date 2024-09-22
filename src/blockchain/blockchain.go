package blockchain

import (
	"sync"

	utils "blockchain-emulator/src/utils"
)

type Blockchain struct {
	Blocks              []*Block
	PendingTransactions *TransactionPool
	Reward              float64
	Mutex               sync.Mutex
}

// Intitialize blockchain with genesis block
func NewBlockchain() *Blockchain {
	return &Blockchain{
		Blocks:              []*Block{GenesisBlock()},
		PendingTransactions: NewTransactionPool(),
		Reward:              utils.StringToFloat(utils.EnvUtils()["REWARD"]),
	}
}

func (bc *Blockchain) AddTransaction(transaction *Transaction) {
	if transaction.IsValid() {
		bc.PendingTransactions.AddTransaction(transaction)
	}
}

// Mine new block
func (bc *Blockchain) MineBlock(minerAddress string) {
	bc.Mutex.Lock()
	defer bc.Mutex.Unlock()
	// Initialize data added to block
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	difficulty := bc.AdjustDifficulty()
	index := prevBlock.Index + 1
	// Create transaction
	transaction := NewTransaction("Blockchain", minerAddress, bc.Reward)
	bc.PendingTransactions.AddTransaction(transaction)
	// Create block and add to blockchain if valid
	newBlock := NewBlock(index, bc.PendingTransactions.Transactions, prevBlock.Hash, difficulty)
	if newBlock.IsValid() {
		bc.Blocks = append(bc.Blocks, newBlock)
		bc.PendingTransactions.Clear()
	}
}

func (bc *Blockchain) GetBalance(address string) float64 {
	balance := 0.0
	for _, block := range bc.Blocks {
		for _, tx := range block.Transactions {
			if tx.From == address {
				balance -= tx.Amount
			}
			if tx.To == address {
				balance += tx.Amount
			}
		}
	}
	return balance
}

// Return all blocks in the blockchain
func (bc *Blockchain) GetBlocks() []*Block {
	return bc.Blocks
}
