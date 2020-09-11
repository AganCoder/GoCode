package main

import (
	"math"
	"runtime"
	"sync"
	"time"
)

func count() {
	x := 0
	for i :=0; i < math.MaxUint32; i++ {
		x += 1
	}

	println(x)
}

func test(n int) {
	for i := 0; i < n; i ++ {
		count()
	}
}

func test2(n int) {
	var wg sync.WaitGroup
	wg.Add(n)

	for i:=0; i<n; i++ {
		go func() {
			count()
			wg.Done()
		}()
	}

	wg.Wait()
}

func main()  {
	n := runtime.GOMAXPROCS(0)
	now := time.Now()

	test(n)
	println(time.Now().Sub(now))

	now = time.Now()

	test2(n)

	println(time.Now().Sub(now))

}