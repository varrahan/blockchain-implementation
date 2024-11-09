package handler

import (
	"encoding/json"
	"log"
	"net/http"

	bc "blockchain-emulator/src/blockchain"
)

func (ctx *HandlerContext) Add_Transaction_Post_Handler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Add transaction post request received")
	var transactionData struct {
		From   string
		To     string
		Amount float64
	}
	// Decode the request body into the transactionData struct
	if err := json.NewDecoder(r.Body).Decode(&transactionData); err != nil {
		err := "Invalid request body: " + err.Error()
		log.Printf("%s", err)
		http.Error(w, err, http.StatusBadRequest)
		return
	}
	// Validate the input
	if transactionData.From == "" || transactionData.To == "" || transactionData.Amount <= 0 {
		err := "Invalid transaction data: 'from', 'to', and 'amount' are required, and amount must be positive"
		log.Printf("%s", err)
		http.Error(w, err, http.StatusBadRequest)
		return
	}
	// Create a new transaction
	transaction := bc.NewTransaction(transactionData.From, transactionData.To, transactionData.Amount)
	ctx.Blockchain.AddTransaction(transaction)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(transaction)
}
