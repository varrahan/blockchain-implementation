package main

import (
	"fmt"
	"net/http"

	blockchain "blockchain-emulator/src/blockchain"
	router "blockchain-emulator/src/router"
	env_utils "blockchain-emulator/src/utils"
)

var bc *blockchain.Blockchain

func main() {
	// Load environment variable for port
	port := env_utils.EnvUtils()["PORT"]
	// Set up server and listen on port specified in dotenv file
	mux := router.CreateRouter()
	http.Handle("/", mux)
	fmt.Printf("Server has started on port: %s", port)
	http.ListenAndServe(":"+port, mux)
}
