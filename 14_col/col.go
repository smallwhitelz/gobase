package main

import "fmt"

func main() {
	a := (5 + 10) * 10
	fmt.Println(a)
	a += 5 //a=a+5
	fmt.Println(a)
	a %= 5 // 取余，取模
	fmt.Println(a)
}
