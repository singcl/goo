
// js深入系列专题：https://github.com/mqyqingfeng/Blog
// js词法作用域
// js函数的作用域（内部属性[[scope]]）是在定义阶段确定的而不是在调用阶段确定的

var a

function main() {
	a = "G"
	console.log(a)
	f1()
}

function f1() {
	var a = "O"
	console.log(a)
	f2()
}

function f2() {
	console.log(a)
}

main()