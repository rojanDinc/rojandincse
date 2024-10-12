package routes

import (
	"html/template"
	"log"
	"net/http"
)

type ContactPage struct {
	PageMeta
}

func ContactHandler(template *template.Template) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cp := ContactPage{
			PageMeta: PageMeta{
				Title:       "Contact",
				Keywords:    "contact, contacts",
				Description: "Contact page",
			},
		}

		if err := template.ExecuteTemplate(w, "contact.html", cp); err != nil {
			log.Println("failed to execute template: ", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
}
