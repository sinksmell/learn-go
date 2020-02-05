package yuheng

import (
	"reflect"
	"fmt"
)

type  Other  struct{

}

type  Adminer  struct{
    Other
}

func (*Other) ToString()  {

}

func (Adminer)test(){}

func main(){
	var u Adminer
	methods:= func(t reflect.Type) {
		for i, n := 0, t.NumMethod(); i < n; i++ {
			f:=t.Method(i)
			fmt.Println(f.Name)
		}
	}

	fmt.Println("--- value interface ---")
	methods(reflect.TypeOf(u))

	fmt.Println("--- pointer interface ---")
	methods(reflect.TypeOf(&u))

}
