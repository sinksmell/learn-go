package yuheng

import (
	"sync"
	"fmt"
	"runtime"
)

func main(){
	// goexit 立刻终止当前goroutine的执行 但是确保已经注册的defer  会被执行
	wg:=sync.WaitGroup{}
	wg.Add(1)

	go func() {
		defer  wg.Done()
		defer fmt.Println("A defer func!")
		func(){
			defer  fmt.Println("B defer func!")
			runtime.Goexit() // 在这之前的defer 都会被执行
			fmt.Println("B")
		}()
		fmt.Println("A")
	}()

	wg.Wait()

}
