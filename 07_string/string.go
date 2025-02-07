package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	s1 := "hello"
	s2 := "你好"
	s3 := `
        第一行
        第二行
        第三行`
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)

	// 字符串常用操作
	var str = "this is string"
	// 长度
	fmt.Println(len(str))
	var str2 = "你好"
	var str3 = "golang"
	// 拼接
	fmt.Println(str2 + "," + str3)
	a := fmt.Sprintf("%s,%s", str2, str3)
	fmt.Println(a)
	// 分割，有字符串分割为数组
	var str4 = "hello,world,golang"
	var arr = strings.Split(str4, ",")
	fmt.Println(arr)
	// join，数组合并为字符串
	var str5 = strings.Join(arr, "-")
	fmt.Println(str5)

	// 字符串遍历
	str6 := "hello 中国"
	// 普通for循环输出的是byte类型
	for i := 0; i < len(str6); i++ {
		fmt.Printf("%v %T\n", str6[i], str6[i])
	}
	// for range循环字符串，输出的是rune类型
	for _, value := range str6 {
		fmt.Printf("%v %T\n", value, value)
		fmt.Println(string(value))
	}

	// int转string
	var num1 int = 10
	//将int类型转换为string类型
	str7 := strconv.Itoa(num1)
	fmt.Printf("%v %T\n", str7, str7)

	// float转string
	var num2 float64 = 20.123123
	// 参数1 要转换的值
	// 参数2 格式化类型
	// 参数3 保留的小数位数
	// 参数4 float的位数
	str8 := strconv.FormatFloat(num2, 'f', 2, 64)
	fmt.Printf("%v %T\n", str8, str8)

	// bool转string
	str9 := strconv.FormatBool(true)
	fmt.Printf("%v %T\n", str9, str9)

	// int64转string
	var num3 int64 = 20
	str10 := strconv.FormatInt(num3, 10) // 这里的10代表10进制
	fmt.Printf("%v %T\n", str10, str10)

	// string转int
	var str11 = "10"
	intNum, _ := strconv.Atoi(str11)
	fmt.Printf("%d %T\n", intNum, intNum)
}
