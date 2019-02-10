package good

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Context interface {
	Responser
	Requester
}

type Requester interface {
	SetRequest(r *http.Request)
	Param(string) string
	Query(key string) string
	Bind(interface{}) error
}

type AppJSONRequest struct {
	r *http.Request
}

func (r *AppJSONRequest) SetRequest(req *http.Request) {
	r.r = req
}

func (r *AppJSONRequest) Param(string) string {
	return ""
}

func (r *AppJSONRequest) Query(key string) string {
	keys, ok := r.r.URL.Query()[key]

	if !ok || len(keys[0]) < 1 {
		return ""
	}

	return keys[0]
}

func (r *AppJSONRequest) Bind(v interface{}) error {
	b, err := ioutil.ReadAll(r.r.Body)
	if err != nil {
		return err
	}
	defer r.r.Body.Close()

	return json.Unmarshal(b, v)
}

type context struct {
	w http.ResponseWriter
	r *http.Request
	Responser
	Requester
}

func NewContext(w http.ResponseWriter, r *http.Request, response Responser, request Requester) Context {
	response.SetWriter(w)
	request.SetRequest(r)
	return &context{w: w, r: r, Responser: response, Requester: request}
}
