package handler

import (
	"encoding/json"
	"log"
	"net/http"
)

func (ctx *HandlerContext) Health_Check(w http.ResponseWriter, r *http.Request) {
	log.Printf("Health check request received")
	health_result := [3]bool{false, false, false}
	// Health checks to determine validity and health of blockchain
	if ctx.Blockchain.Blocks[0].IsValid() {
		health_result[0] = true
	}
	if ctx.Blockchain.GetBlocks() != nil {
		health_result[1] = true
	}
	if ctx.Blockchain.Blocks[0].Transactions != nil {
		health_result[2] = true
	}

	for _, item := range health_result {
		if !item {
			http.Error(w, "Health check has failed", 500)
			return
		}
	}

	json.NewEncoder(w).Encode(map[string][]bool{"data": health_result[:]})
}
