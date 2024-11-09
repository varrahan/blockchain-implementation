package handler

import (
	"encoding/json"
	"log"
	"net/http"
)

func (ctx *HandlerContext) Get_Blocks_Get_Handler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Get blocks request received")
	// Check if blockchain exists and if blocks are not empty
	if ctx.Blockchain == nil {
		err := "Blockchain does not exist. Please restart program"
		log.Printf("%s", err)
		http.Error(w, err, http.StatusBadGateway)
		return
	} else if ctx.Blockchain.Blocks == nil {
		err := "Blockchain do not contain any blocks. Please restart program"
		log.Printf("%s", err)
		http.Error(w, err, http.StatusBadGateway)
		return
	}
	// Return all blocks in the blockchain
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ctx.Blockchain.Blocks)
}
