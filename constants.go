package main

import "fmt"
import "math"

const ss string = "constant"

func main() {
	fmt.Println(ss)

	const n = 500000000
	const d = 3e+20 / n

	// 在函数内部一般用 := 进行变量声明和赋值
	// 当然用var 也是可以的。 var 一般用于声明全部变量
	// go中不允许函数嵌套
	var xy func(i, k int) int
	var h int

	xy = Sum
	h = xy(1,2)
	fmt.Println(h)
	fmt.Println(d)
	fmt.Println(int64(d))
	fmt.Println(math.Sin(n))
}

func Sum(a, b int) int { return a + b }
