// 非结构体类型上的方法
package main

import "fmt"

type IntVector []int

func main() {
	fmt.Println(IntVector{1, 2, 3, 5}.Sum()) // 输出是6
}

func (v IntVector) Sum() (s int) {
	for _, x := range v {
		s += x
	}
	return s
}
