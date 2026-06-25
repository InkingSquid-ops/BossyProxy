package Services

import (
	"net/http"
	"net/url"

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

	baseURL, err := url.Parse(TargetUrl)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := HTML.RewriteHTML(resp.Body, baseURL)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(result)
}

