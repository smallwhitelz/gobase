package main

var i = 0

func Closure(name string) func() string {
	// 闭包：方法+它绑定的运行上下文
	// name变量
	// 方法本身
	// 上下文指方法之外的，所以name和i都是上下文，world不是上下文
	// 闭包如果使用不当可能会引起内存泄漏的问题。即一个对象被闭包引用的话，它是不会被垃圾回收的。
	return func() string {
		i++
		world := "world"
		return name + "hello" + world
	}
}

func ClosureInvoke() {
	c := Closure("zl")
	println(c())
}
