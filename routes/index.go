package routes

import (
	"html/template"
	"log/slog"
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
			slog.Error("failed to execute template", "error", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
}
