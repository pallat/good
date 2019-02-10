package good

import "net/http"

type HandlerFunc func(Context)

type Handler struct {
	Responser
	path   string
	method string
	h      HandlerFunc
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.h(NewContext(w, r, h.Responser))
}
