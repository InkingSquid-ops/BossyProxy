package Backend

import (
	"io"
	"net/http"
)

func BossyProxy(w http.ResponseWriter, url string) {
	resp, err := http.Get(url)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	defer resp.Body.Close()

	io.Copy(w, resp.Body)
}
