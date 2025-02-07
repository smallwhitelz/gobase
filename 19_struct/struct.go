package main

import "fmt"

// 定义一个结构体
// 结构体的大小写，不需要跨包则小写，但是里面的属性尽量全大写
// 如果属性小写，在做json序列化等操作时，会影响结果
type person struct {
	Name string
	City string
	Age  int
}

func main() {
	// 结构体定义：定义一个对象有哪些属性，有哪些动作（方法）
	// 结构体就是自定义的类型，一般作为全局变量去定义
	var p1 person
	p1.Name = "zhangsan"
	p1.City = "beijing"
	p1.Age = 18
	fmt.Printf("%v %T\n", p1, p1)
	// 建议用法，无初始值
	var p2 = new(person)
	var p3 = &person{}
	// 建议用法，有初始值
	var p4 = &person{
		Name: "张三",
		City: "北京",
		Age:  18,
	}
	fmt.Println(p2, p3, p4)
	// 匿名结构体
	// 常用场景：接收借口参数
	//params := new(struct {
	//	Name     string
	//	Password int
	//})
}
