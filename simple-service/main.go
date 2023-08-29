package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Response struct {
	StatusCode int    `json:"status_code"`
	SecretMessage    string `json:"SecretMessage"`
	Version    string `json:"version"`
	ServerName	string `json:"servername"`
}

func main() {
	port := os.Getenv("PORT")
	secret := os.Getenv("SECRET_MESSAGE")
	host, _ := os.Hostname()
	if port == "" {
		panic("No PORT environment variable detected.")
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		response := Response{
			StatusCode: 200,
			SecretMessage: secret,
			Version:    "4.0",
			ServerName: host,
		}

		w.Header().Set("Content-Type", "application/json")

		fmt.Printf("Sending response: %+v\n", response)

		json.NewEncoder(w).Encode(response)
	})

	fmt.Printf("Server starting on port %s\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		fmt.Printf("Failed to start server: %v\n", err)
	}
}
