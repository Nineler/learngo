package main

import "fmt"

func updataslice(s []int) {
	s[0] = 100
	fmt.Println(s)
}

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Println("arr[2:6]=", arr[2:6])
	fmt.Println("arr[:6]=", arr[:6])
	s1 := arr[2:]
	fmt.Println("arr[2:]=", arr[2:])
	s2 := arr[:]
	fmt.Println("arr[:]=arr[]=", arr[:])
	updataslice(s1)
	updataslice(s2)
}
