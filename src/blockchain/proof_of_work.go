package blockchain

import (
	"bytes"
	"crypto/sha256"
	"math/big"
	"strconv"

	utils "blockchain-emulator/src/utils"
)

var InitialDifficulty int = utils.StringToInt(utils.EnvUtils()["INITIAL_DIFFICULTY"])   // Set initial difficulty
var TargetTimePerBlock int = utils.StringToInt(utils.EnvUtils()["TARGET_TIME"])         // Time in seconds
var AdjustmentInterval int = utils.StringToInt(utils.EnvUtils()["ADJUSTMENT_INTERVAL"]) // Number of blocks required to adjust; uses timestamp of block n-1 and block (n-1) - AdjustmentInterval block

type ProofOfWork struct {
	Block  *Block
	Target *big.Int
}

func (bc *Blockchain) AdjustDifficulty() int {
	blockCount := len(bc.Blocks)
	// If not enough blocks for adjustment, use the initial difficulty
	if blockCount < AdjustmentInterval {
		return InitialDifficulty
	}
	// Get the last block and the block from AdjustmentInterval blocks ago
	lastBlock := bc.Blocks[blockCount-1]
	currentDifficulty := lastBlock.Difficulty
	adjustmentBlock := bc.Blocks[blockCount-AdjustmentInterval]
	// Calculate the actual time taken to mine the last `AdjustmentInterval` blocks
	actualTimeTaken := lastBlock.Timestamp.Unix() - adjustmentBlock.Timestamp.Unix()
	expectedTime := AdjustmentInterval * TargetTimePerBlock
	// Calculate adjustment factor and new difficulty to use for block
	adjustmentFactor := float64(expectedTime) / float64(actualTimeTaken)
	newDifficulty := int(float64(currentDifficulty) * adjustmentFactor)
	if newDifficulty > currentDifficulty {
		// Blocks were mined too quickly, increase difficulty
		newDifficulty = currentDifficulty + 1
	} else if newDifficulty < currentDifficulty {
		// Blocks were mined too slowly, decrease difficulty
		newDifficulty = currentDifficulty - 1
	}
	// Ensure new difficulty is never less than minimum
	if newDifficulty < 1 {
		return InitialDifficulty
	}
	return newDifficulty
}

// Initialize new proof of work with target for new block
func NewProofOfWork(block *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-block.Difficulty))
	return &ProofOfWork{block, target}
}

// Method to concat fields within a block into a single slice of bytes that will be converted into the block hash
func (pow *ProofOfWork) PrepareData(nonce int) []byte {
	data := bytes.Join([][]byte{
		pow.Block.PrevBlockHash,
		Convert(pow.Block.Transactions),
		[]byte(strconv.FormatInt(pow.Block.Timestamp.Unix(), 10)),
		[]byte(strconv.Itoa(nonce)),
		[]byte(strconv.Itoa(pow.Block.Difficulty)),
	}, []byte{})
	return data
}

// Rehash data values until Nonce = Target. Once condition true return hash and nonce to set to new block
// This allows
func (pow *ProofOfWork) Run() (int, []byte) {
	var hash [32]byte
	var hashInt big.Int
	nonce := 0
	// loop until nonce is a value that creates a hash equal to target
	for {
		data := pow.PrepareData(nonce)
		hash = sha256.Sum256(data)
		hashInt.SetBytes(hash[:])
		if hashInt.Cmp(pow.Target) == -1 {
			break
		} else {
			nonce++
		}
	}
	return nonce, hash[:]
}
