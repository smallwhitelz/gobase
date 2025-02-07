package main

import "fmt"

func main() {
	// 常量定义
	// 语法：const关键字 变量名 类型 = 值
	// 只定义一次，不可改变，定义了可以不使用
	const a = 1
	// 批量定义
	const (
		b = 2
		c
		d
	)
	fmt.Println(a, b, c, d)
	// 常量命名方式：全大写，用下滑线隔开
	const HTTP_STATUS_OK = 200

}

const (
	Status1 = iota
	Status2

	Status3
	Status4 = 5

	Status5
)
