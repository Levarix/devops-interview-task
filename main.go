package main

import (
	"fmt"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// Get the value of the "secret" environment variable
	secret := os.Getenv("secret")

	if secret == "sUp3rS3creTVa1u3#" {
		// If the secret is set correctly, serve a green page
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("<html><body style=\"background-color: green;\"><h1>Success!</h1></body></html>"))
	} else {
		// If the secret is not set correctly, serve a red page
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("<html><body style=\"background-color: red;\"><h1>Something doesn't work</h1></body></html>"))
	}
}

func main() {
	http.HandleFunc("/", handler)
	port := "8080"
	fmt.Printf("Listening on port %s...\n", port)
	http.ListenAndServe("127.0.0.1:"+port, nil)
}