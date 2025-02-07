package main

import "fmt"

func main() {
	a := 10
	b := 5
	fmt.Println(intSum(a, b))
	fmt.Println(sumAanMin(a, b))
	// 匿名函数
	func_test := func(x, y int) int {
		return x - y
	}
	fmt.Println(func_test(a, b))

	// 闭包函数调用
	fn := AddInt()
	//fmt.Println(AddInt()(1)// 1
	//fmt.Println(AddInt()(1)// 1
	//fmt.Println(AddInt()(1)// 1
	fmt.Println(fn(1)) // 1
	fmt.Println(fn(1)) // 2
	fmt.Println(fn(1)) // 3

	_, _ = Func3(1, 2) // 本质是方法的调用，需要有两个接收参数
	myFunc3 := Func3   // myFunc3本质是一个变量，只不过这个变量指向了Func3
	myFunc3(1, 2)
}

// AddInt 闭包函数，返回值是匿名函数
// 会发生内存逃逸，正常情况下函数执行完，会进行内存回收
// 但是匿名函数会发生内存逃逸，定义的变量会逃到heap堆
func AddInt() func(i int) int {
	var num int
	return func(i int) int {
		num += 1
		return num
	}
}

// 定义函数：func关键字 函数名(接受参数)  (返回值)  {函数体}
// 如参一定要带变量名
// 返回值可以不带变量名
// func intSum(x, y int) (sum int)
func intSum(x, y int) string {
	return fmt.Sprintf("%d + %d = %d", x, y, x+y)
}

func sumAanMin(x, y int) (sum, min int) {
	sum = x + y
	min = x - y
	return
}

// Func2
//
//	返回值可以带名字
//	要么都带名字，要么都不带名字
func Func2(a, b int) (str string, err error) {
	str = "hello"
	return
}

// Func3 虽然带名字，但是也可以不用其名字
func Func3(a, b int) (str string, err error) {
	str = "hello"
	return "", nil
}

func Func4() {
	// 该方法是一个局部方法，只能在Func4方法中使用
	fn := func(name string) string {
		return name + "hello"
	}
	str := fn("zl")
	println(str)
}

// Func5 匿名方法直接调用
func Func5() {
	//
	fn := func(name string) string {
		return name + "hello"
	}("zl")
	println(fn)
}
