package main

import (
	"github.com/pallat/good"
)

// Handler
func hello(c good.Context) {
	c.OK("Hello, World!")
}

func main() {
	// good instance
	g := good.New()

	r := g.Rule()
	r.ContentType(good.TextPlain)
	r.GET("/", hello)

	// Start server
	g.Go(8888)
}
