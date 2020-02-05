package yuheng

import (
	"reflect"
	"fmt"
)
// 字段标签可实现简单元数据编程 如标记ORM Model属性
type  User  struct{
    UserName string `filed:"username" type:"nvarchar(20)"`
    Age int `filed:"age" type:"tinyint"`
}


type  Admin  struct{
    User
	title string
}

func main(){
	// 通过反射获取实例对象字段信息
	var u Admin
	t:=reflect.TypeOf(u)

	for i, n := 0, t.NumField(); i < n; i++ {
		f:=t.Field(i)
		fmt.Println(f.Name,"\t",f.Type)
	}

	// 通过反射获取指针所指对象的字段信息
	u2:=new(Admin)
	t2:=reflect.TypeOf(u2)
	if t2.Kind()==reflect.Ptr{
		t2=t2.Elem()
	}
	for i, n := 0, t2.NumField(); i < n; i++ {
		f:=t2.Field(i)
		fmt.Println(f.Name,"\t",f.Type)
	}

	// 获取字段的tag 信息
	fmt.Println("Use tag")
	var u3 User
	t3:=reflect.TypeOf(u3)
	f,_:=t3.FieldByName("UserName")
	fmt.Println(f.Tag)
	fmt.Println(f.Tag.Get("field"))
	fmt.Println(f.Tag.Get("type"))
}