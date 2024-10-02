package main

import (
	"fmt"
	"sync"
)

var nilai = 0
var mu sync.Mutex
var wg sync.WaitGroup

// tanpa menggunakan sync.WaitGroup
func incrementWithoutWaitGroup(mu *sync.Mutex) {
	// menggunakan mutex untuk melindingi variabel,
	// jadi hanya 1 goroutine yang dapat mengakses variabel
	mu.Lock()
	nilai++
	mu.Unlock()
}

// func main() {
// 	for i := 0; i < 1000; i++ {
// 		go incrementWithoutWaitGroup(&mu)
// 	}

// 	fmt.Println(nilai)
// }

func incrementWithWaitGroup(wg *sync.WaitGroup, mu *sync.Mutex) {
	defer wg.Done()

	mu.Lock()
	nilai++
	mu.Unlock()
}

func main() {

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go incrementWithWaitGroup(&wg, &mu)
	}

	wg.Wait()
	fmt.Println(nilai)
}
