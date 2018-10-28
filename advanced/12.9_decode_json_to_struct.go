package  main

import (
	"encoding/json"
	"fmt"
)

type FamilyMember struct {
	Name string
	Age int
	Parents [] string
}

//解码数据到结构

//事先知道 JSON 数据，我们可以定义一个适当的结构并对 JSON 数据反序列化。
func main () {
	b := []byte(`{"Name": "Wednesday", "Age": 6, "Parents": ["Gomez", "Morticia"]}`)
	var m FamilyMember
	json.Unmarshal(b, &m)
	fmt.Printf("the family member is: %s", b)
}