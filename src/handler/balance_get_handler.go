package handler

import (
	"encoding/json"
	"net/http"

	mux "github.com/gorilla/mux"
)

func (ctx *HandlerContext) GetBalance_Get_Handler(w http.ResponseWriter, r *http.Request) {
	// Get address of user with query parameter
	vars := mux.Vars(r)
	address := vars["address"]
	// Get the balance as a float of each transaction where address existed
	balance := ctx.Blockchain.GetBalance(address)
	json.NewEncoder(w).Encode(map[string]float64{"balance": balance})
}
