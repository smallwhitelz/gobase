package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	var a int8 = 4
	var b int16 = 4
	var c int32 = 4
	var d int64 = 4
	var e int = 4
	//var f uint =4
	fmt.Printf("%T %d\n", a, a)
	fmt.Printf("%T %d\n", b, b)
	fmt.Printf("%T %d\n", c, c)
	fmt.Printf("%T %d\n", d, d)
	fmt.Printf("%T %d\n", e, e)
	//通过reflect打印类型
	fmt.Println(reflect.TypeOf(e))

	// GO Math包中有极值也就是最大值和最小值
	math.Asin(1)
}
