package  main

import "fmt"

func badCall() {
	panic("bad end!")
}

//defer-panic-recover 在某种意义上也是一种像 if，for 这样的控制流机制。

func test() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Printf("Panicing %s \r\n", e)
		}
	}()

	badCall()

	// 上面badCall 函数抛出了panic 所有后面的代码不会执行
	fmt.Printf("After bad call \r\n")
}

func main() {
	fmt.Printf("Calling test\r\n")
	test()
	fmt.Printf("Test completed\r\n")
}

