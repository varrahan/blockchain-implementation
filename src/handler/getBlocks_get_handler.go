package handler

import (
	"encoding/json"
	"net/http"
)

func (ctx *HandlerContext) GetBlocks_Get_Handler(w http.ResponseWriter, r *http.Request) {
	// Return all blocks in the blockchain
	json.NewEncoder(w).Encode(ctx.Blockchain.Blocks)
}
