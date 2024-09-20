package handler

import (
	"encoding/json"
	"net/http"
)

func (ctx *HandlerContext) Balance_Get_Handler(w http.ResponseWriter, r *http.Request) {
    address := r.URL.Query().Get("address")
    balance := ctx.Blockchain.GetBalance(address)
    json.NewEncoder(w).Encode(map[string]float64{"balance": balance})
}