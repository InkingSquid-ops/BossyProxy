package Handlers

import (
	"net/http"

	Services "github.com/InkingSquid-ops/BossyProxy/Backend/Services"
)

func BrowseHandler(w http.ResponseWriter, r *http.Request) {
	target := r.URL.Query().Get("url")
	Services.BossyProxy(w, target)
}
