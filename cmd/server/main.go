package main

import (
	"SmartHome/internal/utils"
	"fmt"
	"log"
	"net/http"
)

func main() {
	port := utils.GetPort()

	mux := http.NewServeMux()
	mux.HandleFunc("/health", healthHandler)

	server := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	log.Printf("Smart Home IoT server running on port %s\n", port)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, `{"status":"ok"}`)
}
