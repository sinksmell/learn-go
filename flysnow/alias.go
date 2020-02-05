package main

import "fmt"

type myint int	// define a new type
type myint2 = int	// type alias

type  User  struct{
}

type myUser =User

func(u User)Hello(){
	fmt.Println("hello")
}

func main(){
	var i int=0
	// var i1 myint =i  error
	var i2 myint2=i
	fmt.Println(i2)

}
