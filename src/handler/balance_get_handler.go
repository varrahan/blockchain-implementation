package handler

import (
	"encoding/json"
	"log"
	"net/http"

	mux "github.com/gorilla/mux"
)

func (ctx *HandlerContext) Get_Balance_Get_Handler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Get balance request received")
	// Get address of user with query parameter
	vars := mux.Vars(r)
	address, ok := vars["address"]
	// Check if address exists in request and is not an empty string
	if !ok {
		err := "Address is not in request. Please ensure that address is used in request"
		log.Printf("%s", err)
		http.Error(w, err, http.StatusBadRequest)
		return
	} else if address == "" {
		err := "Address is empty. Please ensure that address has an value associated with it"
		log.Printf("%s", err)
		http.Error(w, err, http.StatusBadRequest)
		return
	}
	// Get the balance as a float of each transaction where address existed
	balance := ctx.Blockchain.GetBalance(address)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]any{"address": address, "balance": balance})
}
