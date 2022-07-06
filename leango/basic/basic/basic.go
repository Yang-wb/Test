package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

//函数外必须使用 var关键字
var aa = 3
var ss = "kkk"

var (
	cc = 4
	dd = 54
)

func variableZeroValue() {
	var a int
	var s string
	fmt.Printf("%d %q\n", a, s)
}

func variableInitialValue() {
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Printf("%d %d %q\n", a, b, s)
}

func variableTypeDeduction() {
	var a, b, c, d = 3, 4, true, "abc"
	fmt.Println(a, b, c, d)
}

func variableShorter() {
	a, b, c, s := 3, 4, true, "abc"
	fmt.Println(a, b, c, s)
}

func euler() {
	//c := 3 + 4i
	//fmt.Println(cmplx.Abs(c))

	//fmt.Println(cmplx.Pow(math.E, 1i*math.Pi) + 1)

	//fmt.Println(cmplx.Exp(1i*math.Pi) + 1)

	fmt.Printf("%.3f \n", cmplx.Exp(1i*math.Pi)+1)
}

func triangle() {
	var a, b int = 3, 4
	fmt.Println(calcTriangle(a, b))
}

func calcTriangle(a, b int) int {
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	return c
}

const f = "123"
const (
	j = 1
	k = 2
)

func consts() {
	const filename = "abc.txt"
	const a, b = 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

func enmus() {
	//const (
	//	cpp    = 0
	//	java   = 1
	//	python = 2
	//	golang = 3
	//)

	const (
		cpp = iota
		_   //跳过
		java
		python
		golang
	)

	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
}

//func main() {
//	fmt.Println("hello")
//	variableZeroValue()
//	variableInitialValue()
//	variableTypeDeduction()
//	variableShorter()
//	fmt.Println(aa, ss, dd, cc)
//}
