package main

import "time"

func main() {
	done := make(chan struct{})
	c := make(chan int)

	go func() {
		defer close(done)

		for {
			time.Sleep(time.Second)
			// 对于同一通道，接受操作是阻塞的，直到发送者可用，如果通道中没有数据，接受者就阻塞了！
			x, ok := <-c
			if !ok { // 据此判断通道是否被关闭
				return
			}
			println(x)
		}
	}()

	// 对于同一个通道，发送操作在接受者准备好之前是阻塞的
	// 如果通道中无人接收数据，就无法再传给通道传入其他数据，也就是新的数据无法在通道非空的情况下传入
	// 通道值被接收后，channel 会再次变为可用状态

	c <- 1
	println("sended 1")
	c <- 2
	println("sended 2")
	c <- 3
	println("sended 3")

	close(c)

	<-done
}
