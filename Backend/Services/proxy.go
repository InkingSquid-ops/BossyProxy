package Services

import (
	"net/http"

	HTML "github.com/InkingSquid-ops/BossyProxy/Backend/Parser/Html"
)

func Fetch(w http.ResponseWriter, TargetUrl string) {
	if TargetUrl == "" {
		http.Error(w, "missing url parameter", http.StatusBadRequest)
		return
	}

	resp, err := http.Get(TargetUrl)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}

	defer resp.Body.Close()

	result, err := HTML.RewriteHTML(resp.Body, TargetUrl)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(result)
}

