package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")

	// var x = [5]int{1, 2, 3, 4, 5}
	var x [5]int // 基本类型会自动初始化零值
	xs := x[:]   // 从数组创建一个切片 [引用类型]
	xs[0] = 43   // 修改切片 => 切片底层数据 数组也会被修改

	// Spaces are always added between operands and a newline is appended.
	// 总是在操作数之间添加空格，并追加换行符
	fmt.Println(xs)
	// Spaces are added between operands when neither is a string
	// 当两个字符串都不存在时，在操作数之间添加空格。
	fmt.Print(x)
}
