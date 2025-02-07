package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func main() {
	u1 := &User{}
	println(u1)
	u1 = new(User)
	println(u1)
	u2 := User{}
	fmt.Printf("%+v \n", u2)
	u2.Name = "zhangsan"
	u2.Age = 18
	println(u2.Name, u2.Age)

	var u3 User
	fmt.Printf("%+v \n", u3)

	var u4 *User
	fmt.Printf("%+v \n", u4)

	u5 := User{Name: "Jerry"}
	fmt.Printf("%+v \n", u5)
	UserList()
}

func UserList() {
	l1 := LinkedList{}
	l1Ptr := &l1
	var l2 = *l1Ptr
	fmt.Printf("%+v \n", l2)
}
