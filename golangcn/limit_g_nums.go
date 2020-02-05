package main

import "sync"

// 如何控制并发执行的 Goroutine 的最大数目？
// 使用无缓冲的chan来控制


type Worker interface {
	Task()
}

// pool为协程池
type Pool struct {
	work chan Worker
	wg sync.WaitGroup
}

// 根据g的最大数量来进行限制
func NewPool(maxG int)*Pool{
	p:=&Pool{
		work:make(chan Worker),
	}
	p.wg.Add(maxG)
	for i := 0; i < maxG; i++ {
		go func() {
			for w := range p.work {
				w.Task()
			}
			p.wg.Done()
		}()
	}
	return p
}

func (p *Pool)Run(w Worker){
	p.work<-w
}

func(p *Pool)Shutdown(){
	close(p.work)
	p.wg.Wait()
}

func main(){

}