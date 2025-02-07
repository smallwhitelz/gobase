package main

import (
	"fmt"
)

func main() {
	//a := 10
	//if a > 10 && a < 100 {
	//	fmt.Println("hello")
	//} else {
	//	fmt.Println("你好")
	//}
	// 变量写在if条件中，作用域只在if函数内
	if a := 10; a > 10 && a < 100 {
		fmt.Println(a)
		fmt.Println("hello")
	} else {
		fmt.Println("你好")
	}
	//fmt.Println(a)
	// 只在if作用域
	//if err:=json.Unmarshal(); err != nil {
	//    fmt.Println(err)
	//}
	// 在函数作用域
	//err:=json.Unmarshal()
	//if err != nil {
	//    fmt.Println(err)
	//}
}
