package main

import (
	"fmt"
	"time"
)

// 用select 和chan 来实现goroutine的退出
// 由于chan的收发操作比较消耗资源
// 可以用close(ch)来代替退出信号的发送
// close可以用在多个goroutine之间的信号广播
// 程序依然不够稳健 因为goroutine退出后的清理工作不能保证被完成
func worker(cancel chan bool,num int){
	for{
		select {
		case <-cancel:
			return
		default:
			fmt.Println(num,"worker is running")
		}
	}
}

func main(){
	cancel:=make(chan bool)
	for i := 0; i < 5; i++ {
		go worker(cancel,i)
	}
	time.Sleep(time.Second)
	close(cancel)
}
