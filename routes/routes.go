package routes

import (
	"html/template"
	"net/http"
	"rojandincse/middleware"
)

type PageMeta struct {
	Title       string
	Description string
	Keywords    string
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
}
