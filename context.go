package good

import "net/http"

type Context interface {
	Responser
}

type context struct {
	w http.ResponseWriter
	r *http.Request
	Responser
}

func NewContext(w http.ResponseWriter, r *http.Request, response Responser) Context {
	response.SetWriter(w)
	return &context{w, r, response}
}
