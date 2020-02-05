package yuheng

import "fmt"

func main(){
	// 默认的是同步的
	// 生产者和消费者要协调进行，否则被阻塞，直到另一方准备好
	data:=make(chan int)
	exit:=make(chan struct{})

	go func() {
		// consumer
		for d:=range data {
			fmt.Println(d)
		}
		fmt.Println("recv over!")
		exit<- struct{}{}
		close(exit)
	}()

	for i := 0; i < 10; i++ {
		data<-i
	}
	close(data)
	fmt.Println("send over!")
	<-exit
}
