package  main

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"fmt"
)

type Status struct {
	Text string
}

type User struct {
	Status Status
}


// json数据反序列化
// UnMarshal() 的函数签名是 func Unmarshal(data []byte, v interface{}) error 把 JSON 解码为数据结构。
func main() {
	/* 自己写的一个简单的Python web 服务器 返回一个json 数据 */
	res, _ := http.Get("http://127.0.0.1:8080/")
	/* initialize the structure of the JSON response */
	user := User{Status{""}}
	/* unmarshal the JSON into our structures */
	temp, _ := ioutil.ReadAll(res.Body)
	body := []byte(temp)
	json.Unmarshal(body, &user)
	fmt.Printf("status: %s", user.Status.Text)
}

/* Output:
status: This is the Go Json Text: Use Python SimpleHTTPServer build
*/