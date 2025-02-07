package main

import "fmt"

func main() {
	//make和new
	//相同点：都是为变量申请一片内存
	//不同点：
	// make返回的是普通类型，一般用于slice map channel的内存申请
	// new返回的是指针类型，一般用于自定义的结构体
	a := make([]int, 3, 5)
	a = append(a, 1)
	fmt.Printf("%v %T\n", a, a)
	// 只是举个例子
	b := new([]int)
	*b = append(*b, 10)
	fmt.Printf("%v %T\n", b, b)
}
