package main

import (
	"fmt"
	"time"
)

func producter(ch chan<- int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
}

func consumer(ch <-chan int){
	for v := range ch {
		fmt.Println(v)
	}
}

func main() {
	// 缓冲区
	ch := make(chan int, 64)
	go producter(ch)
	go consumer(ch)
	time.Sleep(time.Second*2)
}
