package good

import (
	"io"
	"net/http"
)

type Responser interface {
	SetWriter(w http.ResponseWriter)
	OK(interface{})
	Created(interface{})
	NoContent(interface{})
	NotModified(interface{})
	BadRequest(interface{})
	Unauthorized(interface{})
	Forbidden(interface{})
	NotFound(interface{})
	MethodNotAllowed(interface{})
	InternalServerError(interface{})
}

type TextPlainResponse struct {
	w http.ResponseWriter
	Responser
}

func (r *TextPlainResponse) SetWriter(w http.ResponseWriter) {
	r.w = w
}

func (r *TextPlainResponse) commit(code int, s string) {
	r.w.WriteHeader(code)
	r.w.Header().Set("Content-Type", TextPlain)
	io.WriteString(r.w, s)
}

func (r *TextPlainResponse) out(code int, v interface{}) {
	if s, ok := v.(string); ok {
		r.commit(code, s)
	} else {
		r.commit(http.StatusInternalServerError, "response not string")
	}
}

func (r *TextPlainResponse) OK(v interface{}) {
	r.out(http.StatusOK, v)
}
