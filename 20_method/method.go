package main

import "fmt"

// 方法定义：func关键字 (接受者) 方法名(如参)  (返回值) {方法体}
// 怎么定义方法
// 1. 定义结构体，或者其他类型
type person struct {
	Name string
	Age  int
}

// 2. 定义方法
// p person
// p *person
// person
// *person
// 接收者实际接收的是person类型的实例化对象
func (p person) printInfo() {
	fmt.Printf("姓名: %s, 年龄: %d\n", p.Name, p.Age)
}

// 指针接收者
func (p *person) setInfo(name string, age int) {
	p.Name = name
	p.Age = age
}
func main() {
	p1 := &person{
		Name: "zhangsan",
		Age:  28,
	}
	p1.printInfo()
	p2 := &person{}
	p2.setInfo("lisi", 20)
	p2.printInfo()
}
