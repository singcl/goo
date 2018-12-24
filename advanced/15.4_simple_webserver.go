package main

import (
	"fmt"
	"io"
	"net/http"
)

const form = `
	<html><body>
		<form action="#" method="post" name="bar">
			<input type="text" name="in" />
			<input type="submit" value="submit"/>
		</form>
	</body></html>
`

/* handle a simple get request */
func SimpleServer(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("<h1>Hello World</h1>"))
}

func FormServer(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	switch req.Method {
	case "GET":
		/* display the form to the user */
		io.WriteString(w, form)
	case "POST":
		/* handle the form data, note that ParseForm must
		   be called before we can extract form data */
		//request.ParseForm();
		//io.WriteString(w, request.Form["in"][0])
		io.WriteString(w, req.FormValue("in"))
	}
}

func Home(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "<h1>Home %s </h1>", req.Method)
}

// type HandlerFunc func(ResponseWriter, *Request)
// 这里的类型别名和其他语言的类型别名不一样： HandlerFunc 和 func(ResponseWriter, *Request) 是两个不同的类型
// 要想兼容必须显示强制类型转换

func main() {
	http.HandleFunc("/test1", SimpleServer)
	http.HandleFunc("/test2", FormServer)
	// type Handler interface {
	// 	ServeHTTP(ResponseWriter, *Request)
	// }
	// type HandlerFunc func(ResponseWriter, *Request)
	// http.HandlerFunc 是定义在server.go中的一个类型
	// 这里http.HandlerFunc(Home)不是执行函数，而是强制类型转换,
	// http.HandlerFunc类型又实现了Handler接口，所有又是Handler接口类型
	http.Handle("/", http.HandlerFunc(Home))

	if err := http.ListenAndServe(":50000", nil); err != nil {
		panic(err)
	}
}
