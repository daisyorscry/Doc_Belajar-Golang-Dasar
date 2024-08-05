package main

import (
	"fmt"
	"runtime"
	"testing"
)

func TestGoMaXProc(t *testing.T) {
	totalCpu := runtime.NumCPU()
	fmt.Println(totalCpu)

	totalThread := runtime.GOMAXPROCS(100)
	fmt.Println(totalThread)

	totalThreadd := runtime.GOMAXPROCS(-1)
	fmt.Println(totalThreadd)

	totolGoroutines := runtime.NumGoroutine()
	fmt.Println(totolGoroutines)

}
