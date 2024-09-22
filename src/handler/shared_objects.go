package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	blockchain "blockchain-emulator/src/blockchain"
	utils "blockchain-emulator/src/utils"
)

type HandlerContext struct {
	Blockchain *blockchain.Blockchain
}

type MiningResult struct {
	Success bool
	Status  string
	Block   *blockchain.Block
}

var miningTimeOffset time.Duration = time.Duration(utils.StringToInt(utils.EnvUtils()["TIMER_OFFSET"])) * time.Second
var miningTime time.Duration = time.Duration(utils.StringToInt(utils.EnvUtils()["TARGET_TIME"])) * time.Second

func sendSSE(w http.ResponseWriter, data interface{}) {
	data, err := json.Marshal(data)
	// Check for error
	if err != nil {
		return
	}
	// Write response
	fmt.Fprintf(w, "data: %s\n\n", data)
	// If assertion is successful, flush buffered data to client
	if f, ok := w.(http.Flusher); ok {
		f.Flush()
	}
}
