type myPrintInterface interface {
	print()
}

func f3(x myInterface) {
	x.(myPrintInterface).print()
}

// x 转换为 myPrintInterface 类型是完全动态的：只要 x 的底层类型（动态类型）定义了 print 方法这个调用就可以正常运行。