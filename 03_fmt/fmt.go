package main

import "fmt"

func main() {
	// fmt包的使用
	// fmt.Println
	// 一次输入多个值的时候，中间会有空格，Println会自动换行
	fmt.Println("a", "b", 100)
	fmt.Println("a", "b", 100)
	// fmt.Print，则不会
	fmt.Print("a", "b", 100)
	fmt.Print("a", "b", 100)
	// fmt.Printf，格式化输出，可以将不同的类型，拼接成字符串输出
	// %s 字符串 %d 数字  %v值(各种类型的值 )  %T 类型
	fmt.Printf("\n %s %d %v %T", "abc", 10, true, 10)
	//fmt.Sprintf，跟Printf一样，只是多了个返回值
	// 一般用于不同类型拼接成字符串
	a := fmt.Sprintf("\n %s %d %v %T", "abc", 20, true, 20)
	fmt.Printf(a)
}
