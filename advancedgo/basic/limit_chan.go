package main

import "fmt"

var limit=make(chan int,3)
var work=make([]func(),10)

func init(){
	for i := 0; i < 10; i++ {
		work[i]= func() {
			fmt.Println("hello")
		}
	}
}

func main(){
	// 限制并发数
	for _, w := range work {
		go func() {
			limit<-1
			w()
			<-limit
		}()
	}
	select {}
}
