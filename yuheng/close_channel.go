package yuheng

import (
	"sync"
	"fmt"
	"time"
)

// 用closed channel 通知协程退出
// go 1.8之后用context代替了
func main(){
	wg:=sync.WaitGroup{}
	quit:=make(chan struct{})

	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			task:= func() {
				fmt.Println(id,time.Now().Nanosecond())
				time.Sleep(time.Second)
			}
			for{
				select {
				case <-quit:
					// 从已关闭的channel中读出nil
					return
				default:
					task()

				}
			}
		}(i)
	}

	time.Sleep(time.Second*5)
	close(quit)
	wg.Wait()
}
