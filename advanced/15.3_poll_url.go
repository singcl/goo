package main

import (
	"net/http"
	"fmt"
)

var urls = []string{
	"http://baidu.com",
	"http://taobao.com",
	"http://juejin.com",
}

func main() {
	// Execute an HTTP HEAD request for all url's
	// and returns the HTTP status string or an error string.
	for _, url := range urls {
		resp, err := http.Head(url)
		if err != nil {
			fmt.Println("Error:", url, err)
		}
		fmt.Println(url, ":", resp.Status)
	}
}