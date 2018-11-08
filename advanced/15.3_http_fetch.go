package main

import (
	"net/http"
	"fmt"
	"log"
	"io/ioutil"
)

var urls2 = []string{
	"http://baidu.com",
	"http://taobao.com",
	"http://juejin.com",
}

func main() {
	// Execute an HTTP HEAD request for all url's
	// and returns the HTTP status string or an error string.
	for _, url := range urls2 {
		resp, err := http.Get(url)
		checkError2(err)
		data, err := ioutil.ReadAll(resp.Body)
		checkError2(err)
		fmt.Printf("Got: %s", string(data))
	}
}

func checkError2(err error) {
	if err != nil {
		log.Fatalf("Get : %v", err)
	}
}
