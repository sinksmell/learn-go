package yuheng

import "fmt"

// 可变长参数
func test(s string ,n ...int)string{
	// 将可变长参数作为slice处理
	var x int
	for _,i:=range n  {
		x+=i
	}
	return fmt.Sprintf(s,x)
}
func main(){
	fmt.Println(test("sum: %d",1,2,3,4,5,6))
	// 传入slice时必须要展开
	arr:=[]int{1,2,3,4,5,6}
	fmt.Println(test("sum: %d",arr...))
}
