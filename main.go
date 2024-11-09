package main

import (
	"crypto/tls"
	"log"
	"net/http"

	router "blockchain-emulator/src/router"
	utils "blockchain-emulator/src/utils"
)

func main() {
	// Load port data for local and docker build
	local_port := utils.EnvUtils["CONTAINER_PORT"]
	docker_port := utils.EnvUtils["HOST_PORT"]
	// Certificates paths for certificate and key to enable TLS
	security_certificate := utils.EnvUtils["SECURITY_CERT_PATH"]
	security_key := utils.EnvUtils["SECURITY_KEY_PATH"]

	// HTTPS TLS config
	tlsConfig := &tls.Config{
		MinVersion:               tls.VersionTLS13,
		PreferServerCipherSuites: true,
		CipherSuites: []uint16{
			tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305,
			tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,
			tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
			tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
		},
		InsecureSkipVerify: false,
	}

	// Set up server and listen on port specified in dotenv file
	mux := router.CreateRouter()
	http.Handle("/", mux)

	httpsServer := http.Server{
		Handler:   mux,
		Addr:      ":" + local_port,
		TLSConfig: tlsConfig,
	}
	log.Printf("Server has started on port %s for local build\n", local_port)
	log.Printf("Server is accessible on port %s for docker build\n", docker_port)
	if err := httpsServer.ListenAndServeTLS(security_certificate, security_key); err != http.ErrServerClosed {
		log.Printf("HTTPS server error: %v", err)
	}
}
