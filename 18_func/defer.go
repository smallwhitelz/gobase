package main

// Defer 函数调用顺序 允许从方法返回的前一刻，执行一段逻辑
// 这个就是defer，也叫做延迟调用
// defer类似于栈，也就是后进先出。也就是，先定义的后执行，后定义的先执行
func Defer() {
	defer func() {
		println("第一个defer")
	}()
	defer func() {
		println("第二个defer")
	}()
}

func DeferClosure() {
	j := 0
	defer func() {
		println(j) // 闭包写法
	}()
	j = 1
}

func DeferClosureV1() {
	j := 2
	defer func(val int) {
		println(val) // 参数传递
	}(j)
	j = 1
}

func DeferReturn() int {
	a := 0
	defer func() {
		a = 1
	}()
	return a
}

func DeferReturnV1() (a int) {
	a = 0
	defer func() {
		a = 1
	}()
	return a
}

func DeferReturnV2() *MyStruct {
	res := &MyStruct{
		Name: "123",
	}
	defer func() {
		res.Name = "456"
	}()
	return res
}

type MyStruct struct {
	Name string
}

func DeferClosureLoopV1() {
	for i := 0; i < 10; i++ {
		defer func() {
			println(i)
		}()
	}
}

func DeferClosureLoopV2() {
	for i := 0; i < 10; i++ {
		defer func(val int) {
			println(val)
		}(i)
	}
}

func DeferClosureLoopV3() {
	for i := 0; i < 10; i++ {
		j := i
		defer func() {
			println(j)
		}()
	}
}

// defer 与 闭包
// 作为参数传入的：定义defer的时候就确定了
// 作为闭包传入的：执行defer对应的方法的时候才确定
func main() {
	//Defer()
	//DeferClosure()
	DeferClosureV1()
	println(DeferReturn())
	println(DeferReturnV1())
	println(DeferReturnV2().Name)
	DeferClosureLoopV1()
	println("------------------")
	DeferClosureLoopV2()
	println("------------------")
	DeferClosureLoopV3()
}
