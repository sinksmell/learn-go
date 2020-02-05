package main

import (
	"fmt"
)

func main(){
	str:="hello,world!"
	for _,c:=range str{
		print(string(c))
	}
	// range will copy object
	arr:=[3]int{1,2,3}
	for i,v:=range arr{
		if i==0{
			// 修改arr
			arr[1],arr[2]=100,200
			fmt.Println(arr)
		}
		// i与v都是从复制品中获取的
		arr[i]=v+100
	}
	fmt.Println(arr)
	// 改用 引用类型不会复制底层数据
	arr2:=[]int{1,2,3}
	for i,v:=range arr2{
		if i==0{
			// 修改arr2
			arr2[1],arr2[2]=100,200
			fmt.Println(arr2)
		}
		// i与v都是从复制品中获取的
		arr2[i]=v+100
	}
	fmt.Println(arr2)
}