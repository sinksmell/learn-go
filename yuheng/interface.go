package yuheng

import "fmt"

type  User  struct{
    name string
    id int
}

func(u *User)String()string{
	return fmt.Sprintf("%s %d",u.name,u.id)
}

// 函数直接'实现'接口
type  Tester interface {
    Do()
}

type FuncDo func()

func(self FuncDo)Do(){
	self()
}

func main(){
	var o interface{}=&User{"Tom",1}
	// 判断是否为某个接口
	if i,ok:=o.(fmt.Stringer);ok{
		fmt.Println(i)
	}

	// 转换成具体类型
	// u:=o.(User) interface conversion: interface {} is *main.User, not main.User
	u:=o.(*User)
	fmt.Println(u)

	// 这是什么骚操作?
	var t Tester=FuncDo(func() {
		fmt.Println("hello,world!")
	})
	t.Do()
}
