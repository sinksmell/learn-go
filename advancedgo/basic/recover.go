package main

import (
	"log"
	"fmt"
)

func main(){
	defer func() {
		// 无法捕获异常
		// 在defer中必须直接调用recover
		// 使用recover的包装函数会失败
		if r:= myRecover();r!=nil{
			fmt.Println(r)
		}
	}()
	panic(1)
}

func myRecover()interface{}{
	log.Println("trace...")
	return recover()
}
