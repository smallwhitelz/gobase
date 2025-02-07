package main

import "fmt"

func main() {
	str := "中国hello"
	// rune类型是一个int32
	a := []rune(str)
	// byte类型是一个uint8
	b := []byte(str)
	fmt.Printf("%v %T\n", a, a)
	fmt.Printf("%v %T\n", b, b)
	// 单引号声明rune类型
	c := '中'
	//c := []int32{'中', '国'}
	fmt.Printf("%v %T\n", c, c)
}
