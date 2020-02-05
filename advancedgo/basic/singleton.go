package main

import (
	"sync"
	"sync/atomic"
)

type  singleton  struct{

}

// 使用原子操作 来减少竞争单例时上锁的次数
// 使用数字标志和原子操作 减少锁的使用次数

var(
	instance *singleton
	initialized uint32
	mu sync.Mutex
)

func Instance()*singleton{
	if atomic.LoadUint32(&initialized)==1{
		return instance
	}

	mu.Lock()
	defer mu.Unlock()
	if instance==nil{
		defer atomic.StoreUint32(&initialized,1)
		instance=&singleton{}
	}
	return instance
}