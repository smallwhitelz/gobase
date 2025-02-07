package main

// YourName 不定参数
// Go支持不定参数。不定参数是指最后一个参数，可以传入任意多个值，注意必须是最后一个参数才可以声明为不定参数
// 不定参数在方法内部可以被当作切片来使用
// Option模式大量应用了不定参数
func YourName(name string, aliases ...string) {
	if len(aliases) > 0 {
		println(aliases[0])
	}
}

func YourNameInvoke() {
	YourName("<NAME>")
	YourName("<NAME>", "<NAME>")
	YourName("<NAME>", "<NAME>", "<NAME>")
}
