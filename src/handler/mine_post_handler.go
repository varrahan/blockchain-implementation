package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func (ctx *HandlerContext) Mine_Block_Post_Handler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Mine block post request received")
	if len(ctx.Blockchain.PendingTransactions.Transactions) == 0 {
		err := "No transactions are pending to add to new block. Please ensure that transaction pool is not empty before mining a block onto the blockchain"
		log.Printf("%s", err)
		http.Error(w, err, http.StatusBadRequest)
		return
	}
	// Create channel and go function to process mining on seperate thread
	miningChannel := make(chan MiningResult)
	go func() {
		ctx.Blockchain.MineBlock()
		newBlock := ctx.Blockchain.Blocks[len(ctx.Blockchain.Blocks)-1]
		miningChannel <- MiningResult{
			Success: true,
			Message: "Block has been mined onto the blockchain",
			Block:   newBlock,
		}
	}()
	var result MiningResult
	select {
	// Case when mining is completed, send result response to client
	case miningResult := <-miningChannel:
		w.WriteHeader(http.StatusAccepted)
		result = miningResult
	// Case when mining takes too long, send time out response to client
	case <-time.After(miningTimeOffset + miningTime):
		err := "Mining timed out"
		log.Printf("%s", err)
		w.WriteHeader(http.StatusGatewayTimeout)
		result = MiningResult{
			Success: false,
			Message: err,
			Block:   nil,
		}
	// Case when client disconnects or request is cancelled, send client disconnect error response
	case <-r.Context().Done():
		err := "Client disconnect"
		log.Printf("%s", err)
		w.WriteHeader(499)
		result = MiningResult{
			Success: false,
			Message: err,
			Block:   nil,
		}
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
