package good

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type AppJSONResponse struct {
	w http.ResponseWriter
}

func (r *AppJSONResponse) SetWriter(w http.ResponseWriter) {
	r.w = w
}

func (r *AppJSONResponse) commit(code int, v interface{}) {
	b, err := json.Marshal(v)
	if err != nil {
		s := fmt.Sprintf("application/json not expect %s", err)
		io.WriteString(r.w, s)
	}
	r.w.Header().Set("Content-Type", ApplicationJSON)
	r.w.WriteHeader(code)
	r.w.Write(b)
}

func (r *AppJSONResponse) OK(v interface{}) {
	r.commit(http.StatusOK, v)
}

func (r *AppJSONResponse) Created(v interface{}) {
	r.commit(http.StatusCreated, v)
}

func (r *AppJSONResponse) NoContent(v interface{}) {
	r.commit(http.StatusNoContent, v)
}

func (r *AppJSONResponse) NotModified(v interface{}) {
	r.commit(http.StatusNotModified, v)
}

func (r *AppJSONResponse) BadRequest(v interface{}) {
	r.commit(http.StatusBadRequest, v)
}

func (r *AppJSONResponse) Unauthorized(v interface{}) {
	r.commit(http.StatusUnauthorized, v)
}

func (r *AppJSONResponse) Forbidden(v interface{}) {
	r.commit(http.StatusForbidden, v)
}

func (r *AppJSONResponse) NotFound(v interface{}) {
	r.commit(http.StatusNotFound, v)
}

func (r *AppJSONResponse) MethodNotAllowed(v interface{}) {
	r.commit(http.StatusMethodNotAllowed, v)
}

func (r *AppJSONResponse) InternalServerError(v interface{}) {
	r.commit(http.StatusInternalServerError, v)
}
