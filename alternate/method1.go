package alternate

import (
	"fmt"
	"sync"
)

// 两个协程
func printA(c1, c2 chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i <= 100; i++ {
		<-c1
		if i%2 == 0 {
			fmt.Println("goroutine A is printing:", i)
		}
		c2 <- 100
	}
}
func printB(c1, c2 chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i <= 100; i++ {
		<-c2
		if i%2 == 1 {
			fmt.Println("goroutine B is printing:", i)
		}
		c1 <- 100
	}
}

func Method1() {
	wg := &sync.WaitGroup{}
	c1 := make(chan int,1) // 要加缓冲区，不然死锁
	c2 := make(chan int,1)

	wg.Add(2)
	go printA(c1, c2, wg)
	go printB(c1, c2, wg)
	c1<-1
	wg.Wait()
}
