package main

import "fmt"

func main() {
	// 切片的定义
	// 语法：var关键字 变量名 []类型
	// 长度可变，且不属于类型
	var a []string
	fmt.Printf("%v %T \n", a, a)
	var b = []string{}
	// 长度可变，只能通过append函数追加的方式
	b = append(b, "a")
	fmt.Printf("%v %T \n", b, b)
	var c = []int{1, 2, 3}
	fmt.Printf("%v %T \n", c, c)
	// 用make声明切片，长度为3，容量为5
	var a1 = make([]int, 3, 5)
	fmt.Println(a1, len(a1), cap(a1))

	// slice是引用类型，指向共同的引用地址，当一个发生改变的时候，另一个也会改变
	var slice = []int{1, 2, 3}
	slice2 := slice
	slice2[0] = 3
	fmt.Println(slice2, slice)

	// 切片的长度和容量
	// 长度：元素个数
	// 容量：切片的第一个元素，到底层数组最后一个元素的长度
	x := [8]int{1, 2, 3, 4, 5, 6, 7, 8}
	y := x[3:6]
	fmt.Printf("%v %T \n", x, x)
	fmt.Printf("%v %T \n", y, y)
	fmt.Println(len(y), cap(y))
	z := y[:5]
	fmt.Println(z, len(z), cap(z))

	// append添加1个元素(扩容)
	slice5 := []int{}          // 底层空数组[0]int
	slice5 = append(slice5, 1) // 换了个[1]int{1}的数组
	fmt.Println(slice5)
	for i := 0; i < 10; i++ {
		slice5 = append(slice5, i)
	}
	fmt.Println(slice5)
	// append添加多个(添加切片或数组)
	slice6 := []int{10, 20, 30}
	slice5 = append(slice5, slice6...)
	fmt.Println(slice5)
	// 切片中删除元素
	slice7 := append(slice5[:5], slice5[6:]...)
	fmt.Println(slice7)
	fmt.Println("-------------------")
	ShowSlice()
}

func ShowSlice() {
	s1 := []int{1, 2, 3, 4}
	s2 := s1[2:]
	fmt.Printf("s1:%v len(s1): %d cap(s1): %d\n", s1, len(s1), cap(s1))
	fmt.Printf("s2:%v len(s2): %d cap(s2): %d\n", s2, len(s2), cap(s2))
	s2[0] = 100
	fmt.Printf("s1:%v len(s1): %d cap(s1): %d\n", s1, len(s1), cap(s1))
	fmt.Printf("s2:%v len(s2): %d cap(s2): %d\n", s2, len(s2), cap(s2))
	s2 = append(s2, 999)
	fmt.Printf("s1:%v len(s1): %d cap(s1): %d\n", s1, len(s1), cap(s1))
	fmt.Printf("s2:%v len(s2): %d cap(s2): %d\n", s2, len(s2), cap(s2))
	s2[1] = 19999
	fmt.Printf("s1:%v len(s1): %d cap(s1): %d\n", s1, len(s1), cap(s1))
	fmt.Printf("s2:%v len(s2): %d cap(s2): %d\n", s2, len(s2), cap(s2))
}
