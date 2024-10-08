package routes

import (
	"net/http"
)

type Routes struct {
	mux *http.ServeMux
}

func NewRoutes() *Routes {
	routes := &Routes{
		mux: http.NewServeMux(),
	}

	routes.setup()

	return routes
}

func (r *Routes) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.mux.ServeHTTP(w, req)
}

func (r *Routes) setup() {
	r.mux.Handle("/{$}", IndexHandler())
	r.mux.Handle("/blog/{$}", BlogHandler())
	r.mux.Handle("/post/{post}/{$}", PostHandler())
}
