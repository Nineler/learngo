package main

import (
	"fmt"
)

func main() {
	m := map[string]string{
		"name":    "ccmouse",
		"course":  "golang",
		"site":    "imooc",
		"quality": "motbad",
	}
	m2 := make(map[string]int)
	var m3 map[string]int
	fmt.Println(m, m2, m3)
	for k, v := range m {
		fmt.Println(k, v)
	}
	delete(m, "name")
	fmt.Println(m)
}
