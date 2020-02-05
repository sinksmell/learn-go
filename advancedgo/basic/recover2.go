package main

import "fmt"

func main() {

	defer func() {
		if v := recover(); v != nil {
			switch v.(type) {
			case int:
				fmt.Println(v.(int))
			default:
				fmt.Println("unknown type")
			}
		}
	}()
	panic(1)
	fmt.Println("running next ...")
}
