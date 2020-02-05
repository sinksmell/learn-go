package yuheng

import "fmt"


// 小于1024时slice 2倍扩容
// 大于1024时slice 每次增加1/4
func main(){
	s:=make([]int,0)
	c:=cap(s)
	fmt.Println(c)
	for i := 0; i < 5000; i++ {
		s=append(s,i)
		if n:=cap(s);n>c{
			fmt.Printf("cap: %d -> %d\n",c,n)
			c=n
		}
	}
}
