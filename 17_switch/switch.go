package main

func Switch(status int) string {
	switch status {
	case 1:
		return "成功"
	case 2:
		return "失败"
	default:
		return "未知"
	}
}

func SwitchV1(age int) string {
	switch {
	case age > 18:
		return "成年"
	case age < 6:
		return "未成年"
	}
	return "未知"
}

func main() {
	println(Switch(1))
	println(SwitchV1(20))
}
