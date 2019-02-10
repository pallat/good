package good

import (
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
