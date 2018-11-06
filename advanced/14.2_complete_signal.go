package main

//完整的信号量模式
//用完整的信号量模式对长度为N的 float64 切片进行了 N 个 doSomething() 计算并同时完成
func main() {
	type Empty interface {}
	var empty Empty
	var N = 10

	data := make([]float64, N)
	res := make([]float64 , N)
	sem := make(chan Empty, N)

	for i, ix := range data {
		go func(i int, ix float64) {
			res[i] = doSomeThing(i, ix)
			sem <- empty
		}(i ,ix)
	}

	for i := 0; i < N; i++ {
		<- sem
	}
}

func doSomeThing(i int, ix float64) float64 {
	return ix + float64(i)
}