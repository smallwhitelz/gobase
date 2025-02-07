package main

import "fmt"

// var声明变量可用于函数外，作为全局变量使用
var n2 = 30

// n3 :=40 报错
func main() {
	// 变量定义
	// 语法：var关键字 变量名=值(或者是表达式)
	// 声明出来 必须使用 否则语法检测不通过
	var name string = "zhangsan"
	var age int = 20
	var isOK bool = true
	// 自动推到类型
	var n = 10
	fmt.Println(name, age, isOK, n)
	// := 这种声明方式只能用于函数内，不能用于全局变量声明
	n1 := 10
	fmt.Println(n1)
	fmt.Println(n2)
	// 批量声明变量
	//var username, sex string = "lisi", "男"
	var username, sex string
	username = "lisi"
	sex = "男"
	fmt.Println(username, sex)
	// 批量声明 常用方式
	var (
		a = 1
		b = 2
		c int
	)
	fmt.Println(a, b, c)
	// 变量命名方式：驼峰命名，首字母大写还是小写，取决于要不要跨包调用
	httpStatusOk := 200
	fmt.Println(httpStatusOk)
}
