package main

import (
	"fmt"
	"sync"
	"testing"
)

type Akun struct {
	mutex sync.RWMutex
	Saldo int
}

func (akun *Akun) WriteSaldo(amount int) {
	akun.mutex.Lock()
	akun.Saldo = akun.Saldo + amount
	akun.mutex.Unlock()
}

func (akun *Akun) GetAkun() {
	akun.mutex.RLock()
	saldoAnda := akun.Saldo
	akun.mutex.RUnlock()
	fmt.Println("saldo anda sekarang adalah => ", saldoAnda)
}

func TestAkun_WriteSaldo(t *testing.T) {

	var wg sync.WaitGroup

	numGoroutines := 100
	wg.Add(numGoroutines)
	akun := Akun{}
	for i := 0; i < numGoroutines; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 100; j++ {
				akun.WriteSaldo(1)
				akun.GetAkun()
			}
		}()
	}

	wg.Wait()
	akun.GetAkun()
}
