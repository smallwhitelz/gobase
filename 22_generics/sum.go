package main

import (
	"errors"
	"fmt"
	"io"
)

func Sum[T Number](vals []T) T {
	var res T
	for _, v := range vals {
		res = res + v
	}
	return res
}

func Max[T Number](vals []T) T {
	t := vals[0]
	for i := 1; i < len(vals); i++ {
		if vals[i] > t {
			t = vals[i]
		}

	}
	return t
}

func Min[T Number](vals []T) T {
	t := vals[0]
	for i := 1; i < len(vals); i++ {
		if vals[i] < t {
			t = vals[i]
		}

	}
	return t
}

func DeleteAt[T any](src []T, index int) ([]T, error) {
	length := len(src)
	if index < 0 || index >= length {
		return nil, errors.New("index out of range")
	}
	src = append(src[:index], src[index+1:]...)
	src = Shrink(src)
	return src, nil
}

// 缩容
func Shrink[T any](src []T) []T {
	c, l := cap(src), len(src)
	newCap, changed := CapDown(c, l)
	if !changed {
		return src
	}
	newSrc := make([]T, 0, newCap)
	newSrc = append(newSrc, src...)
	return newSrc
}

func CapDown(c int, l int) (int, bool) {
	// 这里采用java的hashmap底层数组的一个阈值，因为我以前学java，直接拿来用了
	if c < 64 {
		return c, false
	}
	// 如果长度大于等于容量的1/4，则缩容
	newCap := c / 4
	if newCap > l {
		return newCap, true
	}
	return c, false
}

func main() {
	//src, _ := DeleteAt[int]([]int{1, 2, 3, 4}, 2)
	//fmt.Printf("%+v\n", src)
	//src2, err := DeleteAt[string]([]string{"tom", "jerry", "luck"}, 0)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Printf("%+v\n", src2)

	src3 := make([]int, 0, 4)
	src3 = append(src3, 1, 2, 3, 4)
	ans, err := DeleteAt(src3, 2)
	ans, err = DeleteAt(ans, 2)
	ans, err = DeleteAt(ans, 1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("数组：%+v，cap：%d，长度: %d\n", ans, cap(ans), len(ans))

}

// ~代表衍生类型，Integer就是int的衍生类型
type Integer int

type Number interface {
	~int | int32 | int64
}

func UseSum() {
	res := Sum[int]([]int{123, 123})
	println(res)
	res1 := Sum[Integer]([]Integer{123, 123})
	println(res1)
}

func Closable[T io.Closer]() {
	var t T
	t.Close()
}
