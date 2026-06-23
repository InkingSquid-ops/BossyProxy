package main

import (
	"BossyProxy/Backend"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", Backend.HomeHandler)
	http.HandleFunc("/browse", Backend.BrowseHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
