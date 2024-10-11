package routes

import (
	"html/template"
	"log"
	"net/http"
)

type IndexPage struct {
	Title string
}

func IndexHandler(template *template.Template) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := template.ExecuteTemplate(w, "index.html", IndexPage{Title: "Home"}); err != nil {
			log.Println("failed to execute template: ", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
}
