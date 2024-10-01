package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

var (
	port = envOrDefault("PORT", "8080")
)

func main() {
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/blog", handleBlog)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	log.Fatalln(http.ListenAndServe(":"+port, nil))
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	file := "index.html"
	t, err := template.ParseFiles(filepath.Join("public", file))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Error parsing template %s: %v", file, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	t.Execute(w, nil)
}

func handleBlog(w http.ResponseWriter, r *http.Request) {
	file := "blog.html"
	t, err := template.ParseFiles(filepath.Join("public", file))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Error parsing template %s: %v", file, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	t.Execute(w, nil)
}

func envOrDefault(key, defaultValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defaultValue
}
