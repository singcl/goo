package main

import (
	"net/http"
	"io"
	"fmt"
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

func main() {
	http.HandleFunc("/test1", SimpleServer)
	http.HandleFunc("/test2", FormServer)
	http.Handle("/", http.HandlerFunc(Home))

	if err := http.ListenAndServe(":50000", nil); err != nil {
		panic(err)
	}
}