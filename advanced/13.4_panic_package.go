package  main

import (
	"fmt"
	"./parse"
)

// 这是所有自定义包实现者应该遵守的最佳实践：
// 1）在包内部，总是应该从 panic 中 recover：不允许显式的超出包范围的 panic()
// 2）向包的调用者返回错误值（而不是 panic）。

//在包内部，特别是在非导出函数中有很深层次的嵌套调用时，对主调函数来说用 panic 来表示应该被翻译成错误的错误场景是很有用的（并且提高了代码可读性）

func main() {
	var example = []string {
		"1 2 3 4 5",
		"100 50 25 12.5 6.25",
		"2 + 2 = 4",
		"1st class",
		"",
	}

	for _, ex := range example {
		fmt.Printf("Parsing %q: \n", ex)
		nums, err := parse.Parse(ex)
		if err != nil {
			fmt.Println("Singcl - ", err) // here String() method from ParseError is used
			continue
		}
		fmt.Println(nums)
	}
}