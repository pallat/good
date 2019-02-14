package main

import (
	"github.com/pallat/good"
)

// Handler
func hello(c good.Context) {
	c.OK("Hello, World!")
}

type request struct {
	Name string
	Age  int
}

func bind(c good.Context) {
	var req request

	err := c.Bind(&req)
	if err != nil {
		c.InternalServerError(err)
		return
	}

	c.OK(req)
}

func main() {
	// good instance
	g := good.New()
	// g.GracefulShutdown()

	r := g.Rule()
	r.ContentType(good.ApplicationJSON)
	r.POST("/a", bind)

	rs := g.Rule()
	rs.ContentType(good.TextPlain)
	rs.GET("/", hello)

	// Start server
	g.On(8888)
}
