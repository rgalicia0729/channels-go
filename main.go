package main

import (
	"fmt"
	"sync"
	"time"
)

func doSomething(i int, wg *sync.WaitGroup, c chan int) {
	defer wg.Done()

	time.Sleep(2 * time.Second)
	fmt.Printf("Id: %d\n", i)

	<-c
}

func main() {
	var wg sync.WaitGroup
	c := make(chan int, 4)

	for i := 1; i <= 10; i++ {
		wg.Add(1)
		c <- 1

		go doSomething(i, &wg, c)
	}

	wg.Wait()
}
