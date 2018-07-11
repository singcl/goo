// 在 7.6.6 中我们看到了能被搜索和排序的 int 数组、float 数组以及 string 数组，那么对于其他类型的数组呢，是不是我们必须得自己编程实现它们？
// Vector 中存储的所有元素都是 Element 类型，要得到它们的原始类型（unboxing：拆箱）需要用到类型断言。
package main

type Element interface{}

type Vector struct {
	a []Element
}

func (p *Vector) At(i int) Element {
	return p.a[i]
}

func (p *Vector) Set(i int, e Element) {
	p.a[i] = e
}

func main() {}
