package main

import (
	"fmt"
	"time"
)

func main() {
	go fmt.Println("我也开始跑了")
	fmt.Println("我开始跑了")
	time.Sleep(1 * time.Second)
}
