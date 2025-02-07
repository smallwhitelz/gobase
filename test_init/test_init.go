package test_init

import "fmt"

// 跨包调用
var Age = 10

// 小写不能跨包调用
var name = "zhangsan"

func init() {
	fmt.Println("这里是test_init函数")
}
