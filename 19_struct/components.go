package main

type NameI interface {
	Name() string
}

type Outer struct {
	Inner
}

type Outer1 struct {
	*Inner
}

func (i Outer1) Name() string {
	return "outer" // 组合不是继承，更没有多态
}

type Inner struct {
}

func (i Inner) Name() string {
	return "inner"
}

func (i Inner) hello() {
	println("hello 我是,", i.Name())
}

func main() {
	var o Outer
	o.hello()

	o1 := Outer1{
		Inner: &Inner{},
	}
	o1.hello()
}
