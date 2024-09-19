package blockchain

type Blockchain struct {
	Blocks []*Block
}

// Intitialize blockchain with genesis block
func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{GenesisBlock()}}
}

// Add new block to blockchain
func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	difficulty := bc.AdjustDifficulty()
	newBlock := NewBlock(data, prevBlock.Hash, difficulty)
	bc.Blocks = append(bc.Blocks, newBlock)
}

// Return all blocks in the blockchain
func (bc *Blockchain) GetBlocks() []*Block {
	return bc.Blocks
}

