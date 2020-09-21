package main

func main() {
	done := make(chan struct{})

	data := [] chan int {
		make(chan int, 3), // 数据缓冲
	}
	// select 可以用于接收 也可以用于发送
	go func() {
		defer close(done)

		for i := 0; i < 10; i ++ {
			select {
			case data[len(data) - 1] <- i:

			default:
				data = append(data, make(chan int, 3))
			}
		}
	}()

	<- done
	for i := 0; i < len(data); i++ { // 显示所有数据
		c := data[i]
		close(c)

		for x := range  c {
			println(x)
		}
	}

}
