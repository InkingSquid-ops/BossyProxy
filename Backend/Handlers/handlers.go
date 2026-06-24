package Handlers

import (
	"net/http"

	Services "github.com/InkingSquid-ops/BossyProxy/Backend/Services"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "Web/Templates/Index.html")
}

func BrowseHandler(w http.ResponseWriter, r *http.Request) {
	target := r.URL.Query().Get("url")
	Services.BossyProxy(w, target)
}
