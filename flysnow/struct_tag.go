package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type  User  struct{
    Name string `name`
    Age int `age`
}

// string to struct
func decode(){
	var u User
	h:=`{"name":"sinksmell","age":15}`
	if err:=json.Unmarshal([]byte(h),&u);err!=nil{
		fmt.Println(err)
	}else{
		fmt.Println(u)
	}

}

// get tag info
func getTag(){
	var u User
	t:=reflect.TypeOf(u)
	for i ,n:= 0, t.NumField(); i <n; i++ {
		sf:=t.Field(i)
		fmt.Println(sf.Tag)
	}
}

// get value by key in tag info
type  People  struct{
    Name string `json:"name" orm:"o_name"`
    Age int `json:"age" orm:"o_age"`
}
func getTagValue(){
	var p People
	t:=reflect.TypeOf(p)
	for i, n := 0, t.NumField(); i < n; i++ {
		sf:=t.Field(i)
		fmt.Println(sf.Tag.Get("json"),sf.Tag.Get("orm"))
	}
}

func main(){
	decode()
	getTag()
	getTagValue()
}
