package main

import (
	"sync"
	"fmt"
	"time"
)

// 在goroutine_exit的基础上添加waitGroup
// 让主线程在goroutine完成清理后退出

func worker(wg *sync.WaitGroup,cancel <- chan bool,num int ){
	defer wg.Done()
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
	wg:=sync.WaitGroup{}
	cancel:=make(chan bool)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go worker(&wg,cancel,i)
	}
	time.Sleep(time.Second)
	close(cancel)
	wg.Wait()
}