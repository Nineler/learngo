package main

import (
	"fmt"
	"time"
)

func worker(id int, c chan int) {
	for n := range c {
		fmt.Printf("worker %d received %c\n", id, n)
	}
}
func Creatworker(id int) chan<- int {
	c := make(chan int)
	go worker(id, c)
	return c
}

func chanDoem() {
	var channels [10]chan<- int
	for i := 0; i < 10; i++ {
		channels[i] = Creatworker(i)
	}
	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}
	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i
	}
	time.Sleep(time.Millisecond)
}

func bufferedChannel() {
	c := make(chan int, 5) //5为缓冲区大小，表示当没有人接收channel时，可以有5个缓冲区放置数据
	go worker(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	c <- 'e'
	//c<-'f',第6个超出缓冲区大小，程序报错
	time.Sleep(time.Millisecond)
}
func channelClose() {
	c := make(chan int, 5) //5为缓冲区大小，表示当没有人接收channel时，可以有5个缓冲区放置数据
	go worker(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	c <- 'e'
	close(c)
	//c<-'f',第6个超出缓冲区大小，程序报错
	time.Sleep(time.Millisecond)
}

func main() {
	chanDoem()
	bufferedChannel()
	channelClose()
}
