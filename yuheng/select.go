package yuheng

import (
	"fmt"
	"os"
)

func main() {
	a, b := make(chan int, 3), make(chan int)

	go func() {
		v, ok, s := 0, false, ""
		for {
			// select receive
			select {
			case v, ok = <-a:
				s = "a"
			case v, ok = <-b:
				s = "b"
			}
			if ok {
				fmt.Println(s, v)
			} else {
				os.Exit(0)
			}
		}
	}()

	for i := 0; i < 5; i++ {
		// select send
		select {
		case a <- i:
		case b <- i:
		}
	}
	close(a)
	select {}
}
