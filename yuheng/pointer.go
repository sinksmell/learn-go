package yuheng

import (
	"unsafe"
	"time"
)

type  data  struct{
    x [1024*100]byte
}

// uintptr 不能阻止所引用的对象被回收
// unsafe.Pointer 可以阻止所引用对象被回收
func test2()uintptr{
	p:=&data{}
	return uintptr(unsafe.Pointer(p))
}

func main(){
	const N=10000
	cache:=new([N]uintptr)
	for i := 0; i < N; i++ {
		cache[i]=test2()
		time.Sleep(time.Millisecond)
	}
}
