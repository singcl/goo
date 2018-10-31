package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("protect 函数调用函数参数 g 来保护调用者防止从 g 中抛出的运行时 panic，并展示 panic 中的信息")
	protect(func() {
		panic("panic 错误 recover 测试！")
	})
}

func protect(g func()) {
	defer func() {
		log.Println("done")
		// Println executes normally even if there is a panic
		if err := recover(); err != nil {
			log.Printf("run time panic: %v", err)
		}
	}()

	log.Println("start")
	g() //   possible runtime-error
}

