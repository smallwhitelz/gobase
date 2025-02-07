package main

import "fmt"

func main() {
	//数组的定义
	//语法：var关键字 变量名 [长度]类型
	var a [5]int // 只声明，不赋值
	a[0] = 1
	a[1] = 2
	a[3] = 3
	//a[5]=4  数组下标越界
	var b = [3]string{"zhangsan", "lisi", "wangwu"} // 既声明又赋值
	fmt.Printf("%v %T \n", a, a)
	fmt.Printf("%v %T \n", b, b)
	// 声明的时候不定长，必须要有初始值
	// 根据初始值长度自动定义数组长度，不能超过这个定义的长度
	var c = [...]int{1, 2, 3, 4, 5}
	fmt.Printf("%v %T \n", c, c)

	// 不可以这么用，会报错
	//var d [...]int
	//d[0]=1
	//fmt.Println(d)

	// 数组的遍历
	var d = [...]string{"北京", "上海", "深圳"}
	for i := 0; i < len(d); i++ {
		fmt.Println(d[i])
	}
	for k, v := range d {
		fmt.Println(k, v)
	}

	// array是值类型，是将值赋值给arr2，不是内存地址，所以arr2改变时，arr不会改变
	var arr = [3]int{1, 2, 3}
	arr2 := arr
	arr2[0] = 3
	fmt.Println(arr2, arr)
}
