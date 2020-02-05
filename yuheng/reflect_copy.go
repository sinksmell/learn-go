package yuheng

import (
	"reflect"
	"fmt"
)

type  Person  struct{
    Name string
    age int
}

func main(){
	u:=Person{
		"Jack",
		23,
	}
	// value to interface
	v:=reflect.ValueOf(u)
	// pointer to interface
	p:=reflect.ValueOf(&u)
	// struct cast to interface would copy object
	fmt.Println(v.CanSet(),v.FieldByName("Name").CanSet())
	fmt.Println(p.CanSet(),p.Elem().FieldByName("Name").CanSet())

}
