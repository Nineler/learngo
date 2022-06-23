package main

import (
	"fmt"
	"sync"
)

func doworker(id int,
	w Worker) {
	for n := range w.in {
		fmt.Printf("worker %d received %c\n",
			id, n)
		//go func() { done <- true }()
		w.done()
	}
}

type Worker struct {
	in   chan int
	done func()
	//done chan bool
}

func Creatworker(id int, wg *sync.WaitGroup) Worker {
	w := Worker{
		in: make(chan int),
		done: func() {
			wg.Done()
		},
	}
	go doworker(id, w)
	return w
}

func chanDoem() {
	var wg sync.WaitGroup
	var workers [10]Worker
	for i := 0; i < 10; i++ {
		workers[i] = Creatworker(i, &wg)
	}
	wg.Add(20) //等待任务个数
	for i, worker := range workers {
		worker.in <- 'a' + i
	}
	for i, worker := range workers {
		worker.in <- 'A' + i
	}
	wg.Done()
	wg.Wait()
	/*for _, worker := range workers {
		<-worker.done
		<-worker.done
	}*/
}

func main() {
	chanDoem()
}
