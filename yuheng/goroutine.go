package yuheng

import (
	"math"
	"fmt"
	"sync"
)

func sum(id int){
	var x int64
	for i := 0; i < math.MaxUint32; i++ {
		x+=int64(i)
	}
	fmt.Println(id,x)
}

func main(){
	wg:=sync.WaitGroup{}
	wg.Add(2)
	for i := 0; i < 2; i++ {
		go func(id int) {
			defer wg.Done()
			sum(id)
		}(i)
	}
	wg.Wait()
}
