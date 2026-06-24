package Handlers

import (
	"net/http"
	"path/filepath"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, filepath.Join("Web", "Templates", "Index.html"))
}
