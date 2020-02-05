package main

import (
	"net/rpc"
	"net"
	"log"
)

type  HelloService  struct{
}

// golang rpc 规则
// 1.可导出的方法
// 2.只能有两个可序列化的参数， 且第二个参数时指针类型
// 3.返回一个error 类型
func(h *HelloService)Hello(request string ,reply *string)error{
	*reply="hello:"+request
	return nil
}


func main(){
	rpc.RegisterName("HelloService",new(HelloService))
	listener,err:=net.Listen("tcp",":2333")
	if err!=nil{
		log.Fatal("ListenTCP error:",err)
	}
	conn,err:=listener.Accept()
	if err!=nil{
		log.Fatal("Accept error:",err)
	}
	rpc.ServeConn(conn)
}

