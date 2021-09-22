package main

import (
	"fmt"
	"net"
)

func handle(conn net.Conn) {
	defer conn.Close()

	for {
		buf := make([]byte, 100)
		n, err := conn.Read(buf)

		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("read data size %d msg: %s", n, string(buf[0:n]))
		msg := []byte("hello world \n")
		conn.Write(msg)
	}

}

func main() {
	fmt.Println(" Start server.... ")
	listen, err := net.Listen("tcp", "0.0.0.0:3000")

	if err != nil {
		fmt.Println("Listen failed! msg:", err)
		return
	}

	for {
		conn, errs := listen.Accept()

		if errs != nil {
			fmt.Println("accept failed")
			continue
		}

		go handle(conn)
	}
}
