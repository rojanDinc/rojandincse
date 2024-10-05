package routes

import (
	"net/http"
	"path/filepath"
)

func IndexHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, filepath.Join("templates", "index.html"))
	})
}
