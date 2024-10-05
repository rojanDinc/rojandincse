package routes

import "net/http"

type Routes struct {
	mux *http.ServeMux
}

func NewRoutes() *Routes {
	return &Routes{
		mux: http.NewServeMux(),
	}
}

func (r *Routes) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.setup()
	r.mux.ServeHTTP(w, req)
}

func (r *Routes) setup() {
	r.mux.Handle("/", IndexHandler())
}
