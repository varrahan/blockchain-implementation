package handler

import (
	"time"

	"blockchain-emulator/src/blockchain"
	"blockchain-emulator/src/utils"
)

type HandlerContext struct {
	Blockchain *blockchain.Blockchain
}

type MiningResult struct {
	Success bool
	Message string
	Block   *blockchain.Block
}

var miningTimeOffset time.Duration = time.Duration(utils.StringToInt(utils.EnvUtils["TIMER_OFFSET"])) * time.Second
var miningTime time.Duration = time.Duration(utils.StringToInt(utils.EnvUtils["TARGET_TIME"])) * time.Second
