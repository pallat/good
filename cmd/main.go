package main

import (
	"github.com/pallat/good"
)

// Handler
func hello(c good.Context) {
	c.OK(map[string]string{
		"message": "Hello, World!",
	})
}

func main() {
	// good instance
	g := good.New()

	r := g.Rule()
	r.ContentType(good.ApplicationJSON)
	r.GET("/", hello)

	// Start server
	g.On(8888)
}
