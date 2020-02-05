package main

type  HelloService  struct{
}

func(h *HelloService)Hello(request *String,reply *String)error{
	reply.Value="hello:"+request.GetValue()
	return nil
}
