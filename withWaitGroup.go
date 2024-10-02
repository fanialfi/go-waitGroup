package main

import (
	"fmt"
	"sync"
)

var value = 0

func incrementWithWaitGroup(wg *sync.WaitGroup) {
	defer wg.Done()
	value++
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go incrementWithWaitGroup(&wg)
	}

	wg.Wait()
	fmt.Println(value)
}
