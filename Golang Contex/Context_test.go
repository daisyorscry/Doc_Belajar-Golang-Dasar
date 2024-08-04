package main

import (
	"context"
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	result := Print()
	assert.NotEqual(t, "hellos", result)
	assert.Equal(t, "hello", result)
}

func TestContext(t *testing.T) {
	background := context.Background()
	fmt.Println("Background context:", background)

	todo := context.TODO()
	fmt.Println("TODO context:", todo)

	if background == nil {
		t.Error("Expected background context to be non-nil")
	}

	if todo == nil {
		t.Error("Expected TODO context to be non-nil")
	}
}

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
