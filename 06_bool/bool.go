package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var b = true
	var c bool
	fmt.Println(b, c)
	// 查看变量占用的字节大小
	fmt.Println(unsafe.Sizeof(b))
}
