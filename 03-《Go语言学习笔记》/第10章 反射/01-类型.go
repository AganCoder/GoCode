package main

import (
	"fmt"
	"reflect"
)

type X int


// Type: 表示真实类型（静态类型）  --- 表示 type 类型 X
// Kind：表示基础结构（底层类型）类别 -- 表示最底层的 int

func main() {
	var a X = 10
	t := reflect.TypeOf(a)

	fmt.Println(t.Name(), t.Kind())
}
