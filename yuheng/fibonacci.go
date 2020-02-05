package yuheng

import "time"
import "fmt"

func main() {
	for i := 0; i < 7; i++ {
		start := time.Now()
		Fibonacci(i)
		end := time.Now()
		deta := end.Sub(start)
		fmt.Println(deta)
	}

}

func Fibonacci(n int) int {
	if n < 0 {
		return -1
	} else if n == 0 || n == 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}
