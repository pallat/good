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
	TextPlain       = "text/plain"
	ApplicationJSON = "application/json"
)

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

func (g *Goods) On(port int) error {
	addr := fmt.Sprintf(":%s", strconv.Itoa(port))
	log.Println("serve on " + addr)
	return http.ListenAndServe(addr, g.handler())
}
