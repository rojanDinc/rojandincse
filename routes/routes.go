package routes

import (
	"html/template"
	"net/http"
)

type Routes struct {
	template *template.Template
	mux      *http.ServeMux
}

func NewRoutes() *Routes {
	routes := &Routes{
		template: template.Must(template.ParseGlob("templates/*.html")),
		mux:      http.NewServeMux(),
	}

	routes.setup()

	return routes
}

func (r *Routes) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.mux.ServeHTTP(w, req)
}

func (r *Routes) setup() {
	r.mux.Handle("/{$}", IndexHandler(r.template))
	r.mux.Handle("/blog/{$}", BlogHandler(r.template))
	r.mux.Handle("/contact/{$}", ContactHandler(r.template))
	r.mux.Handle("/healthz", HealthzHandler())
	r.mux.Handle("/post/{post}/{$}", PostHandler(r.template))
	r.mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
}
