package alternate

import (
	"fmt"
	"sync"
)

// 这种使用channel的方法其实是有问题的，因为对于无缓冲的channel，一边进了之后，另一边的出会同时满足，所以如果去掉奇偶判断，就会出现错误的结果。
var ChanSignal = make(chan int)

var wg = sync.WaitGroup{}

func Method2() {
	wg.Add(2)
	go printAA()
	go printBB()
	wg.Wait()
}

func printAA() {
	defer wg.Done()
	for i := 0; i < 101; i++ {
		ChanSignal <- 1

		if i%2 == 1 {
			fmt.Println("AAAA:", i)
		}
	}
}

func printBB() {
	defer wg.Done()
	for i := 0; i < 101; i++ {
		<-ChanSignal

		if i%2 == 0 {
			fmt.Println("BBBB:", i)
		}
	}

}
