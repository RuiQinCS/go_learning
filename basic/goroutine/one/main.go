package main

import "fmt"

// 测试 子goroutine 未结束，main 会不会结束
func main() {

	go func() {
		for i := 0; i < 1000; i++ {
			fmt.Println(i)
		}
	}()

}
