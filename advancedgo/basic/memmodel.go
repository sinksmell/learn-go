package main

// 不同goroutine之间不保证顺序一致性内存模型

var a string
var done bool

func setup(){
	a="hello,world!"
	done =true
}

func main(){
	go setup()
	for !done  {}
	print(a)
}
