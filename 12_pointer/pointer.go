package main

import "fmt"

func main() {
	// 指针的定义
	// 语法：var关键字 变量名 *类型
	var a = 10
	var b *int = &a
	fmt.Println(a, b, *b)
	fmt.Println(&b, &a)

	//var a int = 127
	//var b float64
	//
	//c := int8(a)
	//
	//fmt.Printf("type:%T,val:%d\n", a, a)
	//fmt.Printf("type:%T,val:%d\n", b, b)
	//fmt.Printf("type:%T,val:%d\n", c, c)
}
