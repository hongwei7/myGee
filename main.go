package main

import (
	"gee"
	"net/http"
)

func helloHandler(c *gee.Context) {
	c.String(http.StatusOK, "hello world")
}

func handler1(c *gee.Context) {
	m := map[string]interface{}{
		"name": "hongwei7",
		"age":  25,
	}
	c.JSON(http.StatusOK, m)
}

func main() {
	gee := gee.New()
	gee.GET("/hello", helloHandler)
	gee.GET("/", handler1)
	gee.Run()
}
