package main

import (
	"fmt"
	"net/http"
)

func main() {
	//conn, err := net.Dial("tcp", "127.0.0.1:3000")
	//if err != nil {
	//	fmt.Println("err, dialing:", err.Error())
	//	return
	//}
	//defer conn.Close()
	//
	//inputReader := bufio.NewReader(os.Stdin)
	//
	//
	//for  {
	//	str, _ := inputReader.ReadString('\n')
	//	data := strings.Trim(str, "\n")
	//
	//	if data == "quit" {
	//		return
	//	}
	//
	//	_,err := conn.Write([]byte(data))
	//	if err != nil {
	//		fmt.Println("send data error:", err)
	//		return
	//	}
	//
	//	buf := make([]byte, 512)
	//
	//	n, err := conn.Read(buf)
	//
	//	fmt.Println("from server", string((buf[:n])))
	//
	//}
	resp, err := http.Get("https://service.paper.meiyuan.in/api/v2/columns")
	fmt.Println(resp, err)

}
