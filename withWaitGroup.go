package main

import (
	"fmt"
	"sync"
)

var value = 0
var mutex sync.Mutex

func incrementWithWaitGroup(wg *sync.WaitGroup, mu *sync.Mutex) {
	defer wg.Done()

	mu.Lock()
	value++
	mu.Unlock()
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go incrementWithWaitGroup(&wg, &mutex)
	}

	wg.Wait()
	fmt.Println(value)
}
