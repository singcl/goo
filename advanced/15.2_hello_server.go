package main

import (
	"net/http"
	"fmt"
)

type Hello struct {}

func (h *Hello) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello Singcl!")
}

func main() {
	var h *Hello
	http.ListenAndServe(":40000", h)
}