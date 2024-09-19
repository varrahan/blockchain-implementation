package blockchain

import (
	"time"
)

type Block struct {
	Timestamp     int64
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
	Nonce         int
	Difficulty    int
}

// Creates a new block to add to blockchain
func NewBlock(data string, prevHash []byte, difficulty int) *Block {
	// Create a new block
	newBlock := &Block{time.Now().Unix(),
		[]byte(data),
		prevHash,
		[]byte{},
		0,
		difficulty}
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
	return NewBlock("Genesis Block, initial block in blockchain", []byte{}, InitialDifficulty)
}
