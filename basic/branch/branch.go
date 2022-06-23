package main

import (
	"fmt"
	"io/ioutil"
)

func grade(score int) string {
	g := ""
	switch {
	case score < 0 || score > 100:
		panic(fmt.Sprintf(
			"Wrong score: %d", score))
	case score < 60:
		g = "f"
	case score < 80:
		g = "c"
	case score < 90:
		g = "b"
	case score < 100:
		g = "a"
	}
	return g
}

func main() {
	const filename = "abc.txt"
	if contents, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}

	fmt.Println(
		grade(15),
		grade(70),
		grade(89),
		grade(99),
	)
}
