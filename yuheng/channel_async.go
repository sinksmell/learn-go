package main

import "fmt"

func main() {
	// 有缓冲的channel 是异步的 当缓冲区已满的时候才会导致阻塞
	data := make(chan int, 10)
	exit := make(chan struct{})

	for i := 0; i < 10; i++ {
		data <- i
	}
	close(data)
	fmt.Println("send over!")

	go func() {
		for d := range data {
			fmt.Println(d)
		}
		fmt.Println("recv over!")
		exit <- struct{}{}
	}()

	<-exit
}
