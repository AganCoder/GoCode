package main

import (
	"fmt"
	"reflect"
)
// Elem 返回指针，数组，切片。字典或通道的基本类型
func main() {
	 fmt.Println(reflect.TypeOf(map[string]int{}).Elem())
	fmt.Println(reflect.TypeOf(map[string]int32{}).Elem())

}
