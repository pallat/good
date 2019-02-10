/*
Example:

  package main

  import (
	"github.com/chonla/cotton/response"
    "net/http"

    "github.com/pallat/good/v1"
    "github.com/pallat/good/v1/middleware"
  )

  // Handler
  func hello(c good.Context) error {
    return c.OK("Hello, World!")
  }

  func main() {
    // good instance
    g := good.New()

    // Middleware
    g.Use(middleware.Logger())
    g.Use(middleware.Recover())
    g.Use(middleware.GracefulShoutdown())

    r := g.Rule()
    r.ContentType(g.TextPlain)
    r.GET("/", hello)

    // Start server
    g.Logger.Fatal(g.Start(":8888"))
  }

Learn more at https://good.odds.io
*/

package good

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/pallat/good/middleware"
)

const (
	TextPlain = "text/plain"
)

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

type Goods struct {
	rules []*Rule
}

func New() *Goods {
	return &Goods{}
}

func (g *Goods) Use(middleware.MiddlewareFunc) {

}

func (g *Goods) newRule() *Rule {
	return &Rule{}
}

func (g *Goods) Rule() *Rule {
	r := g.newRule()
	g.rules = append(g.rules, r)
	return r
}

func (g *Goods) handler() http.Handler {
	mux := http.NewServeMux()
	for _, rule := range g.rules {
		hands := rule.handler()
		for _, h := range hands {
			mux.Handle(h.path, h)
		}
	}
	return mux
}

func (g *Goods) Go(port int) error {
	addr := fmt.Sprintf(":%s", strconv.Itoa(port))
	log.Println("serve on " + addr)
	return http.ListenAndServe(addr, g.handler())
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
	if ct == "text/plain" {
		r.response = &TextPlainResponse{}
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

type Route struct {
	method string
	path   string
	h      HandlerFunc
}
