package routes

import (
	"html/template"
	"log"
	"net/http"
)

type IndexPage struct {
	PageMeta
}

func IndexHandler(template *template.Template) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip := IndexPage{
			PageMeta: PageMeta{
				Title:       "Home",
				Keywords:    "home, homepage",
				Description: "Home page",
			},
		}
		if err := template.ExecuteTemplate(w, "index.html", ip); err != nil {
			log.Println("failed to execute template: ", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
}
