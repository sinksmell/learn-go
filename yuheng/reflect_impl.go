package yuheng

import (
	"reflect"
	"fmt"
)

type  Data  struct{

}

func(*Data)String()string{
	return ""
}

func main(){
	var d *Data
	t:=reflect.TypeOf(d)

	fmt.Println(t.Implements(reflect.TypeOf((*fmt.Stringer)(nil)).Elem()))
}