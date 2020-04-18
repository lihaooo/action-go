package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	e := gin.Default()
	e.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	err := e.Run(":50000")
	if err != nil {
		log.Fatal(err)
	}
}
