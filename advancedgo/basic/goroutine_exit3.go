package main

import (
	"context"
	"sync"
	"fmt"
	"time"
)

// 使用context进行管理goroutine的退出
func worker(ctx context.Context, wg *sync.WaitGroup) error {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			fmt.Println("worker is running")
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.TODO())
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go worker(ctx, &wg)
	}

	time.Sleep(time.Second)
	cancel()
	wg.Wait()
}
