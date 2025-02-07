package main

import "fmt"

func main() {
	// map定义
	// 语法：var关键字 变量名 map[键类型]值类型
	// 只声明变量，没有分配内存
	//var mp map[string]string
	//mp["name"] = "zhangsan"
	//fmt.Println(mp) // panic: assignment to entry in nil map
	// 正确的使用姿势，初始化分配内存
	// make一般用于去做引用类型的声明，slice map channel
	var mp2 = make(map[string]string)
	mp2["name"] = "lisi"
	mp2["age"] = "10"
	mp2["sex"] = "male"
	fmt.Println(mp2)
	var mp3 = map[string]string{}
	mp3["name"] = "wangwu"
	fmt.Println(mp3)

	// 判断某个键是否存在
	mp4 := map[string]string{
		"username": "zhangsan",
		"password": "123456",
	}
	v, ok := mp4["username"]
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("not found")
	}
	// 删除某个key和value
	delete(mp4, "username")
	fmt.Println(mp4)
	// 遍历
	mp5 := map[string]int{
		"zhangsan": 30,
		"lisi":     20,
		"wangwu":   10,
	}
	for k, v := range mp5 {
		fmt.Println(k, v)
	}
	// 只遍历key
	for k := range mp5 {
		fmt.Println(k)
	}
	// 只遍历value
	for _, v := range mp5 {
		fmt.Println(v)
	}

	mp6 := map[string]string{
		"key1": "value1",
	}
	println(mp6["key2"]) // 没有的key会返回数据类型的零值，而不是panic
}
