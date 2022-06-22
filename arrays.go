package main

import "fmt"

func printfarray(arr *[5]int) {
	//range遍历数组
	arr[0] = 100
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func main() {
	var arr1 [5]int
	//:=方式必须要有初始值
	arr2 := [3]int{1, 2, 2}
	arr3 := [...]int{1, 2, 3, 4, 5}
	var grid [4][5]int

	fmt.Println(arr1, arr2, arr3)
	fmt.Println(grid)

	printfarray(&arr1)
	printfarray(&arr3)
	fmt.Println(arr1, arr3)
}
