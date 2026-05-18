package routes

import (
	"html/template"
	"log/slog"
	"net/http"
	"rojandincse/middleware"
)

type PageMeta struct {
	Title       string
	Description string
	Keywords    string
	FrontMatter
}

type Routes struct {
	template *template.Template
	mux      *http.ServeMux
}

func NewRoutes(template *template.Template) *Routes {
	routes := &Routes{
		template: template,
		mux:      http.NewServeMux(),
	}

	routes.setup()

	return routes
}

func (r *Routes) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	middleware.Logger(r.mux).ServeHTTP(w, req)
}

func (r *Routes) setup() {
	r.mux.Handle("/{$}", IndexHandler(r.template))
	r.mux.Handle("/blog/{$}", BlogHandler(r.template))
	r.mux.Handle("/contact/{$}", ContactHandler(r.template))
	r.mux.Handle("/healthz", HealthzHandler())
	r.mux.Handle("/post/{post}/{$}", PostHandler(r.template))
	r.mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	r.mux.Handle("/", NotFoundHandler(r.template))
}

type NotFoundPage struct {
	PageMeta
}

func NotFoundHandler(template *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		page := NotFoundPage{
			PageMeta: PageMeta{
				Title:       "404",
				Description: "Page not found",
				Keywords:    "",
			},
		}
		if err := template.ExecuteTemplate(w, "404.html", page); err != nil {
			slog.Error("failed to execute 404 template", "error", err)
			http.Error(w, "", http.StatusInternalServerError)
			return
		}
	}
}
