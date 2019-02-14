package good

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"

	"github.com/pallat/good/middleware"
)

const (
	TextPlain       = "text/plain"
	ApplicationJSON = "application/json"
)

type Goods struct {
	rules []*Rule
	*http.Server
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

func (g *Goods) On(port int) {
	addr := fmt.Sprintf(":%s", strconv.Itoa(port))
	log.Println("serve on " + addr)

	g.Server = &http.Server{
		Addr:           addr,
		Handler:        g.handler(),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	func() {
		log.Fatal(g.Server.ListenAndServe())
	}()

	g.GracefulShutdown()
}

func (g *Goods) GracefulShutdown() {
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := g.Server.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
}
