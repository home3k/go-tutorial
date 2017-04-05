package main

import (
	"fmt"
	"reflect"
)

func main() {
	t := reflect.TypeOf(3)
	fmt.Println(t.String())
	fmt.Println(t)
	// 内部就是用reflect.TypeOf
	fmt.Printf("%T\n", 3)

	v := reflect.ValueOf(5)
	fmt.Println(v.String())
	fmt.Println(v)
	fmt.Printf("%v\n", 5)

	t = v.Type()
	fmt.Println(t)

	x := 2
	a := reflect.ValueOf(2)
	b := reflect.ValueOf(x)
	c := reflect.ValueOf(&x)
	d := c.Elem()

	fmt.Println(a.CanAddr())
	fmt.Println(b.CanAddr())
	fmt.Println(c.CanAddr())
	fmt.Println(d.CanAddr())

	z := reflect.ValueOf(&x).Elem()
	px := z.Addr().Interface().(*int)
	*px = 300
	fmt.Println(x)

	z.Set(reflect.ValueOf(400))
	fmt.Println(x)


}
