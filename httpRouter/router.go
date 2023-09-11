package httpRouter

import "net/http"

type Router struct{}

func CreateNewRouter() *Router {
	router := &Router{}

	return router
}

func (r *Router) Run() {
	http.ListenAndServe(":8080", nil)
}

func (r *Router) GET(path string, handler http.HandlerFunc) {
	http.HandleFunc(path, handler)
}

func (r *Router) POST(path string, handler http.HandlerFunc) {
	http.HandleFunc(path, handler)
}

func (r *Router) PUT(path string, handler http.HandlerFunc) {
	http.HandleFunc(path, handler)
}

func (r *Router) DELETE(path string, handler http.HandlerFunc) {
	http.HandleFunc(path, handler)
}
