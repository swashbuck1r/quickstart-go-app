package main

import (
	"flag"
	"fmt"

	"github.com/gin-gonic/gin"
)

var (
	port int
)

func main() {
	flag.IntVar(&port, "port", 8080, "listen port")
	flag.Parse()

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello World!")
	})

	r.GET("/:message", func(c *gin.Context) {
		c.String(200, fmt.Sprintf("Hello %s!", c.Param("message")))
	})

	r.Run(fmt.Sprintf(":%d", port))
}
