package  main

import (
	"os"
	"encoding/json"
	"log"
)
type FamilyMember2 struct {
	Name string
	Age int
	Parents [] string
}

//json 包提供 Decoder 和 Encoder 类型来支持常用 JSON 数据流读写。
// NewDecoder 和 NewEncoder 函数分别封装了 io.Reader 和 io.Writer 接口。

// 下面演示了JSON先编码再解码的完整过程：json数据先解码到指定的GO结构，再从指定的GO结构解码到指定的文件流
func main() {

	// 新建一个json数据强制转换为切片
	b := []byte(`{"Name": "Wednesday", "Age": 6, "Parents": ["Gomez", "Morticia"]}`)

	// 解码json数据到指定的GO 结构
	var m FamilyMember2
	json.Unmarshal(b, &m)

	// OpenFile
	fp, _ := os.OpenFile("family_member.json", os.O_CREATE|os.O_WRONLY, 0666)
	defer fp.Close()

	// 创建一个json编码器
	enc := json.NewEncoder(fp)

	// 将GO 结构 编码到打开的文件流中
	err := enc.Encode(m)
	if err != nil {
		log.Println("Error in encoding json")
	}
}