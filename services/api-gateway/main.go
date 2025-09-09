package main

import (
	"log"
	"net/http"

	"ride-sharing/shared/env"
)

var (
	httpAddr = env.GetString("GATEWAY_HTTP_ADDR", ":8081")
)

func main() {
	log.Println("Starting API Gateway")

	mux := http.NewServeMux()
	log.Println("Listening on " + httpAddr)
	log.Println("Calling handleTripPreview")
	mux.HandleFunc("POST /trip/preview", handleTripPreview)

	server := &http.Server{
		Addr:    httpAddr,
		Handler: mux,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Printf("Failed to start server: %v", err)
	}
}
