package main

import (
	"encoding/json"
	"fmt"
)

//json 包使用 map[string]interface{} 和 []interface{} 储存任意的 JSON 对象和数组；
// 其可以被反序列化为任何的 JSON blob 存储到接口值中。
func main() {
	b := []byte(`{"Name": "Wednesday", "Age": 6, "Parents": ["Gomez", "Morticia"]}`)
	var f interface{}
	json.Unmarshal(b, &f)

	//f 指向的值是一个 map，key 是一个字符串，value 是自身存储作为空接口类型的值：
	/*
	map[string]interface{} {
		"Name": "Wednesday",
		"Age":  6,
		"Parents": []interface{} {
			"Gomez",
			"Morticia",
		},
	}
	*/
	/*要访问这个数据，我们可以使用类型断言*/
	m := f.(map[string]interface{})
	//我们可以通过 for range 语法和 type switch 来访问其实际类型：
	for k, v := range m {
		switch vv := v.(type) {
		case string:
			fmt.Println(k, "is string", vv)
		case int:
			fmt.Println(k, "is int", vv)
		case []interface{}:
			fmt.Println(k, "is an array")
			for i, u := range vv {
				fmt.Println(i, u)
			}
		default:
			fmt.Println(k , "is of a type I don’t know how to handle")
		}
	}
	//通过这种方式，你可以处理未知的 JSON 数据，同时可以确保类型安全。
}
