package service

import "fmt"

var (
	Name = "zhangsan"
	Age  = 10
	Sex  = "male"
	mail = "mail.com.qq.cn"
)

func Get() {
	fmt.Println(mail)
}
