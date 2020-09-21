package main

// 异步模式在缓冲区未满或者数据未读完前，不会进行阻塞
func main() {
	c := make(chan  int, 3)

	c <- 1
	c <- 2

	println(<-c)
	println(<-c)
}
