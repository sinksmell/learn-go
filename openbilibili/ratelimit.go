package main

import "fmt"

type  Book  struct{
    name string
}

func Test(b *Book){
	fmt.Println("hello")
	return
}

func main(){
	var done func(i int,s string)
	done(1,"hello")
}

