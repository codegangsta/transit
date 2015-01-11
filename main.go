package transit

import (
	"net/http"
	"net/url"

	"github.com/julienschmidt/httprouter"
)

type Router struct {
	router *httprouter.Router
}

func (ro *Router) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	ro.router.ServeHTTP(rw, r)
}

func (r *Router) Handle(method, path string, handler http.Handler) {
	r.router.Handle(method, path, func(rw http.ResponseWriter, r *http.Request, params httprouter.Params) {
		if len(params) > 0 {
			r.URL.RawQuery = toValues(params).Encode() + "&" + r.URL.RawQuery
		}
		handler.ServeHTTP(rw, r)
	})
}

func (r *Router) HandleFunc(method, path string, handler http.HandlerFunc) {
	r.Handle(method, path, handler)
}

func (r *Router) Get(path string, handler http.HandlerFunc) {
	r.HandleFunc("GET", path, handler)
}

func (r *Router) Post(path string, handler http.HandlerFunc) {
	r.HandleFunc("POST", path, handler)
}

func (r *Router) Put(path string, handler http.HandlerFunc) {
	r.HandleFunc("PUT", path, handler)
}

func (r *Router) Patch(path string, handler http.HandlerFunc) {
	r.HandleFunc("PATCH", path, handler)
}

func (r *Router) Delete(path string, handler http.HandlerFunc) {
	r.HandleFunc("DELETE", path, handler)
}

func (r *Router) Head(path string, handler http.HandlerFunc) {
	r.HandleFunc("Head", path, handler)
}

func New() *Router {
	return &Router{httprouter.New()}
}

func toValues(params httprouter.Params) url.Values {
	v := url.Values{}
	for _, p := range params {
		v.Set(p.Key, p.Value)
	}
	return v
}
