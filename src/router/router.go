package router

import (
	blockchain "blockchain-emulator/src/blockchain"
	"blockchain-emulator/src/handler"

	mux "github.com/gorilla/mux"
)

func CreateRouter() *mux.Router {
	// Create new blockchain
	bc := blockchain.NewBlockchain()

	// Create a new handler context with the blockchain
	ctx := &handler.HandlerContext{
		Blockchain: bc,
	}

	// Set up multiplexer router
	mux := mux.NewRouter()
	mux.HandleFunc("/mineblock", ctx.MineBlock_Post_handler).Methods("POST")
	mux.HandleFunc("/maketransaction", ctx.AddTransaction_Post_Handler).Methods("POST")
	mux.HandleFunc("/getbalance", ctx.GetBalance_Get_Handler).Methods("GET")
	mux.HandleFunc("/getblocks", ctx.GetBlocks_Get_Handler).Methods("GET")
	return mux
}
