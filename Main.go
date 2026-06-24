package main

import (
	"log"
	"net/http"

	handlers "github.com/InkingSquid-ops/BossyProxy/Backend/Handlers"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handlers.HomeHandler)
	mux.HandleFunc("/browse", handlers.BrowseHandler)

	log.Fatal(http.ListenAndServe(":8080", mux))
}
