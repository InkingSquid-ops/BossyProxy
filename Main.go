package main

import (
	"log"
	"net/http"

	handlers "github.com/InkingSquid-ops/BossyProxy/Backend/Handlers"
)

func main() {
	mux := http.NewServeMux()

	fs := http.FileServer(http.Dir("Web/Templates/Static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	mux.HandleFunc("/", handlers.HomeHandler)
	mux.HandleFunc("/browse", handlers.BrowseHandler)

	log.Fatal(http.ListenAndServe(":8080", mux))
}
