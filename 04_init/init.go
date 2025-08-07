package main

import (
	"fmt"
	"gobase/test_init"
)

// 在main函数之前自动调用
// init函数可以有多个
//func init() {
//	fmt.Println("我是init1函数")
//}
//
//func init() {
//	fmt.Println("我是init2函数")
//}

// main函数作为程序主入口，一个程序只能有一个
// 正式项目中，非常非常不建议使用init函数！！！！！！！
// 如果你项目中有多个init，那么他会按照包的顺序去执行，这样很容易出现问题
func main() {
	fmt.Println(test_init.Age)
	fmt.Println("我是main函数")
}
