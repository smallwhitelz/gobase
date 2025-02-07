package main

import (
	"fmt"
)

func main() {
	// for循环
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	var i = 0
	for i < 10 {
		fmt.Println(i)
		i++
	}

	// 模仿while循环
	k := 1
	for {
		if k <= 10 {
			fmt.Println("k =", k)
		} else {
			break
		}
		k++
		//5秒执行一次，类似cron
		//time.Sleep(5*1000)
	}
	// 在go里面最常用的for range 循环
	str := "你好golang"
	for _, value := range str {
		fmt.Println(string(value))
	}
	//switch case
	// 一般情况下case的代码块中用比较简单的表达式或代码
	score := "C"
	switch score {
	case "A": // score = "A"
		fmt.Println("非常棒")
	case "B": //score = "B"
		fmt.Println("优秀")
	default: // score != "A" && score != "B"
		fmt.Println("不及格")
	}

	// continue
	for i := 0; i < 10; i++ {
		// i%2 i能被2整除，余数为0
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}
	fmt.Println("over")
	LoopBug()
	LoopBug2()
	// return
	//for i := 1; i < 10; i++ {
	//	// i%2 i能被2整除，余数为0
	//	if i%2 == 0 {
	//		return
	//	}
	//	fmt.Println(i)
	//}
	//fmt.Println("over")
}

func LoopBug() {
	users := []User{
		{
			name: "Tom",
		},
		{
			name: "Jerry",
		},
	}
	m := make(map[string]*User)
	for _, u := range users {
		m[u.name] = &u
	}
	for name, u := range m {
		println(name, u.name)
	}
}

type User struct {
	name string
}

func LoopBug2() {
	for i := range 7 {
		fmt.Println(i)
	}
	fmt.Println("over")
}
