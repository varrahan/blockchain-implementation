package handler

import (
	blockchain "blockchain-emulator/src/blockchain"
)

type HandlerContext struct {
	Blockchain *blockchain.Blockchain
}
