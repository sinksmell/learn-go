package main

import "fmt"

type user struct {
	name string
}

func main() {
	m := make(map[int]user)
	m[1] = user{"user1"}

	// 取出的value是一个复制品
	// map被设计成不可寻址的 通过key来修改value中的值被禁止
	// m[1].name="Tom"   cannot assign to struct field m[1].name in map

	// 正确地写法
	u:=m[1]
	u.name="Tom"
	m[1]=u

	// 迭代时可以删除key
	for i := 0; i < 5; i++ {
		m:=map[int]string{
			0:"a",1:"a",2:"a",3:"a",4:"a",
		}
		for k:=range m {
			m[k+k]="x"
			delete(m,k)
		}
		fmt.Println(m)
	}

	// 可以使用 map[sting]struct{} 来模拟set

}
