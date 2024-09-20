package handler

import (
	"encoding/json"
	"net/http"
	"blockchain-emulator/src/blockchain"
)

func (ctx *HandlerContext) AddTransaction(w http.ResponseWriter, r *http.Request) {
    var tx blockchain.Transaction
    if err := json.NewDecoder(r.Body).Decode(&tx); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    ctx.Blockchain.AddTransaction(&tx)
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(tx)
}
