package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

const templateDir = "templates"

var (
	port = envOrDefault("PORT", "8080")
)

type FooterPage struct {
	Year int
}

func main() {
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/blog", handleBlog)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	log.Fatalln(http.ListenAndServe(":"+port, nil))
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	file := "index.html"
	t, err := template.ParseFiles(filepath.Join(templateDir, file))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Error parsing template %s: %v", file, err)
		return
	}

	t.Execute(w, nil)

	if err := writeFooter(w, r); err != nil {
		log.Printf("Error writing footer: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
}

func handleBlog(w http.ResponseWriter, r *http.Request) {
	file := "blog.html"
	t, err := template.ParseFiles(filepath.Join(templateDir, file))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Error parsing template %s: %v", file, err)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	t.Execute(w, nil)
}

func writeFooter(w http.ResponseWriter, r *http.Request) error {
	file := "footer.html"
	t, err := template.ParseFiles(filepath.Join(templateDir, file))
	if err != nil {
		return err
	}

	t.Execute(w, FooterPage{
		Year: time.Now().Year(),
	})

	return nil
}

func envOrDefault(key, defaultValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defaultValue
}
