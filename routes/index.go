package routes

import (
	"net/http"
	"path/filepath"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, filepath.Join("public", "index.html"))
}
