package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Println("Request started")
	defer log.Println("Request finished")

	select {
	case <-time.After(5 * time.Second):
		// Show on command line stdout
		log.Println("Request processed with success.\n")
		// Show on browser stdout
		w.Write([]byte("Request processed with success.\n"))
	case <-ctx.Done():
		// Show on command line stdout
		log.Println("Request cancelled by the client.\n")
		// Show on browser stdout
		http.Error(w, "Request cancelled by the client.\n", http.StatusRequestTimeout)
	}
}
