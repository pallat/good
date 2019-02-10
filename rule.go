package good

import "net/http"

type Route struct {
	method string
	path   string
	h      HandlerFunc
}

type Rule struct {
	response Responser
	routes   []*Route
}

func (r *Rule) handler() []*Handler {
	h := []*Handler{}
	for _, route := range r.routes {
		h = append(h, &Handler{
			Responser: r.response,
			path:      route.path,
			method:    route.method,
			h:         route.h,
		})
	}
	return h
}

func (r *Rule) ContentType(ct string) {
	if ct == TextPlain {
		r.response = &TextPlainResponse{}
	}
	if ct == ApplicationJSON {
		r.response = &AppJSONResponse{}
	}
}

func (r *Rule) append(method, path string, h HandlerFunc) {
	r.routes = append(r.routes, &Route{
		method: method,
		path:   path,
		h:      h,
	})
}

func (r *Rule) Add(method, path string, h HandlerFunc) {
	r.append(method, path, h)
}

func (r *Rule) GET(path string, h HandlerFunc) {
	r.Add(http.MethodGet, path, h)
}
