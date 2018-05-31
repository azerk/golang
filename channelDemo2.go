package main

import (
	"fmt"
	//	"time"
)

func produce(p chan<- int) {
	for i := 0; i < 500; i++ {
		p <- i
		fmt.Println("send:", i)
	}
}

//func consumer(c <-chan int) {
//	for i := 0; i < 100; i++ {
//		v := <-c
//		fmt.Println("receive:", v)
//	}
//}
func main() {
	ch := make(chan int, 72)
	go produce(ch)
	//	go consumer(ch)
	for i := 0; i < 500; i++ {
		v := <-ch
		fmt.Println("receive:", v)
	}

	//	time.Sleep(1 * time.Second)
	fmt.Println("完成了，胜利了")
}
