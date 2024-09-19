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
	mux.HandleFunc("/addblock", ctx.PostHandler).Methods("POST")
	mux.HandleFunc("/getblocks", ctx.GetHandler).Methods("GET")
	return mux
}
