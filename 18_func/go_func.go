package main

import (
	"fmt"
	"time"
)

func main() {

	for i := 0; i < 5; i++ {
		// 启动一个带参数的匿名 goroutine，把当前的 i 值传进去
		j := i
		go func() {
			// val 是函数参数，已经在调用时拷贝好了
			fmt.Println("B:", j)
		}()
	}

	// 等一会儿，让所有 goroutine 都有机会跑
	time.Sleep(100 * time.Millisecond)
}
