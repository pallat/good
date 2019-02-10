package good

import (
	"io"
	"net/http"
)

type TextPlainResponse struct {
	w http.ResponseWriter
	Responser
}

func (r *TextPlainResponse) SetWriter(w http.ResponseWriter) {
	r.w = w
}

func (r *TextPlainResponse) commit(code int, s string) {
	r.w.Header().Set("Content-Type", TextPlain)
	r.w.WriteHeader(code)
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

func (r *TextPlainResponse) Created(v interface{}) {
	r.out(http.StatusCreated, v)
}

func (r *TextPlainResponse) NoContent(v interface{}) {
	r.out(http.StatusNoContent, v)
}

func (r *TextPlainResponse) NotModified(v interface{}) {
	r.out(http.StatusNotModified, v)
}

func (r *TextPlainResponse) BadRequest(v interface{}) {
	r.out(http.StatusBadRequest, v)
}

func (r *TextPlainResponse) Unauthorized(v interface{}) {
	r.out(http.StatusUnauthorized, v)
}

func (r *TextPlainResponse) Forbidden(v interface{}) {
	r.out(http.StatusForbidden, v)
}

func (r *TextPlainResponse) NotFound(v interface{}) {
	r.out(http.StatusNotFound, v)
}

func (r *TextPlainResponse) MethodNotAllowed(v interface{}) {
	r.out(http.StatusMethodNotAllowed, v)
}

func (r *TextPlainResponse) InternalServerError(v interface{}) {
	r.out(http.StatusInternalServerError, v)
}
