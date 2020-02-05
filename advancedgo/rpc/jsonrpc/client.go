package main

import (
	"net"
	"log"
	"net/rpc"
	"net/rpc/jsonrpc"
	"fmt"
)

func main(){
	conn,err:=net.Dial("tcp","localhost:2333")
	if err!=nil{
		log.Fatal("net.Dial",err)
	}
	client:=rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))
	var reply string
	err=client.Call("HelloService.Hello","hello",&reply)
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println(reply)
}
