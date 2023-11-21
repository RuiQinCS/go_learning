package main

import "fmt"

// 测试 子goroutine 未结束，main 会不会结束

func main() { // 代表gin.Engine
	stop := make(chan bool)

	go func() { // 代表pubDetail
		go func() { // 代表pubDetail中的goroutine
			for i := 0; i < 1000; i++ {
				fmt.Println(i)
			}
			stop <- true
		}()
	}()

	<-stop
}
