package main

import "fmt"

func P(fn func()){
	fmt.Printf("%p",fn)
}

func main(){
	fn:=func() {
		fmt.Println("hello")
	}
	fmt.Printf("%p\n",fn)
	P(fn)
}
