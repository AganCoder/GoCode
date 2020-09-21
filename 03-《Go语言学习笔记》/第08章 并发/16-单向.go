package main

import (
	"fmt"
)

func main() {
	var send = make(chan <- int)
	fmt.Printf("%v, %T", send, send)

	var recv = make(<- chan  int)

	fmt.Printf("%v, %T", recv, recv)

	//go func() {
	//	for recv
	//
	//}()
	//
	//go func() {
	//	defer close(send)
	//
	//	for i :=0; i<3; i++ {
	//		send <- i
	//	}
	//}()
}