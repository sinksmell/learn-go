package yuheng

import (
	"fmt"
	"reflect"
)

type Data struct {
}

func (*Data) Test(x, y int) (int, int) {
	return x + 100, y + 100
}

func (*Data) Sum(s string, x ...int) string {
	c := 0
	for _, n := range x {
		c += n
	}
	return fmt.Sprintf(s, c)
}

func info(m reflect.Method) {
	t := m.Type
	fmt.Println(m.Name)
	// in 为传入的参数
	for i, n := 0, t.NumIn(); i < n; i++ {
		fmt.Printf(" in[%d] %v\n", i, t.In(i))
	}
	// out 为返回值
	for i, n := 0, t.NumOut(); i < n; i++ {
		fmt.Printf(" out[%d] %v\n", i, t.Out(i))
	}

}

// 从打印的结果来看 receiver 是被隐式传入方法中的 本例中第一个参数为 *main.Data
/*
Test
 in[0] *main.Data
 in[1] int
 in[2] int
 out[0] int
 out[1] int
Sum
 in[0] *main.Data
 in[1] string
 in[2] []int
 out[0] string
*/
func main() {
	d := new(Data)
	t := reflect.TypeOf(d)
	test, _ := t.MethodByName("Test")
	info(test)
	sum, _ := t.MethodByName("Sum")
	info(sum)

	// 动态调用方法
	v := reflect.ValueOf(d)
	exec := func(name string, in []reflect.Value) {
		m := v.MethodByName(name)
		out := m.Call(in)
		for _, value := range out {
			fmt.Println(value.Interface())
		}
	}
	fmt.Println("--- 动态调用方法 ---")
	// 不需要传入receiver
	exec("Test", []reflect.Value{
		reflect.ValueOf(1),
		reflect.ValueOf(2),
	})

	fmt.Println("-------------------")
	exec("Sum", []reflect.Value{
		reflect.ValueOf("result = %d"),
		reflect.ValueOf(1),
		reflect.ValueOf(2),
	})
}
