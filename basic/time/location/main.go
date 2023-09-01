package main

import (
	"fmt"
	"time"
)

func main() {
	//fmt.Print(time.Now().Before(time.Time{}))

	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		panic(err)
	}

	t, err := time.ParseInLocation("2006-01-02 15:04:05", "2023-09-02 12:23:34", loc)
	if err != nil {
		panic(err)
	}

	fmt.Println(t.String())
}
