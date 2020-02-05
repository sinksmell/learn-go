package yuheng

import (
	"sync"
	"fmt"
	"runtime"
)

// Gosched 让出底层线程，将当前goroutine暂停，放回队列等待下次运行
// 类似线程的yield函数
func main(){
	wg:=sync.WaitGroup{}
	wg.Add(2)
	runtime.GOMAXPROCS(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 6; i++ {
			fmt.Println(i)
			if i==3{
				runtime.Gosched()
			}
		}
	}()

	go func() {
		defer wg.Done()
		fmt.Println("Hello,world!")
	}()
	wg.Wait()
}
