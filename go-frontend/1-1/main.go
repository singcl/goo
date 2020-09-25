package main

import (
	"fmt"
	js "syscall/js"
)

func main() {
	btn := js.Global().Get("document").Call("getElementById", "test")

	signal := make(chan int)

	callback := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		fmt.Println(this)
		fmt.Println(args)
		fmt.Println("button clicked")
		return nil
	})

	btn.Set("innerHTML", "changed by go")
	btn.Call("addEventListener", "click", callback)

	<-signal
}
