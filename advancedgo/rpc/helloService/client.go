package main

import (
	"net/rpc"
	"log"
	"fmt"
)

func main(){
	client,err:=rpc.Dial("tcp","localhost:2333")
	if err!=nil{
		log.Fatal("dialing:",err)
	}
	var reply string
	err=client.Call("HelloService.Hello","hello",&reply)
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println(reply)
}


