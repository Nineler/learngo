package main

import (
	"GOProject/functional/fib"
	"bufio"
	"fmt"
	"os"
)

//一个函数有多个defer时，defer类似于栈的顺序，后进先出（后面的defer先执行）
/*func tryDefer() {
	defer fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println(3)
}*/
func writeFile(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		if pathError, ok := err.(*os.PathError); !ok {
			panic(err)
		} else {
			fmt.Printf("%s,%s,%s\n",
				pathError.Op,
				pathError.Path,
				pathError.Err)
		}
		return
	}
	defer file.Close()
	f := fib.Fibonacci()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, f())
	}
}
func main() {
	writeFile("fib.txt")
}
