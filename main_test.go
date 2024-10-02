package main

import "testing"

func BenchmarkWithoutWaitGroup(b *testing.B) {
	for i := 0; i < b.N; i++ {
		go incrementWithoutWaitGroup(&mu)
	}
}

func BenchmarkWithWaitGroup(b *testing.B) {
	for i := 0; i < b.N; i++ {
		wg.Add(1)
		go func() {
			incrementWithWaitGroup(&wg, &mu)
		}()
		wg.Wait()
	}
}
