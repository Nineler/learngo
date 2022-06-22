package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

var (
	aa = 123
	bb = true
	cc = "abc"
)

//定义变量的多种方式
func variableZeroValue() {
	var a int
	var s string
	fmt.Printf("%d,%q\n", a, s)
}

func variableInitialValue() {
	var a, b int = 3, 4
	var s string = "kkk"
	fmt.Println(a, b, s)
}

func variableTypeDeduction() {
	var a, b, c = 3, true, "def"
	fmt.Println(a, b, c)
}

func variableShort() {
	//:=形式只能在函数内部使用,多次对变量赋值第一次：=就行了
	a, b, c := 3, true, "asd"
	a = 4
	fmt.Println(a, b, c)
}

//欧拉公式
func euler() {
	fmt.Println(cmplx.Exp(1i*math.Pi) + 1)
}

//强制转换
func Triangle() {
	a, b := 3, 4
	var c int = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

//常量
func consts() {
	const (
		filename = "123.txt"
		a, b     = 3, 4
	)
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(filename, c)
}

//枚举
func enums() {
	//iota自增值
	const (
		cpp = iota
		java
		php
	)
	fmt.Println(cpp, java, php)
}
func main() {
	fmt.Println("hello world")
	variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	variableShort()
	fmt.Println(aa, bb, cc)

	euler()
	Triangle()
	consts()
	enums()
}
