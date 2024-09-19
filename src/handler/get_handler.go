package handler

import (
	"encoding/json"
	"net/http"
)

func (ctx *HandlerContext) GetHandler (w http.ResponseWriter, r *http.Request) {
    blocks := ctx.Blockchain.GetBlocks()
	// build response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(blocks)
}