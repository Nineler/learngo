package main

import (
	"fmt"
	"sync"
	"time"
)

type atomicInt struct {
	value int
	lock  sync.Mutex
}

func (a *atomicInt) incremeent() {
	fmt.Println("safe increment")
	func() {
		a.lock.Lock()
		a.value++
		defer a.lock.Unlock()
	}()

}

func (a *atomicInt) get() int {
	a.lock.Lock()
	defer a.lock.Unlock()
	return a.value
}

func main() {
	var a atomicInt
	a.incremeent()
	go func() {
		a.incremeent()
	}()
	time.Sleep(time.Millisecond)
	fmt.Println(a.value)
}
