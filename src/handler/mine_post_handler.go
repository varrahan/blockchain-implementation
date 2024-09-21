package handler

import (
	"encoding/json"
	"net/http"
)

func (ctx *HandlerContext) MineBlock_Post_handler(w http.ResponseWriter, r *http.Request) {
	// Get address from query parameters
	minerAddress := r.URL.Query().Get("address")
	// If address is blank raise error
	if minerAddress == "" {
		http.Error(w, "Miner address is required", http.StatusBadRequest)
		return
	}
	ctx.Blockchain.MineBlock(minerAddress)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(ctx.Blockchain.Blocks[len(ctx.Blockchain.Blocks)-1])
}
