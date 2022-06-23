package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 1000; i++ {
		go func(ii int) {
			for {
				fmt.Printf("hello from"+"goroutine %d\n", ii)
			}
		}(i)
	}
	time.Sleep(time.Minute)
}
