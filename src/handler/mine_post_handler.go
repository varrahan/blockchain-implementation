package handler

import (
	"net/http"
	"time"
)

func (ctx *HandlerContext) MineBlock_Post_handler(w http.ResponseWriter, r *http.Request) {
	// Send response to indicate mining has started
	sendSSE(w, map[string]string{
		"status": "Mining started",
	})
	// Create channel and go function to process mining on seperate thread
	result := make(chan MiningResult)
	go func() {
		ctx.Blockchain.MineBlock()
		newBlock := ctx.Blockchain.Blocks[len(ctx.Blockchain.Blocks)-1]
		result <- MiningResult{
			Success: true,
			Block:   newBlock,
		}
	}()
	select {
	// Case when mining is completed, send result response to client
	case miningResult := <-result:
		sendSSE(w, miningResult)
	// Case when mining takes too long, send time out response to client
	case <-time.After(miningTimeOffset + miningTime):
		sendSSE(w, map[string]string{"status": "Mining timed out"})
	// Case when client disconnects or request is cancelled, stop function
	case <-r.Context().Done():
		return
	}
}
