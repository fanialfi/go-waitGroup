package main

import (
	"fmt"
	"sync"
)

var nilai = 0
var mu sync.Mutex

// tanpa menggunakan sync.WaitGroup
func incrementWithoutWaitGroup(mu *sync.Mutex) {
	// menggunakan mutex untuk melindingi variabel,
	// jadi hanya 1 goroutine yang dapat mengakses variabel
	mu.Lock()
	nilai++
	mu.Unlock()
}

func main() {
	for i := 0; i < 1000; i++ {
		go incrementWithoutWaitGroup(&mu)
	}

	fmt.Println(nilai)
}
