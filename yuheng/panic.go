package yuheng

import "fmt"

func recv(){
	defer func() {
		if err:=recover();err!=nil{
			fmt.Println(err.(string))
		}
	}()
	panic("panic error!")
}


// 仅最后一个错误可以被捕获
func recv2(){
	defer func() {
		fmt.Println(recover())
	}()

	defer func() {
		panic("defer panic!")
	}()

	panic("panic error!")
}

func main(){
	recv()
	recv2()
}