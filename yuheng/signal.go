package yuheng

import (
	"os/signal"
	"fmt"
	"os"
	"syscall"
)

func main(){
	// 信号的到来时异步的 可以实现实时控制 使用 interrupt实现
	c:=make(chan os.Signal,1)
	signal.Notify(c,syscall.SIGINT,syscall.SIGTERM,syscall.SIGTSTP)
	for{
		s:=<-c
		fmt.Printf("get a signal %s\n",s.String())
		switch s{
		case syscall.SIGINT,syscall.SIGTERM:
			fmt.Println("exit!")
			return
		case syscall.SIGTSTP:
			fmt.Println("stop!")
		default:
			fmt.Println("hello!")
		}
	}

}