package main

import (
	"fmt"
	"time"
)

func fabonacciMethod(n int) int {
	if n < 2 {
		return n
	}
	return fabonacciMethod(n-1) + fabonacciMethod(n-2)
}
func main() {
	//	var currentTiem int64
	currentTiem := time.Now().UnixNano()
	fmt.Println(currentTiem)
	var n int
	for n = 0; n < 50; n++ {
		fmt.Println(n, fabonacciMethod(n))
	}
	now := time.Now().UnixNano() - currentTiem
	//	fmt.Println(time.Now().Unix())
	fmt.Println("用时：", now)
}
