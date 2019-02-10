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

func (r *TextPlainResponse) OK(v interface{}) {
	if s, ok := v.(string); ok {
		r.w.WriteHeader(http.StatusOK)
		r.w.Header().Set("Content-Type", TextPlain)
		io.WriteString(r.w, s)
	}
}
