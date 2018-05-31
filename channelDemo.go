package main

import (
	"fmt"
	"time"
)

func produce(p chan<- int) {
	for i := 0; i < 10; i++ {
		p <- i
		fmt.Println("子线程send:", i)
	}
}
func consumer(c <-chan int) {
	for i := 0; i < 10; i++ {
		v := <-c
		fmt.Println("子线程receive:", v)
	}
}
func main() {
	ch := make(chan int)
	go produce(ch)
	go consumer(ch)
	fmt.Println("我是主线程")
	time.Sleep(1 * time.Second)
}
