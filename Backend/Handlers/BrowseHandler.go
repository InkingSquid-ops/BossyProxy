package Handlers

import (
	"net/http"
	"fmt"

	Security "github.com/InkingSquid-ops/BossyProxy/Backend/Security"
	Services "github.com/InkingSquid-ops/BossyProxy/Backend/Services"
)

func BrowseHandler(w http.ResponseWriter, r *http.Request) {
	target := r.URL.Query().Get("url")

	validTarget, err := Security.ValidateURL(target)
	
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid URL: %v", err.Error()), http.StatusBadRequest)
		return
	}

	Services.Fetch(w, validTarget)
}
