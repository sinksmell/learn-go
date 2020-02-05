package main

import (
	"net/rpc"
	"net/http"
	"io"
	"net/rpc/jsonrpc"
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
	// 类似rest 规范的接口
	rpc.RegisterName("HelloService",new(HelloService))
	http.HandleFunc("/jsonrpc", func(writer http.ResponseWriter, request *http.Request) {
		var conn io.ReadWriteCloser= struct {
			io.Writer
			io.ReadCloser
		}{
			ReadCloser: request.Body,
			Writer: writer,
		}
		rpc.ServeRequest(jsonrpc.NewServerCodec(conn))
	})
	http.ListenAndServe(":1234",nil)
}
