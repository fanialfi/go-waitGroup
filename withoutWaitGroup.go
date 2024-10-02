package main

import "fmt"

var nilai = 0

// tanpa menggunakan sync.WaitGroup
func incrementWithoutWaitGroup() {
	nilai++
}

func main() {
	for i := 0; i <= 1000; i++ {
		go incrementWithoutWaitGroup()
	}

	fmt.Println(nilai)
}
