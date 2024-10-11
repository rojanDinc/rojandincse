package routes

import (
	"html/template"
	"log"
	"net/http"
)

type ContactPage struct {
	Title string
}

func ContactHandler(template *template.Template) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := template.ExecuteTemplate(w, "contact.html", ContactPage{Title: "Contact"}); err != nil {
			log.Println("failed to execute template: ", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
}
