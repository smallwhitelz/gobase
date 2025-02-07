package main

import (
	"fmt"
	// 只导入不使用，只会声明包下面的全局变量，以及执行init函数
	_ "gobase/service"
	// 包名的别名
	svc "gobase/service"
	//"gobase/service"
)

func main() {
	//fmt.Println(service.Age)
	//service.Get()
	fmt.Println(svc.Age)
	svc.Get()
}
