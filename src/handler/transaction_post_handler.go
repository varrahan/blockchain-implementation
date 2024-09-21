package handler

import (
	"encoding/json"
	"net/http"

	"blockchain-emulator/src/blockchain"
)

func (ctx *HandlerContext) AddTransaction_Post_Handler(w http.ResponseWriter, r *http.Request) {
	var transactionData struct {
		From   string
		To     string
		Amount float64
	}
	// Decode the request body into the transactionData struct
	if err := json.NewDecoder(r.Body).Decode(&transactionData); err != nil {
		http.Error(w, "Invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}
	// Validate the input
	if transactionData.From == "" || transactionData.To == "" || transactionData.Amount <= 0 {
		http.Error(w, "Invalid transaction data: 'from', 'to', and 'amount' are required, and amount must be positive", http.StatusBadRequest)
		return
	}
	// Create a new transaction
	tx := blockchain.NewTransaction(transactionData.From, transactionData.To, transactionData.Amount)
	ctx.Blockchain.AddTransaction(tx)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(tx)
}
