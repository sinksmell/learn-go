package main

import (
	"sync"
	"fmt"
)

func main(){
	var mu sync.Mutex
	mu.Lock()
	go func() {
		fmt.Println("hello")
		mu.Unlock()
	}()
	// 第二次锁定时，如果锁未释放，那么 主线程就阻塞
	mu.Lock()
}

