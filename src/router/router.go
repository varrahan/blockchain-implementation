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
	mux.HandleFunc("/mineblock", ctx.Mine_Block_Post_Handler).Methods("POST")
	mux.HandleFunc("/maketransaction", ctx.Add_Transaction_Post_Handler).Methods("POST")
	mux.HandleFunc("/getbalance/{address}", ctx.Get_Balance_Get_Handler).Methods("GET")
	mux.HandleFunc("/getblocks", ctx.Get_Blocks_Get_Handler).Methods("GET")
	mux.HandleFunc("/health", ctx.Health_Check).Methods("GET")
	return mux
}
