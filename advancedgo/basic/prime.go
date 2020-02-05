package main

import "fmt"

// 并发版本的素数晒

// 返回生成自然数序列的管道 从2开始
func GenerateNatural() chan int {
	// 同步的管道
	ch := make(chan int)
	go func() {
		for i := 2; ; i++ {
			ch <- i
		}
	}()

	return ch
}

// 管道过滤器 删除能被素数整除的数
func PrimeFilter(in <-chan int,prime int)chan int{
	out:=make(chan int)
	go func() {
		for{
			if i:=<-in;i%prime!=0{
				out<-i
			}
		}
	}()
	return out
}

func main(){
	ch:=GenerateNatural()
	for i := 0; i < 100; i++ {
		prime:=<-ch // 新出现的素数
		fmt.Printf("%v: %v\n",i+1,prime)
		ch=PrimeFilter(ch,prime) // 使用新素数进行过滤
	}
}