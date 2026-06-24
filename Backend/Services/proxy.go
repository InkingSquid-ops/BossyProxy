package Services

import (
	"io"
	"net/http"
)

func Fetch(w http.ResponseWriter, url string) {
	resp, err := http.Get(url)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}

	defer resp.Body.Close()

	io.Copy(w, resp.Body)
}
