package yuheng

import (
	"reflect"
	"fmt"
)

// 利用Make、New函数 实现近似泛型的操作
var (
	Int    = reflect.TypeOf(0)
	String = reflect.TypeOf("")
)

func Make(T reflect.Type, fptr interface{}) {
	swap := func(in []reflect.Value) []reflect.Value {
		return []reflect.Value{
			reflect.MakeSlice(
				reflect.SliceOf(T),
				int(in[0].Int()),
				int(in[1].Int()),
			),
		}
	}
	// 省略算法内容
	fn:=reflect.ValueOf(fptr).Elem()
	v:=reflect.MakeFunc(fn.Type(),swap)
	fn.Set(v)
}

func main() {
	var makeints func(int,int)[]int
	var makestrings func(int ,int)[]string
	Make(Int,&makeints)
	Make(String,&makestrings)

	x:=makeints(5,10)
	fmt.Printf("%#v\n",x)
	s:=makestrings(3,10)
	fmt.Printf("%#v\n",s)
}
