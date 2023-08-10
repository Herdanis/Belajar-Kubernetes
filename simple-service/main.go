package main

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	StatusCode int `json:"status_code"`
	Message string `json:"message"`
	Version string `json:"version"`
}

func main()  {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		response := Response{
			StatusCode: 200,
			Message:    "OK",
			Version: "2.0",
		}

		// Set the content type to JSON
		w.Header().Set("Content-Type", "application/json")

		// Marshal the struct into JSON and write to the response
		json.NewEncoder(w).Encode(response)
	})

	http.ListenAndServe(":8089",nil)
}