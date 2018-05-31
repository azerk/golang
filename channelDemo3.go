package main

import (
	"fmt"
	"time"
)

func load(c chan int) {
	for i := 0; i < 100000; i++ {
		fmt.Print(i)
		if i == 99999 {
			fmt.Println(i)
		}
	}
	c <- 0
}
func load1() {
	for i := 0; i < 100000; i++ {
		fmt.Print(i)
		if i == 99999 {
			fmt.Println(i)
		}
	}

}
func main() {

	c := make(chan int, 10)
	now := time.Now().UnixNano()
	go load(c)

	fmt.Println(<-c)
	now1 := time.Now().UnixNano()
	fmt.Println("用时1:", now1-now, now, now1)

	now2 := time.Now().UnixNano()
	load1()
	now3 := time.Now().UnixNano()
	fmt.Println("用时2:", now3-now2, now2, now3)

	fmt.Println("胜利了")
}
