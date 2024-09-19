package handler

import (
	"encoding/json"
	"net/http"
)

func (ctx *HandlerContext) PostHandler(w http.ResponseWriter, r *http.Request) {
	// Map to hold json
	var requestData map[string]interface{}
	// Decode JSON request body and store in map
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&requestData); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	// Extract the value for the key 'data'
	data, ok := requestData["data"].(string)
	if !ok {
		http.Error(w, "Missing or invalid 'data' field", http.StatusBadRequest)
		return
	}
	// Add block to blockchain with parsed data
	ctx.Blockchain.AddBlock(data)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Block added successfully"))
}
