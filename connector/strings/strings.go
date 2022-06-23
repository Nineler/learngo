package main

import "fmt"

func main() {
	s := "Yes我爱慕课网！" //UTF-8
	fmt.Println(s)
	for _, b := range []byte(s) {
		fmt.Printf("%X", b)
	}
	fmt.Println()
	for i, ch := range s {
		fmt.Printf("(%d %X)", i, ch)
	}
	fmt.Println()
	for i, ch := range []rune(s) {
		fmt.Printf("(%d %c)", i, ch)
	}
}
