package blockchain

import (
	"time"
)

type Block struct {
	Index         int64
	Timestamp     time.Time
	Transactions  []*Transaction
	PrevBlockHash []byte
	Hash          []byte
	Nonce         int
	Difficulty    int
}

// Creates a new block to add to blockchain
func NewBlock(index int64, transactions []*Transaction, prevHash []byte, difficulty int) *Block {
	// Create a new block
	newBlock := &Block{
		Index:         index,
		Timestamp:     time.Now(),
		Transactions:  transactions,
		PrevBlockHash: prevHash,
		Difficulty:    difficulty}
	// Create proof of work for new block
	pow := NewProofOfWork(newBlock)
	// Run proof of work to find and set the nonce and hash
	nonce, hash := pow.Run()
	newBlock.Nonce = nonce
	newBlock.Hash = hash
	return newBlock
}

// Creates a genesis block to begin blockchain; does not reference a previous block
func GenesisBlock() *Block {
	return NewBlock(0, []*Transaction{}, []byte{}, InitialDifficulty)
}

func (b *Block) IsValid() bool {
	return b.Hash != nil || b.Transactions != nil
}
