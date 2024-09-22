package main

import (
	"fmt"
	"net/http"

	router "blockchain-emulator/src/router"
	env_utils "blockchain-emulator/src/utils"
)

func main() {
	// Load environment variable for port
	// Container port represents the port inside the docker container that the app is listening on
	// Host port represents the port that is mapped to the contaner port, used to prevent issues of applications using the same port on host machine
	local_port := env_utils.EnvUtils()["CONTAINER_PORT"]
	docker_port := env_utils.EnvUtils()["HOST_PORT"]
	// Set up server and listen on port specified in dotenv file
	mux := router.CreateRouter()
	http.Handle("/", mux)
	fmt.Printf("Server has started on port %s for local build\n", local_port)
	fmt.Printf("Server is accessible on port %s for docker build", docker_port)
	http.ListenAndServe(":"+local_port, mux)
}
