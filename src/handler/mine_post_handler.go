package handler

import (
	"encoding/json"
	"net/http"

	mux "github.com/gorilla/mux"
)

func (ctx *HandlerContext) MineBlock_Post_handler(w http.ResponseWriter, r *http.Request) {
	// Get address from query parameters
	vars := mux.Vars(r)
	minerAddress := vars["address"]
	// If address is blank raise error
	if minerAddress == "" {
		http.Error(w, "Miner address is required", http.StatusBadRequest)
		return
	}
	// Mine the block to add new block to blockchain
	ctx.Blockchain.MineBlock(minerAddress)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(ctx.Blockchain.Blocks[len(ctx.Blockchain.Blocks)-1])
}
