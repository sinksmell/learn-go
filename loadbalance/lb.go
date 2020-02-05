package main

import (
	"encoding/json"
	"fmt"
)

type  proxy  struct{
    HttpPort int `json:"httpPort"`
    HttpsPort int `json:"httpsPort"`
}


func main(){
	pr:=proxy{
		HttpPort:80,
	}

	str,_:=json.Marshal(pr)
	fmt.Println(string(str))

	str2:=""
	pr2:=&proxy{}
	json.Unmarshal([]byte(str2),pr2)
	fmt.Printf("%+v\n",pr2)
}