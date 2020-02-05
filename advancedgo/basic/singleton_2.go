package main

import (
	"sync"
	"sync/atomic"
)

type  Once  struct{
    m sync.Mutex
    done uint32
}

func(o *Once)Do(f func()){
	if atomic.LoadUint32(&o.done)==1{
		return
	}
	o.m.Lock()
	defer o.m.Unlock()
	if o.done==0{
		defer atomic.StoreUint32(&o.done,1)
		f()
	}
}

// 重新实现单例模式
var(
	instance *singleton
	once sync.Once
)

func GetInstance()*singleton{
	once.Do(func() {
		instance=&singleton{}
	})
	return instance
}

func main(){

}
