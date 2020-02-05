package main

import "fmt"

// 模拟搜索引擎，获取到了最先返回的结果就退出
func main() {
	ch := make(chan string, 32)
	go func() {
		ch <- "Bing"
	}()

	go func() {
		ch <- "Google"
	}()

	go func() {
		ch <- "Baidu"
	}()
	fmt.Println(<-ch)
}
