package Handlers

import (
	"net/http"

	Security "github.com/InkingSquid-ops/BossyProxy/Backend/Security"
	Services "github.com/InkingSquid-ops/BossyProxy/Backend/Services"
)

func BrowseHandler(w http.ResponseWriter, r *http.Request) {
	target := r.URL.Query().Get("url")

	validTarget := Security.ValidateURL(target)

	Services.Fetch(w, validTarget)
}
