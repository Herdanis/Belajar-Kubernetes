package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	url := os.Getenv("CHECKSITE_URL")
	if url == "" {
		panic("Please set the CHECKSITE_URL environment variable to the URL you want to check.")
	}

	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("Failed to fetch URL %s: %s\n", url, err)
		os.Exit(1)
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("Failed to read response body: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("HTTP Status Code: %d\n", response.StatusCode)
	fmt.Printf("Response Body:\n%s\n", string(body))
}
