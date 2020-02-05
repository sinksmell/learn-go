package main

import (
	"unsafe"
	"fmt"
)

// type change
// cannot convert ip (type *int) to type *float64
//func typechange(){
//	i:=10
//	ip:=&i
//	var fp *float64=(*float64)(ip)
//	fmt.Println(fp)
//}

func typechange2(){
	i:=10
	ip:=&i
	var fp *float64=(*float64)(unsafe.Pointer(ip))
	fmt.Println(*fp)
}

type myType int
type Pointer *myType
/*
任何指针都可以转换为unsafe.Pointer
unsafe.Pointer可以转换为任何指针
uintptr可以转换为unsafe.Pointer
unsafe.Pointer可以转换为uintptr
*/
// 通过unsafe.Pointer这个万能的指针，我们可以在*T之间做任何转换。
func typechange3(){
	type  user  struct{
	    name string
	    age int
	}

	u:=new(user)
	fmt.Println(*u)
	pName:=(*string)(unsafe.Pointer(u))
	*pName="sinksmell"
	pAge:=(*int)(unsafe.Pointer((uintptr)(unsafe.Pointer(u))+unsafe.Offsetof(u.age)))
	*pAge=20
	fmt.Println(*u)

}
func main(){
	typechange2()
	typechange3()
}
