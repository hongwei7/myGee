package main

import (
	"fmt"
	"gee"
	"net/http"
)

func helloHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello %q", req.URL.Path)
}

func main() {
	gee := gee.New()
	gee.GET("/hello", helloHandler)
	gee.Run()
}
