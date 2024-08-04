package main

import (
	"sync"
	"testing"
)

// ini adalah contoh untuk menguji seberapa ringan goroutines
// perulangan ini akan running goroutines sebanyak 100_0000k gorotuines
// anggapannya seperti kita membuka sebanyak 100_0000k aplikasi
func TestLoopGoroutines(t *testing.T) {
	var wg sync.WaitGroup

	numGoroutines := 100000
	wg.Add(numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		func() {
			defer wg.Done()
			go CreateGolangGoroutines(i)
		}()
	}
	wg.Wait()
}
