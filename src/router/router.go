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
	mux.HandleFunc("/mine", ctx.Mine_Post_handler).Methods("POST")
	mux.HandleFunc("/transaction", ctx.Balance_Get_Handler).Methods("POST")
	mux.HandleFunc("/balance", ctx.Balance_Get_Handler).Methods("GET")
	mux.HandleFunc("/getblocks", ctx.GetBlocks_Get_Handler).Methods("GET")
	return mux
}
