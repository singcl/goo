package main

import (
	"net/http"
	"fmt"
	"log"
)

func HelloServer(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Inside HelloServer handler")
	fmt.Fprintf(w, "Hello, " + req.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", HelloServer)
	err := http.ListenAndServe(":50000", nil)
	// 或者
	//http.ListenAndServe(":50000", http.HandlerFunc(HelloServer)) //  http.HandlerFunc(HelloServer)强制类型转换
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}